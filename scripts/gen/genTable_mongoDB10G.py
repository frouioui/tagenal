import json
from random import random
import numpy as np
from PIL import Image
from shutil import copyfile
import os

USERS_NUM = 10000
ARTICLES_NUM = 10000
READS_NUM = 1000000

uid_region = {}
aid_lang = {}


# Beijing:60%   Hong Kong:40%
# en:20%    zh:80%
# 20 depts
# 3 roles
# 50 tags
# 0~99 credits

def gen_an_user (i):
    timeBegin = 1506328859000
    user = {}
    user["timestamp"] = str(timeBegin + i)
    user["id"] = 'u'+str(i)
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
    user["preferTags"] = "tags%d" % int(random() * 50)
    user["obtainedCredits"] = str(int(random() * 100))

    uid_region[user["uid"]] = user["region"]
    return user

# science:45%   technology:55%
# en:50%    zh:50%
# 50 tags
# 2000 authors
def gen_an_article (i):
    timeBegin = 1506000000000
    article = {}
    article["id"] = 'a'+str(i)
    article["timestamp"] = str(timeBegin + i)
    article["aid"] = str(i)
    article["title"] = "title%d" % i
    article["category"] = "science" if random() > 0.55 else "technology"
    article["abstract"] = "abstract of article %d" % i
    article["articleTags"] = "tags%d" % int(random() * 50)
    article["authors"]  = "author%d" % int(random() * 2000)
    article["language"] = "en" if random() > 0.5 else "zh"

    # create text
    article["text"] = "text_a"+str(i)+'.txt'
    path = './articles/article'+str(i)
    if not os.path.exists(path):
        os.makedirs(path) 
    num = int(random()*1000)
    text = ['Tsinghua']*num
    f = open(path+"/text_a"+str(i)+'.txt','w+',encoding="utf8")
    f.write("".join(text))
    f.close()

    # create images
    image_num = int(random()*5)+1
    image_str = ""
    for j in range(image_num):
        image_str+= 'image_a'+str(i)+'_'+str(j)+'.jpg,'
    article["image"] = image_str
    num_image = int(random()*5)+1
    for j in range(num_image):
        a = np.random.randint(0,255,(360,480,3))
        img = Image.fromarray(a.astype('uint8')).convert('RGB')
        img.save(path+'/image_a'+str(i)+'_'+str(j)+'.jpg')

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
    return article

# user in Beijing read/agree/comment/share an english article with the probability 0.6/0.2/0.2/0.1
# user in Hong Kong read/agree/comment/share an Chinese article with the probability 0.8/0.2/0.2/0.1
p = {}
p["Beijing"+"en"] = [0.6,0.2,0.2,0.1]
p["Beijing"+"zh"] = [1,0.3,0.3,0.2]
p["Hong Kong"+"en"] = [1,0.3,0.3,0.2]
p["Hong Kong"+"zh"] = [0.8,0.2,0.2,0.1]
def gen_an_read (i):
    timeBegin = 1506332297000
    read = {}
    read["timestamp"] = str(timeBegin + i*10000)
    read["id"] = 'r'+str(i)
    read["uid"] = str(int(random() * USERS_NUM))
    read["aid"] = str(int(random() * ARTICLES_NUM))
    
    region = uid_region[read["uid"]]
    lang = aid_lang[read["aid"]]
    ps = p[region + lang]

    if (random() > ps[0]):
        # read["readOrNot"] = "0";
        return gen_an_read (i)
    else:
        read["readOrNot"] = "1"
        read["readTimeLength"] = str(int(random() * 100))
        read["readSequence"] = str(int(random() * 4))
        read["agreeOrNot"] = "1" if random() < ps[1] else "0"
        read["commentOrNot"] = "1" if random() < ps[2] else "0"
        read["shareOrNot"] = "1" if random() < ps[3] else "0"
        read["commentDetail"] = "comments to this article: (" + read["uid"] + "," + read["aid"] + ")" 
    return read

with open("user.dat", "w+") as f:
    for i in range (USERS_NUM):
        json.dump(gen_an_user(i), f)
        f.write("\n")

if not os.path.exists('./articles'):
    os.makedirs('./articles')
with open("article.dat", "w+") as f:
    for i in range(ARTICLES_NUM):
        json.dump(gen_an_article(i), f)
        f.write("\n")


with open("read.dat", "w+") as f:
    for i in range(READS_NUM):
        json.dump(gen_an_read(i), f)
        f.write("\n")

