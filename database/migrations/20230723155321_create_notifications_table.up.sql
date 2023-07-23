CREATE TABLE notifications (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY idx_notifications_email (email),
  KEY idx_notifications_created_at (created_at),
  KEY idx_notifications_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
