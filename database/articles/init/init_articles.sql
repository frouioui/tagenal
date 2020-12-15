DROP TABLE IF EXISTS articles_lookup;
DROP TABLE IF EXISTS be_read;
DROP TABLE IF EXISTS popularity;
DROP TABLE IF EXISTS article;

CREATE TABLE articles_lookup (
  id INT NOT NULL,
  keyspace_id VARBINARY(128),

  PRIMARY KEY(id)

);

CREATE TABLE article (
  id BIGINT NOT NULL,
  timestamp BIGINT DEFAULT 0,
  aid CHAR(7) DEFAULT NULL,
  title CHAR(15) DEFAULT NULL,
  category VARBINARY(256) DEFAULT NULL,
  abstract CHAR(30) DEFAULT NULL,
  article_tags CHAR(14) DEFAULT NULL,
  authors CHAR(40) DEFAULT NULL,
  language CHAR(3) DEFAULT NULL,
  text TEXT(500) DEFAULT NULL,
  image CHAR(255) DEFAULT NULL,
  video CHAR(255) DEFAULT NULL,

  PRIMARY KEY(id)

) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE be_read (
  id BIGINT NOT NULL,
  timestamp BIGINT DEFAULT 0,
  aid BIGINT NOT NULL,
  reads_nb INT,
  read_uid_list CHAR(255) DEFAULT NULL,
  comments_nb INT,
  comment_uid_list CHAR(255) DEFAULT NULL,
  agrees_nb INT,
  agree_uid_list CHAR(255) DEFAULT NULL,
  shares_nb INT,
  share_uid_list CHAR(255) DEFAULT NULL,

  PRIMARY KEY(id),
  FOREIGN KEY(aid) REFERENCES article(id)

) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE popularity (
  id BIGINT NOT NULL,
  timestamp BIGINT DEFAULT 0,
  temporality ENUM('daily', 'weekly', 'monthly'),
  aid BIGINT NOT NULL,

  PRIMARY KEY(id),
  FOREIGN KEY(aid) REFERENCES article(id)

) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE read_stats (
  id BIGINT NOT NULL AUTO_INCREMENT,
  timestamp BIGINT DEFAULT 0,
  aid BIGINT NOT NULL,
  reads_nb INT DEFAULT 0,
  comments_nb INT DEFAULT 0,
  agrees_nb INT DEFAULT 0,
  shares_nb INT DEFAULT 0,

  PRIMARY KEY(id),
  FOREIGN KEY(aid) REFERENCES article(id)

) ENGINE=InnoDB DEFAULT CHARSET=utf8;