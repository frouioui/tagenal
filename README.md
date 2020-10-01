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

**User collection:**
```json
{
    "_id": "objectID",
    "timestamp": "1506328859000", 
    "id": "u0",
    "uid": "0",
    "name": "user0",
    "gender": "male",
    "email": "email0",
    "phone": "phone0",
    "dept": "dept3",
    "grade": "grade4",
    "language": "en",
    "region": "Beijing",
    "role": "role1",
    "preferTags": "tags24",
    "obtainedCredits": "84",
}
```

**Article collection:**
```json
{
    "_id": "objectID",
    "id": "a0",
    "timestamp": "1506000000000",
    "aid": "0",
    "title": "title0",
    "category": "technology",
    "abstract": "abstract of article 0",
    "articleTags": "tags9",
    "authors": "author1407",
    "language": "en",
    "text": "text_a0.txt",
    "image": "image_a0_0.jpg,image_a0_1.jpg,",
    "video": ""
}
```

**Read collection:**
```json
{
    "_id": "objectID",
    "timestamp": "1506332297000",
    "id": "r0",
    "uid": "7312",
    "aid": "9630",
    "readOrNot": "1",
    "readTimeLength": "99",
    "readSequence": "3",
    "agreeOrNot": "0",
    "commentOrNot": "0",
    "shareOrNot": "0",
    "commentDetail": "comments to this article: (7312,9630)"
}
```