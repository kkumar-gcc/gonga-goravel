CREATE TABLE medias (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  url varchar(255) NOT NULL,
  type varchar(255) NOT NULL,
  owner_id bigint(20) unsigned NOT NULL,
  owner_type varchar(255) NOT NULL,
  PRIMARY KEY (id),
  KEY idx_medias_owner_id_owner_type (owner_id, owner_type),
  KEY idx_medias_created_at (created_at),
  KEY idx_medias_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
