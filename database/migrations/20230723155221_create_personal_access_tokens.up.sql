CREATE TABLE personal_access_tokens (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  deleted_at datetime(3) DEFAULT NULL,
  name varchar(255) NOT NULL,
  token varchar(255) NOT NULL,
  last_used_at datetime(3) DEFAULT NULL,
  expires_at datetime(3) DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY idx_personal_access_tokens_token (token),
  KEY idx_personal_access_tokens_expires_at (expires_at),
  KEY idx_personal_access_tokens_created_at (created_at),
  KEY idx_personal_access_tokens_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
