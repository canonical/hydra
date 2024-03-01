// Copyright © 2024 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package flow

import (
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func (f *DeviceFlow) setDeviceRequest(r *DeviceUserAuthRequest) {
	f.ID = r.ID
	f.CSRF = r.CSRF
	f.Verifier = r.Verifier
	f.Client = r.Client
	f.RequestURL = r.RequestURL
	f.RequestedAt = r.RequestedAt
	f.RequestedScope = r.RequestedScope
	f.RequestedAudience = r.RequestedAudience
	f.WasHandled = r.WasHandled
	f.HandledAt = r.HandledAt
}

func (f *DeviceFlow) setHandledDeviceRequest(r *HandledDeviceUserAuthRequest) {
	f.ID = r.ID
	f.Client = r.Client
	f.RequestURL = r.RequestURL
	f.RequestedAt = r.RequestedAt
	f.RequestedScope = r.RequestedScope
	f.RequestedAudience = r.RequestedAudience
	f.Error = r.Error
	f.RequestedAt = r.RequestedAt
	f.RequestID = r.DeviceCodeRequestID
	f.WasHandled = r.WasHandled
	f.HandledAt = r.HandledAt
}

func TestDeviceFlow_GetDeviceUserAuthRequest(t *testing.T) {
	t.Run("GetDeviceUserAuthRequest should set all fields on its return value", func(t *testing.T) {
		f := DeviceFlow{}
		expected := DeviceUserAuthRequest{}
		assert.NoError(t, faker.FakeData(&expected))
		f.setDeviceRequest(&expected)
		actual := f.GetDeviceUserAuthRequest()
		assert.Equal(t, expected, *actual)
	})
}

func TestDeviceFlow_GetHandledDeviceUserAuthRequest(t *testing.T) {
	t.Run("GetHandledDeviceUserAuthRequest should set all fields on its return value", func(t *testing.T) {
		f := DeviceFlow{}
		expected := HandledDeviceUserAuthRequest{}
		assert.NoError(t, faker.FakeData(&expected))
		f.setHandledDeviceRequest(&expected)
		actual := f.GetHandledDeviceUserAuthRequest()
		assert.NotNil(t, actual.Request)
		expected.Request = nil
		actual.Request = nil
		assert.Equal(t, expected, *actual)
	})
}

func TestDeviceFlow_NewDeviceFlow(t *testing.T) {
	t.Run("NewDeviceFlow and GetDeviceUserAuthRequest should use all DeviceUserAuthRequest fields", func(t *testing.T) {
		expected := &DeviceUserAuthRequest{}
		assert.NoError(t, faker.FakeData(expected))
		actual := NewDeviceFlow(expected).GetDeviceUserAuthRequest()
		assert.Equal(t, expected, actual)
	})
}

func TestDeviceFlow_HandleDeviceUserAuthRequest(t *testing.T) {
	t.Run(
		"HandleDeviceUserAuthRequest should ignore RequestedAt in its argument and copy the other fields",
		func(t *testing.T) {
			f := DeviceFlow{}
			assert.NoError(t, faker.FakeData(&f))
			f.State = DeviceFlowStateInitialized

			r := HandledDeviceUserAuthRequest{}
			assert.NoError(t, faker.FakeData(&r))
			r.ID = f.ID
			f.WasHandled = false
			f.RequestedAudience = r.RequestedAudience
			f.RequestedScope = r.RequestedScope
			f.RequestURL = r.RequestURL

			assert.NoError(t, f.HandleDeviceUserAuthRequest(&r))

			actual := f.GetHandledDeviceUserAuthRequest()
			assert.NotEqual(t, r.RequestedAt, actual.RequestedAt)
			r.Request = f.GetDeviceUserAuthRequest()
			actual.RequestedAt = r.RequestedAt
			assert.Equal(t, r, *actual)
		},
	)
}
