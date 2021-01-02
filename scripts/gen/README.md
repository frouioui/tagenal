# Generation of data

## How to run

Execute the following command:

```
python3 generation_data.py --number-users 5000 --number-reads 20000 --number-articles 1000
```

The data will be generated in the `article.sql`, `user.sql`, `user_read.sql`. Media and text files related to articles, will be store in the `articles` folder, where each subfolder corresponds to an article.

## Requirements

To use the data generation script you must have installed the following python packages:

- numpy
- pillow