CREATE TABLE mentions (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  user_id bigint(20) unsigned NOT NULL,
  owner_id bigint(20) unsigned NOT NULL,
  owner_type varchar(191) NOT NULL,
  position int(10) DEFAULT NULL,
  deleted_at datetime(3) DEFAULT NULL,
  PRIMARY KEY (id),
  KEY idx_mentions_user_id (user_id),
  KEY idx_mentions_owner_id (owner_id),
  KEY idx_mentions_owner_type (owner_type),
  KEY idx_mentions_created_at (created_at),
  KEY idx_mentions_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
