DROP TABLE IF EXISTS users_seq;
DROP TABLE IF EXISTS user_read_seq;
DROP TABLE IF EXISTS article_seq;
DROP TABLE IF EXISTS be_read_seq;
DROP TABLE IF EXISTS popularity_seq;
DROP TABLE IF EXISTS read_stats_seq;

CREATE TABLE users_seq(id INT, next_id BIGINT, cache BIGINT, PRIMARY KEY(id)) comment 'vitess_sequence';
INSERT INTO users_seq(id, next_id, cache) VALUES(0, 1, 100);

CREATE TABLE user_read_seq(id INT, next_id BIGINT, cache BIGINT, PRIMARY KEY(id)) comment 'vitess_sequence';
INSERT INTO user_read_seq(id, next_id, cache) VALUES(0, 1, 100);

CREATE TABLE article_seq(id INT, next_id BIGINT, cache BIGINT, PRIMARY KEY(id)) comment 'vitess_sequence';
INSERT INTO article_seq(id, next_id, cache) VALUES(0, 1, 100);

CREATE TABLE be_read_seq(id INT, next_id BIGINT, cache BIGINT, PRIMARY KEY(id)) comment 'vitess_sequence';
INSERT INTO be_read_seq(id, next_id, cache) VALUES(0, 1, 100);

CREATE TABLE popularity_seq(id INT, next_id BIGINT, cache BIGINT, PRIMARY KEY(id)) comment 'vitess_sequence';
INSERT INTO popularity_seq(id, next_id, cache) VALUES(0, 1, 100);

CREATE TABLE read_stats_seq(id INT, next_id BIGINT, cache BIGINT, PRIMARY KEY(id)) comment 'vitess_sequence';
INSERT INTO read_stats_seq(id, next_id, cache) VALUES(0, 1, 100);