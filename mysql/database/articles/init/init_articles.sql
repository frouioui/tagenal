DROP TABLE IF EXISTS article;

CREATE TABLE article (
  _id INT NOT NULL auto_increment,
  timestamp CHAR(14) DEFAULT NULL,
  id CHAR(7) DEFAULT NULL,
  aid CHAR(7) DEFAULT NULL,
  title CHAR(15) DEFAULT NULL,
  category VARBINARY(256) DEFAULT NULL,
  abstract CHAR(30) DEFAULT NULL,
  articleTags CHAR(14) DEFAULT NULL,
  authors CHAR(40) DEFAULT NULL,
  language CHAR(3) DEFAULT NULL,
  text CHAR(500) DEFAULT NULL,
  image CHAR(500) DEFAULT NULL,
  video CHAR(500) DEFAULT NULL,
  PRIMARY KEY(_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;