CREATE TABLE users_seq(id INT, next_id BIGINT, cache BIGINT, PRIMARY KEY(id)) comment 'vitess_sequence';
INSERT INTO users_seq(id, next_id, cache) VALUES(0, 1, 100);

CREATE TABLE user_read_seq(id INT, next_id BIGINT, cache BIGINT, PRIMARY KEY(id)) comment 'vitess_sequence';
INSERT INTO user_read_seq(id, next_id, cache) VALUES(0, 1, 100);

CREATE TABLE articles_seq(id INT, next_id BIGINT, cache BIGINT, PRIMARY KEY(id)) comment 'vitess_sequence';
INSERT INTO articles_seq(id, next_id, cache) VALUES(0, 1, 100);