CREATE TABLE tags (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  title varchar(191) NOT NULL,
  cover_image varchar(255) DEFAULT NULL,
  background_image varchar(255) DEFAULT NULL,
  description text DEFAULT NULL,
  color varchar(191) DEFAULT NULL,
  slug varchar(191) NOT NULL,
  user_id bigint(20) unsigned NOT NULL,
  deleted_at datetime(3) DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY idx_tags_title (title),
  KEY idx_tags_user_id (user_id),
  KEY idx_tags_created_at (created_at),
  KEY idx_tags_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
