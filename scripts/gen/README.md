# Tagenal

## Generation of the datasets

### Requirements

To generate the data you must have all the required python packages:

- numpy
- pillow

Here are the command lines to download the required packages:
```shell
python3 -m pip install Pillow
python3 -m pip install numpy
```

### Generation

To generate the data we use the following command: `python3 genTable_mongoDB10G.py`. We can also use the 50Gb or 100Gb versions. The data will be generated a the root of this folder in the a `articles` folder, and 3 files: `article.dat`, `read.dat`, `user.dat`, respectively.

A JSON version of the data generation is also available by using: `python3 genTable_mongoDB10G_json.py`.
