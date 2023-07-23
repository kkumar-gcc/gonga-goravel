CREATE TABLE comments (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  user_id bigint(20) unsigned NOT NULL,
  post_id bigint(20) unsigned NOT NULL,
  body text NOT NULL,
  parent_id bigint(20) unsigned DEFAULT NULL,
  deleted_at datetime(3) DEFAULT NULL,
  PRIMARY KEY (id),
  KEY idx_comments_user_id (user_id),
  KEY idx_comments_post_id (post_id),
  KEY idx_comments_created_at (created_at),
  KEY idx_comments_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
