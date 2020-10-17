DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `_id` BIGINT NOT NULL auto_increment,
  `timestamp` CHAR(14) DEFAULT NULL,
  `id` CHAR(5) DEFAULT NULL,
  `uid` CHAR(5) DEFAULT NULL,
  `name` CHAR(9) DEFAULT NULL,
  `gender` CHAR(7) DEFAULT NULL,
  `email` CHAR(10) DEFAULT NULL,
  `phone` CHAR(10) DEFAULT NULL,
  `dept` CHAR(9) DEFAULT NULL,
  `grade` CHAR(7) DEFAULT NULL,
  `language` CHAR(3) DEFAULT NULL,
  `region` CHAR(10) DEFAULT NULL,
  `role` CHAR(6) DEFAULT NULL,
  `preferTags` CHAR(7) DEFAULT NULL,
  `obtainedCredits` CHAR(3) DEFAULT NULL,
  PRIMARY KEY(_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;