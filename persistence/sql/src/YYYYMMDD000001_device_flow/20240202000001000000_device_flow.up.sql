CREATE TABLE IF NOT EXISTS hydra_oauth2_device_code
(
    signature             VARCHAR(255) NOT NULL PRIMARY KEY,
    request_id            VARCHAR(40) NOT NULL,
    requested_at          TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    client_id             VARCHAR(255) NOT NULL,
    scope                 TEXT NOT NULL,
    granted_scope         TEXT NOT NULL,
    form_data             TEXT NOT NULL,
    session_data          TEXT NOT NULL,
    subject               VARCHAR(255) NOT NULL DEFAULT '',
    active                BOOL NOT NULL DEFAULT true,
    requested_audience    TEXT NULL DEFAULT '',
    granted_audience      TEXT NULL DEFAULT '',
    challenge_id          VARCHAR(40) NULL,
    nid                   UUID NULL
);
CREATE INDEX hydra_oauth2_device_code_request_id_idx ON hydra_oauth2_device_code (request_id, nid);
CREATE INDEX hydra_oauth2_device_code_client_id_idx ON hydra_oauth2_device_code (client_id, nid);
CREATE INDEX hydra_oauth2_device_code_challenge_id_idx ON hydra_oauth2_device_code (challenge_id, nid);

CREATE TABLE IF NOT EXISTS hydra_oauth2_user_code
(
    signature          VARCHAR(255) NOT NULL PRIMARY KEY,
    request_id         VARCHAR(255) NOT NULL,
    requested_at       TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    client_id          VARCHAR(255) NOT NULL,
    scope              TEXT         NOT NULL,
    granted_scope      TEXT         NOT NULL,
    form_data          TEXT         NOT NULL,
    session_data       TEXT         NOT NULL,
    subject            VARCHAR(255) NOT NULL DEFAULT '',
    active             BOOL         NOT NULL DEFAULT true,
    requested_audience TEXT         NULL DEFAULT '',
    granted_audience   TEXT         NULL DEFAULT '',
    challenge_id       VARCHAR(40)  NULL,
    nid                UUID         NULL
);
CREATE INDEX hydra_oauth2_user_code_request_id_idx ON hydra_oauth2_user_code (request_id, nid);
CREATE INDEX hydra_oauth2_user_code_client_id_idx ON hydra_oauth2_user_code (client_id, nid);
CREATE INDEX hydra_oauth2_user_code_challenge_id_idx ON hydra_oauth2_user_code (challenge_id, nid);

CREATE TABLE IF NOT EXISTS hydra_oauth2_device_flow (
    challenge             VARCHAR(255)                NOT NULL PRIMARY KEY,
    nid                   UUID                        NULL,
    request_id         VARCHAR(255) NOT NULL,
    request_url           TEXT                        NOT NULL,
    client_id             VARCHAR(255)                NOT NULL,
    verifier              VARCHAR(40)                 NOT NULL,
    csrf                  VARCHAR(40)                 NOT NULL,
    requested_at          TIMESTAMP                   DEFAULT CURRENT_TIMESTAMP NOT NULL,
    state                 INTEGER                     NOT NULL,
    requested_scope       TEXT                        NOT NULL DEFAULT '[]',
    requested_at_audience TEXT                        NOT NULL DEFAULT '[]',
    was_handled    BOOLEAN                     DEFAULT false NOT NULL,
    handled_at  TIMESTAMP                   NULL,
    error                 TEXT                        NULL
);

CREATE INDEX hydra_oauth2_device_flow_verifier_idx ON hydra_oauth2_device_flow (verifier, nid);
CREATE INDEX hydra_oauth2_device_flow_challenge_idx ON hydra_oauth2_device_flow (challenge, nid);
CREATE INDEX hydra_oauth2_device_flow_cid_idx ON hydra_oauth2_device_flow (client_id);
ALTER TABLE hydra_oauth2_flow ADD COLUMN device_flow_id VARCHAR(255) NULL;
ALTER TABLE hydra_oauth2_flow ADD COLUMN device_code_request_id VARCHAR(255) NULL;
