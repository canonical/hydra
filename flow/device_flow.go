// Copyright © 2024 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package flow

import (
	"context"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"

	"github.com/ory/hydra/v2/client"
	"github.com/ory/hydra/v2/oauth2/flowctx"
	"github.com/ory/hydra/v2/x"
	"github.com/ory/x/sqlcon"
	"github.com/ory/x/sqlxx"
)

// TODO(nsklikas): These constants don't start from 1 to avoid conflicts with the flow
// table, we need to evaluate if this makes sense. Should/Can we merge these two tables?
//
// DeviceFlowState* constants enumerate the states of a flow. The below graph
// describes possible flow state transitions.
//
// graph TD
//
//	DEVICE_INITIALIZED --> DEVICE_UNUSED
//	DEVICE_UNUSED --> DEVICE_USED
//	DEVICE_UNUSED --> DEVICE_ERROR
const (
	// DeviceFlowStateLoginInitialized applies before the login app either
	// accepts or rejects the login request.
	DeviceFlowStateInitialized = int16(7)

	// DeviceFlowStateUnused indicates that the login has been authenticated, but
	// the User Agent hasn't picked up the result yet.
	DeviceFlowStateUnused = int16(8)

	// DeviceFlowStateUsed indicates that the User Agent is requesting consent and
	// Hydra has invalidated the login request. This is a short-lived state
	// because the transition to DeviceFlowStateConsentInitialized should happen while
	// handling the request that triggered the transition to DeviceFlowStateUsed.
	DeviceFlowStateUsed = int16(9)
	// DeviceFlowStateError indicates that an error has occured in the handling of the
	// device verification flow.
	DeviceFlowStateError = int16(127)
)

// DeviceFlow contains information about the device authorization flow.
type DeviceFlow struct {
	ID  string    `db:"challenge"`
	NID uuid.UUID `db:"nid"`

	// RequestID is the device authorization request's ID.
	RequestID string `db:"request_id"`
	// RequestURL is the original OAuth 2.0 Device Authorization URL requested by the OAuth 2.0 client. This URL is typically not
	// needed, but might come in handy if you want to deal with additional request parameters.
	//
	// required: true
	RequestURL string `db:"request_url"`

	// RequestedScope contains the OAuth 2.0 Scope requested by the OAuth 2.0 Client.
	//
	// required: true
	RequestedScope sqlxx.StringSliceJSONFormat `db:"requested_scope"`

	// RequestedAudience contains the access token audience as requested by the OAuth 2.0 Client.
	//
	// required: true
	RequestedAudience sqlxx.StringSliceJSONFormat `db:"requested_at_audience"`

	// Client is the OAuth 2.0 Client that initiated the request.
	//
	// required: true
	Client   *client.Client `db:"-"`
	ClientID string         `db:"client_id"`

	Verifier string `db:"verifier"`
	CSRF     string `db:"csrf"`

	RequestedAt time.Time `db:"requested_at"`

	State int16 `db:"state"`
	// The user_code was already handled.
	WasHandled bool `db:"was_handled"`
	// HandledAt contains the timestamp the device user verification request was handled.
	HandledAt sqlxx.NullTime      `db:"handled_at"`
	Error     *RequestDeniedError `db:"error"`
}

// NewDeviceFlow return a new DeviceFlow from a DeviceUserAuthRequest.
func NewDeviceFlow(r *DeviceUserAuthRequest) *DeviceFlow {
	f := &DeviceFlow{
		ID:                r.ID,
		Client:            r.Client,
		RequestURL:        r.RequestURL,
		Verifier:          r.Verifier,
		CSRF:              r.CSRF,
		RequestedAt:       r.RequestedAt,
		RequestedScope:    r.RequestedScope,
		RequestedAudience: r.RequestedAudience,
		WasHandled:        r.WasHandled,
		HandledAt:         r.HandledAt,
		State:             DeviceFlowStateInitialized,
	}
	if r.Client != nil {
		f.ClientID = r.Client.GetID()
	}
	return f
}

// GetDeviceUserAuthRequest return the DeviceUserAuthRequest from a DeviceFlow.
func (f *DeviceFlow) GetDeviceUserAuthRequest() *DeviceUserAuthRequest {
	return &DeviceUserAuthRequest{
		ID:                f.ID,
		Client:            f.Client,
		RequestURL:        f.RequestURL,
		Verifier:          f.Verifier,
		CSRF:              f.CSRF,
		RequestedAt:       f.RequestedAt,
		RequestedScope:    f.RequestedScope,
		RequestedAudience: f.RequestedAudience,
		WasHandled:        f.WasHandled,
		HandledAt:         f.HandledAt,
	}
}

// GetHandledDeviceUserAuthRequest return the HandledDeviceUserAuthRequest from a DeviceFlow.
func (f *DeviceFlow) GetHandledDeviceUserAuthRequest() *HandledDeviceUserAuthRequest {
	return &HandledDeviceUserAuthRequest{
		ID:                  f.ID,
		Client:              f.Client,
		Request:             f.GetDeviceUserAuthRequest(),
		DeviceCodeRequestID: f.RequestID,
		RequestURL:          f.RequestURL,
		RequestedAt:         f.RequestedAt,
		RequestedScope:      f.RequestedScope,
		RequestedAudience:   f.RequestedAudience,
		WasHandled:          f.WasHandled,
		HandledAt:           f.HandledAt,
		Error:               f.Error,
	}
}

// HandleDeviceUserAuthRequest updates the flows fields from a handled request.
func (f *DeviceFlow) HandleDeviceUserAuthRequest(h *HandledDeviceUserAuthRequest) error {
	if f.WasHandled {
		return errors.WithStack(x.ErrConflict.WithHint("The user_code was already used and can no longer be changed."))
	}

	if f.State != DeviceFlowStateInitialized && f.State != DeviceFlowStateUnused && f.State != DeviceFlowStateError {
		return errors.Errorf("invalid flow state: expected %d/%d/%d, got %d", DeviceFlowStateInitialized, DeviceFlowStateUnused, DeviceFlowStateError, f.State)
	}

	if f.ID != h.ID {
		return errors.Errorf("flow device challenge ID %s does not match HandledDeviceUserAuthRequest ID %s", f.ID, h.ID)
	}

	f.State = DeviceFlowStateUnused
	if h.Error != nil {
		f.State = DeviceFlowStateError
	}
	f.Client = h.Client
	f.ClientID = h.Client.GetID()
	f.RequestID = h.DeviceCodeRequestID
	f.HandledAt = h.HandledAt
	f.WasHandled = h.WasHandled
	f.RequestedScope = h.RequestedScope
	f.RequestedAudience = h.RequestedAudience
	f.Error = h.Error

	return nil
}

// InvalidateDeviceRequest shifts the flow state to DeviceFlowStateUsed. This
// transition is executed upon device completion.
func (f *DeviceFlow) InvalidateDeviceRequest() error {
	if f.State != DeviceFlowStateUnused && f.State != DeviceFlowStateError {
		return errors.Errorf("invalid flow state: expected %d or %d, got %d", DeviceFlowStateUnused, DeviceFlowStateError, f.State)
	}
	if f.WasHandled {
		return errors.New("device verifier has already been used")
	}
	f.WasHandled = true
	f.State = DeviceFlowStateUsed
	return nil
}

// ToDeviceChallenge converts the flow into a device challenge.
func (f *DeviceFlow) ToDeviceChallenge(ctx context.Context, cipherProvider CipherProvider) (string, error) {
	return flowctx.Encode(ctx, cipherProvider.FlowCipher(), f, flowctx.AsDeviceChallenge)
}

// ToDeviceVerifier converts the flow into a device verifier.
func (f *DeviceFlow) ToDeviceVerifier(ctx context.Context, cipherProvider CipherProvider) (string, error) {
	return flowctx.Encode(ctx, cipherProvider.FlowCipher(), f, flowctx.AsDeviceVerifier)
}

// TableName returns the DeviceFlow database table name.
func (DeviceFlow) TableName() string {
	return "hydra_oauth2_device_flow"
}

// BeforeSave get clientID before storing the flow in the database.
func (f *DeviceFlow) BeforeSave(_ *pop.Connection) error {
	if f.Client != nil {
		f.ClientID = f.Client.GetID()
	}
	return nil
}

// AfterFind fetches the client object and populates the relevant DeviceFlow field.
func (f *DeviceFlow) AfterFind(c *pop.Connection) error {
	f.Client = &client.Client{}
	return sqlcon.HandleError(c.Where("id = ? AND nid = ?", f.ClientID, f.NID).First(f.Client))
}
