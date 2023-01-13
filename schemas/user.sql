  CREATE TABLE IF NOT EXISTS user (
  id bigint(20) NOT NULL AUTO_INCREMENT,
  username varchar(64) DEFAULT '',
  profile_img varchar(200) DEFAULT '',
  password varchar(280) DEFAULT '',
  email   varchar(64) DEFAULT '',
  token varchar(64) DEFAULT '',
  created_at bigint(20) DEFAULT '0',
  updated_at bigint(20) DEFAULT '0',
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;