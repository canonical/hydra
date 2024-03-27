ALTER TABLE hydra_oauth2_device_code DROP FOREIGN KEY IF EXISTS hydra_oauth2_device_code_challenge_id_fk;
ALTER TABLE hydra_oauth2_device_code DROP FOREIGN KEY IF EXISTS hydra_oauth2_device_code_client_id_fk;
ALTER TABLE hydra_oauth2_device_code DROP FOREIGN KEY IF EXISTS hydra_oauth2_device_code_nid_fk_idx;

DROP TABLE IF EXISTS hydra_oauth2_device_code;

ALTER TABLE hydra_oauth2_user_code DROP FOREIGN KEY IF EXISTS hydra_oauth2_user_code_challenge_id_fk;
ALTER TABLE hydra_oauth2_user_code DROP FOREIGN KEY IF EXISTS hydra_oauth2_user_code_client_id_fk;
ALTER TABLE hydra_oauth2_user_code DROP FOREIGN KEY IF EXISTS hydra_oauth2_user_code_nid_fk_idx;

DROP TABLE IF EXISTS hydra_oauth2_user_code;

ALTER TABLE hydra_oauth2_flow DROP COLUMN IF EXISTS device_challenge_id;
ALTER TABLE hydra_oauth2_flow DROP COLUMN IF EXISTS device_code_request_id;
ALTER TABLE hydra_oauth2_flow DROP COLUMN IF EXISTS device_verifier;
ALTER TABLE hydra_oauth2_flow DROP COLUMN IF EXISTS device_csrf;
ALTER TABLE hydra_oauth2_flow DROP COLUMN IF EXISTS device_user_code_accepted_at;
ALTER TABLE hydra_oauth2_flow DROP COLUMN IF EXISTS device_was_used;
ALTER TABLE hydra_oauth2_flow DROP COLUMN IF EXISTS device_handled_at;
ALTER TABLE hydra_oauth2_flow DROP COLUMN IF EXISTS device_error;
