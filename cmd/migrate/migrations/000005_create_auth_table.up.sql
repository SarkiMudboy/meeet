CREATE TABLE auth (
  auth_id SMALLINT UNSIGNED AUTO_INCREMENT,
  user_id SMALLINT UNSIGNED NOT NULL,
  session_token VARCHAR(60),
  csrf_token VARCHAR(60),
  password_hash VARCHAR(500),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT pk_user_auth PRIMARY KEY (auth_id),
  CONSTRAINT fk_user_auth FOREIGN KEY(user_id) REFERENCES users(user_id)
);
