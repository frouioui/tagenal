CREATE TABLE users_seq(id INT, next_id BIGINT, cache BIGINT, PRIMARY KEY(id)) comment "vitess_sequence";
INSERT INTO users_seq(id, next_id, cache) VALUES(0, 100, 100);