# Tagenal

## Description of the project
<img align="right" width="100" height="100" src="./assets/img/Tsinghua_University_Logo.png">

This project is part of the **Distributed Database Systems** class of the Advanced Computer Science master degree at **Tsinghua University**.

<br>

## List of main features

- Bulk load of the User, Article, and Read tables
- Query users, articles, and users' readings
- Insert new data in the Be-Read table
- Query top 5 daily/weekly/monthly articles, with their details
- Efficient execution of the data insert, update, and queries
- Monitoring of the whole distributed system

## Generation of the database

All the necessary scripts for the databases' data generation can be found in `./scripts/gen/`.

## Data models

### user table
```sql
CREATE TABLE user (
  _id INT NOT NULL auto_increment,
  timestamp CHAR(14) DEFAULT NULL,
  id CHAR(5) DEFAULT NULL,
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
  preferTags CHAR(7) DEFAULT NULL,
  obtainedCredits CHAR(3) DEFAULT NULL,
  PRIMARY KEY(_id)
);
```

### article table
```sql
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
  text TEXT(500) DEFAULT NULL,
  image CHAR(255) DEFAULT NULL,
  video CHAR(255) DEFAULT NULL,
  PRIMARY KEY(_id)
);
```