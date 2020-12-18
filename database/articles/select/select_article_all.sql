\! echo 'Articles'

USE articles;

SELECT * FROM article LIMIT 20;

\! echo ''
\! echo '------------------------------------'
\! echo 'Keyspace: articles, Shard: -80'

USE articles/-80;

\! echo 'article table'
SELECT * FROM article LIMIT 20;

\! echo 'be_read table'
SELECT article.id AS id, article.timestamp, article.id AS aid, be_read.reads_nb AS reads_nb, CONCAT('ruid_', article.id) AS read_uid_list, be_read.comments_nb AS comments_nb, CONCAT('cuid_', article.id) AS comment_uid_list, be_read.agrees_nb AS agrees_nb, CONCAT('auid_', article.id) AS agree_uid_list, be_read.shares_nb AS shares_nb, CONCAT('suid_', article.id) AS share_uid_list FROM article, be_read WHERE article.id=be_read.aid;

\! echo ''
\! echo '------------------------------------'
\! echo 'Keyspace: articles, Shard: 80-'

USE articles/80-;

\! echo 'article table'
SELECT * FROM article LIMIT 20;

\! echo 'be_read table'
SELECT article.id AS id, article.timestamp, article.id AS aid, be_read.reads_nb AS reads_nb, CONCAT('ruid_', article.id) AS read_uid_list, be_read.comments_nb AS comments_nb, CONCAT('cuid_', article.id) AS comment_uid_list, be_read.agrees_nb AS agrees_nb, CONCAT('auid_', article.id) AS agree_uid_list, be_read.shares_nb AS shares_nb, CONCAT('suid_', article.id) AS share_uid_list FROM article, be_read WHERE article.id=be_read.aid;