DROP TABLE IF EXISTS users_lookup;
DROP TABLE IF EXISTS user_read;
DROP TABLE IF EXISTS user;

CREATE TABLE users_lookup (
  id INT NOT NULL,
  keyspace_id VARBINARY(128),

  PRIMARY KEY(id)

);

CREATE TABLE user (
  id INT NOT NULL,
  timestamp BIGINT DEFAULT 0,
  uid CHAR(5) DEFAULT NULL,
  name CHAR(9) DEFAULT NULL,
  gender CHAR(7) DEFAULT NULL,
  email CHAR(10) DEFAULT NULL,
  phone CHAR(10) DEFAULT NULL,
  dept CHAR(9) DEFAULT NULL,
  grade CHAR(7) DEFAULT NULL,
  language CHAR(3) DEFAULT NULL,
  region VARBINARY(256),
  role CHAR(6) DEFAULT NULL,
  prefer_tags CHAR(7) DEFAULT NULL,
  obtained_credits CHAR(3) DEFAULT NULL,
  
  PRIMARY KEY(id)

) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE user_read (
  id INT NOT NULL,
  timestamp BIGINT DEFAULT 0,
  uid INT NOT NULL,
  aid BIGINT NOT NULL,
  read_or_not CHAR(2) DEFAULT NULL,
  read_time_length CHAR(3) DEFAULT NULL,
  read_sequence CHAR(2) DEFAULT NULL,
  agree_or_not CHAR(2) DEFAULT NULL,
  comment_or_not CHAR(2) DEFAULT NULL,
  share_or_not CHAR(2) DEFAULT NULL,
  comment_detail CHAR(100) DEFAULT NULL,

  PRIMARY KEY(id),
  FOREIGN KEY(uid) REFERENCES user(id)

) ENGINE=InnoDB DEFAULT CHARSET=utf8;
