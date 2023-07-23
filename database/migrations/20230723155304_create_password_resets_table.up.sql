CREATE TABLE password_resets (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  email varchar(255) NOT NULL,
  token varchar(255) NOT NULL,
  expires_at datetime(3) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY idx_password_resets_email (email),
  KEY idx_password_resets_created_at (created_at),
  KEY idx_password_resets_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
