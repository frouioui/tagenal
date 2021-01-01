import json
from random import random,randrange
import numpy as np
from PIL import Image
from shutil import copyfile
import os

USERS_NUM = 5500
ARTICLES_NUM = 4000
READS_NUM = 20000

uid_region = {}
aid_lang = {}

# Beijing:60%   Hong Kong:40%
# en:20%    zh:80%
# 20 depts
# 3 roles
# 50 tags
# 0~99 credits

def gen_an_user (i):
    timeBegin = 1605624528
    user = {}
    user["timestamp"] = str(timeBegin + i)
    user["uid"] = str(i)
    user["name"] = "user%d" % i
    user["gender"] = "male" if random() > 0.33 else "female"
    user["email"] = "email%d" % i
    user["phone"] = "phone%d" % i
    user["dept"]  = "dept%d" % int(random() * 20)
    user["grade"] = "grade%d" % int(random() * 4 + 1)
    user["language"] = "en" if random() > 0.8 else "zh"
    user["region"] = "Beijing" if random() > 0.4 else "Hong Kong"
    user["role"] = "role%d" % int(random() * 3)
    user["prefer_tags"] = "tags%d" % int(random() * 50)
    user["obtained_credits"] = str(int(random() * 100))

    uid_region[user["uid"]] = user["region"]
    return "(" +  \
              user["timestamp"] + ", " + \
            "\"" + user["uid"] + "\", " + \
            "\"" + user["name"] + "\", " + \
            "\"" + user["gender"] + "\", " + \
            "\"" + user["email"] + "\", " + \
            "\"" + user["phone"] + "\", " + \
            "\"" + user["dept"] + "\", " + \
            "\"" + user["grade"] + "\", " + \
            "\"" + user["language"] + "\", " + \
            "\"" + user["region"] + "\", " + \
            "\"" + user["role"] + "\", " + \
            "\"" + user["prefer_tags"] + "\", " + \
            "\"" + user["obtained_credits"] + "\")";

# science:45%   technology:55%
# en:50%    zh:50%
# 50 tags
# 2000 authors
def gen_an_article (i):
    timeBegin = 1605883728
    article = {}
    article["timestamp"] = str(timeBegin + i)
    article["aid"] = str(i)
    article["title"] = "title%d" % i
    article["category"] = "science" if random() > 0.55 else "technology"
    article["abstract"] = "abstract of article %d" % i
    article["article_tags"] = "tags%d" % int(random() * 50)
    article["authors"]  = "author%d" % int(random() * 2000)
    article["language"] = "en" if random() > 0.5 else "zh"
    # create text
    article["text"] = "text_a"+str(i)+'.txt'
    path = './articles/article'+str(i)
    if not os.path.exists(path):
        os.makedirs(path) 
    num = int(random()*1000)
    text = ['tsinghua ']*num
    f = open(path+"/text_a"+str(i)+'.txt','w+',encoding="utf8")
    f.write("".join(text))
    f.close()

    # create images
    image_num = int(random()*5)+1
    image_str = ""
    for j in range(image_num):
        a = np.random.randint(0,255,(360,480,3))
        # img = Image.fromarray(a.astype('uint8')).convert('RGB')
        image_str+= 'image_a'+str(i)+'_'+str(j)+'.jpg,'
        # img.save(path + '/image_a' + str(i) + '_' + str(j) + '.jpg')
        copyfile('./images/' + str(randrange(1,12)) + '.jpg', path + '/image_a' + str(i) + '_' + str(j) + '.jpg')

    article["image"] = image_str

    # create video
    if random() < 0.05:
        #has one video
        article["video"] = "video_a"+str(i)+'_video.flv'
        if random()<0.5:
            copyfile('./video/video1.flv',path+"/video_a"+str(i)+'_video.flv')
        else:
            copyfile('./video/video2.flv',path+"/video_a"+str(i)+'_video.flv')
    else:
        article["video"] = ""

    aid_lang[article["aid"]] = article["language"]
    return "(" +  \
            " " + article["timestamp"] + " , " + \
            "\"" + article["aid"] + "\", " + \
            "\"" + article["title"] + "\", " + \
            "\"" + article["category"] + "\", " + \
            "\"" + article["abstract"] + "\", " + \
            "\"" + article["article_tags"] + "\", " + \
            "\"" + article["authors"] + "\", " + \
            "\"" + article["language"] + "\", " + \
            "\"" + article["text"] + "\", " + \
            "\"" + article["image"] + "\", " + \
            "\"" + article["video"] + "\")";

# user in Beijing read/agree/comment/share an english article with the probability 0.6/0.2/0.2/0.1
# user in Hong Kong read/agree/comment/share an Chinese article with the probability 0.8/0.2/0.2/0.1
p = {}
p["Beijing"+"en"] = [0.6,0.2,0.2,0.1]
p["Beijing"+"zh"] = [1,0.3,0.3,0.2]
p["Hong Kong"+"en"] = [1,0.3,0.3,0.2]
p["Hong Kong"+"zh"] = [0.8,0.2,0.2,0.1]
def gen_an_read (i):
    timeBegin = 1606834128
    read = {}
    read["timestamp"] = str(timeBegin + i*10000)
    read["uid"] = str(int(random() * (USERS_NUM-1)) + 1)
    read["aid"] = str(int(random() * (ARTICLES_NUM-1)) + 1)
    
    region = uid_region[read["uid"]]
    lang = aid_lang[read["aid"]]
    ps = p[region + lang]

    if (random() > ps[0]):
        # read["readOrNot"] = "0";
        return gen_an_read (i)
    else:
        read["read_or_not"] = "1"
        read["read_time_length"] = str(int(random() * 100))
        read["read_sequence"] = str(int(random() * 4))
        read["agree_or_not"] = "1" if random() < ps[1] else "0"
        read["comment_or_not"] = "1" if random() < ps[2] else "0"
        read["share_or_not"] = "1" if random() < ps[3] else "0"
        read["comment_detail"] = "comments to this article: (" + read["uid"] + "," + read["aid"] + ")" 
    return "(" +  \
            " " + read["timestamp"] + " , " + \
            " " + read["uid"] + " , " + \
            " " + read["aid"] + " , " + \
            "\"" + read["read_or_not"] + "\", " + \
            "\"" + read["read_time_length"] + "\", " + \
            "\"" + read["read_sequence"] + "\", " + \
            "\"" + read["agree_or_not"] + "\", " + \
            "\"" + read["comment_or_not"] + "\", " + \
            "\"" + read["share_or_not"] + "\", " + \
            "\"" + read["comment_detail"] + "\")"

with open("user.sql", "w+") as f:
    f.write("USE users;\n")
    f.write("INSERT INTO user (timestamp,uid,name,gender,email,phone,dept,grade,language,region,role,prefer_tags,obtained_credits) VALUES\n")
    for i in range (USERS_NUM-1):
        f.write("  " + gen_an_user(i+1) + ",\n")
    f.write("  " + gen_an_user(USERS_NUM) + ";\n")

with open("article.sql", "w+") as f:
    f.write("USE articles;\n")
    f.write("INSERT INTO article (timestamp,aid,title,category,abstract,article_tags,authors,language,text,image,video) VALUES\n")
    for i in range (ARTICLES_NUM-1):
        f.write("  " + gen_an_article(i+1) + ",\n")
    f.write("  " + gen_an_article(ARTICLES_NUM-1) + ";\n")

with open("user_read.sql", "w+") as f:
    f.write("USE users;\n")
    f.write("INSERT INTO user_read (timestamp,uid,aid,read_or_not,read_time_length,read_sequence,agree_or_not,comment_or_not,share_or_not,comment_detail) VALUES\n")
    for i in range (READS_NUM-1):
        f.write("  " + gen_an_read(i+1) + ",\n")
    f.write("  " + gen_an_read(READS_NUM) + ";\n")

