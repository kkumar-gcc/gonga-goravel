CREATE TABLE likes (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  user_id bigint(20) unsigned NOT NULL,
  likeable_id bigint(20) unsigned NOT NULL,
  likeable_type varchar(255) NOT NULL,
  deleted_at datetime(3) DEFAULT NULL,
  PRIMARY KEY (id),
  KEY idx_likes_user_id (user_id),
  KEY idx_likes_likeable_id (likeable_id),
  KEY idx_likes_likeable_type (likeable_type),
  KEY idx_likes_created_at (created_at),
  KEY idx_likes_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
