CREATE TABLE follows (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  follower_id bigint(20) unsigned NOT NULL,
  following_id bigint(20) unsigned NOT NULL,
  deleted_at datetime(3) DEFAULT NULL,
  PRIMARY KEY (id),
  KEY idx_follows_follower_id (follower_id),
  KEY idx_follows_following_id (following_id),
  KEY idx_follows_created_at (created_at),
  KEY idx_follows_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
