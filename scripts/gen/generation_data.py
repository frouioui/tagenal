import json
from random import random, randrange
import argparse
import numpy as np
from PIL import Image
import shutil
import os

class Generator:
    def __init__(self, args):
        self.images_path = args.images_path
        self.output_path = args.output_path

        self.number_articles = args.number_articles
        self.number_reads = args.number_reads
        self.number_users = args.number_users

        self.articles_images_max = args.articles_images_max
        self.articles_start_id = args.articles_start_id
        self.articles_nb_authors = args.articles_nb_authors
        self.articles_nb_tags = args.articles_nb_tags
        self.articles_ratio_english = args.articles_ratio_english
        self.articles_ratio_science = args.articles_ratio_science
        self.articles_videos_percentage = args.articles_videos_percentage
        self.articles_words_max = args.articles_words_max
        
        self.reads_start_id = args.reads_start_id
        self.reads_beijing_en_ratio_agree = args.reads_beijing_en_ratio_agree
        self.reads_beijing_en_ratio_comment = args.reads_beijing_en_ratio_comment
        self.reads_beijing_en_ratio_read = args.reads_beijing_en_ratio_read
        self.reads_beijing_en_ratio_share = args.reads_beijing_en_ratio_share
        self.reads_beijing_zh_ratio_agree = args.reads_beijing_zh_ratio_agree
        self.reads_beijing_zh_ratio_comment = args.reads_beijing_zh_ratio_comment
        self.reads_beijing_zh_ratio_read = args.reads_beijing_zh_ratio_read
        self.reads_beijing_zh_ratio_share = args.reads_beijing_zh_ratio_share

        self.reads_hong_kong_en_ratio_agree = args.reads_hong_kong_en_ratio_agree
        self.reads_hong_kong_en_ratio_comment = args.reads_hong_kong_en_ratio_comment
        self.reads_hong_kong_en_ratio_read = args.reads_hong_kong_en_ratio_read
        self.reads_hong_kong_en_ratio_share = args.reads_hong_kong_en_ratio_share
        self.reads_hong_kong_zh_ratio_agree = args.reads_hong_kong_zh_ratio_agree
        self.reads_hong_kong_zh_ratio_comment = args.reads_hong_kong_zh_ratio_comment
        self.reads_hong_kong_zh_ratio_read = args.reads_hong_kong_zh_ratio_read
        self.reads_hong_kong_zh_ratio_share = args.reads_hong_kong_zh_ratio_share

        self.users_start_id = args.users_start_id
        self.users_max_credits = args.users_max_credits
        self.users_nb_depts = args.users_nb_depts
        self.users_nb_roles = args.users_nb_roles
        self.users_nb_tags = args.users_nb_tags
        self.users_ratio_beijing = args.users_ratio_beijing
        self.videos_path = args.videos_path

        self.uid_region = {}
        self.aid_lang = {}
        self.p = {}
        self.p["Beijing" + "en"] = [
            self.reads_beijing_en_ratio_agree,
            self.reads_beijing_en_ratio_comment,
            self.reads_beijing_en_ratio_read,
            self.reads_beijing_en_ratio_share
        ]
        self.p["Beijing" + "zh"] = [
            self.reads_beijing_zh_ratio_agree,
            self.reads_beijing_zh_ratio_comment,
            self.reads_beijing_zh_ratio_read,
            self.reads_beijing_zh_ratio_share
        ]
        self.p["Hong Kong" + "en"] = [
            self.reads_hong_kong_en_ratio_agree,
            self.reads_hong_kong_en_ratio_comment,
            self.reads_hong_kong_en_ratio_read,
            self.reads_hong_kong_en_ratio_share 
        ]
        self.p["Hong Kong"+"zh"] = [
            self.reads_hong_kong_zh_ratio_agree,
            self.reads_hong_kong_zh_ratio_comment,
            self.reads_hong_kong_zh_ratio_read,
            self.reads_hong_kong_zh_ratio_share
        ]

    pass

class GeneratorSQL(Generator):
    def __init__(self, args):
        super().__init__(args=args)
        pass

    def gen_an_user(self, i):
        timeBegin = 1605624528
        user = {}
        user["timestamp"] = str(timeBegin + i)
        user["uid"] = str(i)
        user["name"] = "user%d" % i
        user["gender"] = "male" if random() > 0.33 else "female"
        user["email"] = "email%d" % i
        user["phone"] = "phone%d" % i
        user["dept"]  = "dept%d" % int(random() * self.users_nb_depts)
        user["grade"] = "grade%d" % int(random() * 4 + 1)
        user["language"] = "en" if random() > 0.8 else "zh"
        user["region"] = "Beijing" if random() < (self.users_ratio_beijing/100) else "Hong Kong"
        user["role"] = "role%d" % int(random() * self.users_nb_roles)
        user["prefer_tags"] = "tags%d" % int(random() * self.users_nb_tags)
        user["obtained_credits"] = str(int(random() * self.users_max_credits))
        self.uid_region[user["uid"]] = user["region"]
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


    def gen_an_article(self, i):
        print(i)
        timeBegin = 1605883728
        article = {}
        article["timestamp"] = str(timeBegin + i)
        article["aid"] = str(i)
        article["title"] = "title%d" % i
        article["category"] = "science" if random() < (self.articles_ratio_science/100) else "technology"
        article["abstract"] = "abstract of article %d" % i
        article["article_tags"] = "tags%d" % int(random() * self.articles_nb_tags)
        article["authors"]  = "author%d" % int(random() * self.articles_nb_authors)
        article["language"] = "en" if random() > (self.articles_ratio_english/100) else "zh"
        # article path
        path = './articles/article'+str(i)
        if not os.path.exists(path):
            os.makedirs(path) 
        # create text
        article["text"] = "text_a"+str(i)+'.txt'
        num = int(random()*self.articles_words_max)
        text = ['tsinghua ']*num
        f = open(path+"/text_a"+str(i)+'.txt','w+',encoding="utf8")
        f.write("".join(text))
        f.close()
        # create images
        image_num = int(random()*self.articles_images_max)+1
        image_str = ""
        for j in range(image_num):
            image_str+= 'image_a'+str(i)+'_'+str(j)+'.jpg,'
            shutil.copyfile(self.images_path+'/' + str(randrange(1,12)) + '.jpg', path + '/image_a' + str(i) + '_' + str(j) + '.jpg')
        article["image"] = image_str
        # create video
        if random() < (self.articles_videos_percentage/100):
            article["video"] = "video_a" + str(i) + '_video.flv'
            shutil.copyfile(self.videos_path+'/video'+ str(randrange(1,2)) + '.flv', path + '/video_a' + str(i) +'_video.flv')
        else:
            article["video"] = ""

        self.aid_lang[article["aid"]] = article["language"]
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

    def gen_an_read(self, i):
        timeBegin = 1606834128
        read = {}
        read["timestamp"] = str(timeBegin + i*10000)
        read["uid"] = str(int(random() * (self.number_users-1)) + self.users_start_id)
        read["aid"] = str(int(random() * (self.number_articles-1)) + self.articles_start_id)
        
        region = self.uid_region[read["uid"]]
        lang = self.aid_lang[read["aid"]]
        ps = self.p[region + lang]

        if (random() > ps[0]):
            return self.gen_an_read(i)
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

    def generate(self):
        with open("user.sql", "w+") as f:
            f.write("USE users;\n")
            f.write("INSERT INTO user (timestamp,uid,name,gender,email,phone,dept,grade,language,region,role,prefer_tags,obtained_credits) VALUES\n")
            for i in range (self.users_start_id-1, self.users_start_id + self.number_users-2):
                f.write("  " + self.gen_an_user(i+1) + ",\n")
            f.write("  " + self.gen_an_user(self.users_start_id + self.number_users-1) + ";\n")

        with open("article.sql", "w+") as f:
            f.write("USE articles;\n")
            f.write("INSERT INTO article (timestamp,aid,title,category,abstract,article_tags,authors,language,text,image,video) VALUES\n")
            for i in range(self.articles_start_id - 1, self.articles_start_id + self.number_articles-2):
                f.write("  " + self.gen_an_article(i+1) + ",\n")
            f.write("  " + self.gen_an_article(self.articles_start_id + self.number_articles-1) + ";\n")

        with open("user_read.sql", "w+") as f:
            f.write("USE users;\n")
            f.write("INSERT INTO user_read (timestamp,uid,aid,read_or_not,read_time_length,read_sequence,agree_or_not,comment_or_not,share_or_not,comment_detail) VALUES\n")
            for i in range (self.reads_start_id-1, self.reads_start_id+self.number_reads-2):
                f.write("  " + self.gen_an_read(i+1) + ",\n")
            f.write("  " + self.gen_an_read(self.reads_start_id+self.number_reads-1) + ";\n")

    pass


def get_args():
    ap = argparse.ArgumentParser()
    ap.add_argument('--output-path', required=False, default='./')
    ap.add_argument('--images-path', required=False, default='./images')
    ap.add_argument('--videos-path', required=False, default='./videos')

    apgarticle = ap.add_argument_group('articles')
    apgarticle.add_argument('--number-articles', required=True, type=int)
    apgarticle.add_argument('--articles-start-id', required=False, default=1, type=int)
    apgarticle.add_argument('--articles-images-max', required=False, default=5, type=int)
    apgarticle.add_argument('--articles-videos-percentage', required=False, default=5, type=int)
    apgarticle.add_argument('--articles-words-max', required=False, default=1000, type=int)
    apgarticle.add_argument('--articles-nb-tags', required=False, default=50, type=int)
    apgarticle.add_argument('--articles-nb-authors', required=False, default=2000, type=int)
    apgarticle.add_argument('--articles-ratio-science', required=False, default=45, type=int)
    apgarticle.add_argument('--articles-ratio-english', required=False, default=50, type=int)

    apgusers = ap.add_argument_group('users')
    apgusers.add_argument('--number-users', required=True, type=int)
    apgusers.add_argument('--users-start-id', required=False, default=1, type=int)
    apgusers.add_argument('--users-nb-depts', required=False, default=20, type=int)
    apgusers.add_argument('--users-nb-roles', required=False, default=3, type=int)
    apgusers.add_argument('--users-nb-tags', required=False, default=50, type=int)
    apgusers.add_argument('--users-max-credits', required=False, default=99, type=int)
    apgusers.add_argument('--users-ratio-beijing', required=False, default=60, type=int)

    apgreads = ap.add_argument_group('reads')
    apgreads.add_argument('--number-reads', required=True, type=int)
    apgreads.add_argument('--reads-start-id', required=False, default=1, type=int)

    apgreads.add_argument('--reads-beijing-en-ratio-read', required=False, default=0.6, type=float)
    apgreads.add_argument('--reads-beijing-en-ratio-agree', required=False, default=0.2, type=float)
    apgreads.add_argument('--reads-beijing-en-ratio-comment', required=False, default=0.2, type=float)
    apgreads.add_argument('--reads-beijing-en-ratio-share', required=False, default=0.1, type=float)
    apgreads.add_argument('--reads-beijing-zh-ratio-read', required=False, default=0.6, type=float)
    apgreads.add_argument('--reads-beijing-zh-ratio-agree', required=False, default=0.2, type=float)
    apgreads.add_argument('--reads-beijing-zh-ratio-comment', required=False, default=0.2, type=float)
    apgreads.add_argument('--reads-beijing-zh-ratio-share', required=False, default=0.1, type=float)
    
    apgreads.add_argument('--reads-hong-kong-en-ratio-read', required=False, default=0.8, type=float)
    apgreads.add_argument('--reads-hong-kong-en-ratio-agree', required=False, default=0.2, type=float)
    apgreads.add_argument('--reads-hong-kong-en-ratio-comment', required=False, default=0.2, type=float)
    apgreads.add_argument('--reads-hong-kong-en-ratio-share', required=False, default=0.1, type=float)
    apgreads.add_argument('--reads-hong-kong-zh-ratio-read', required=False, default=0.8, type=float)
    apgreads.add_argument('--reads-hong-kong-zh-ratio-agree', required=False, default=0.2, type=float)
    apgreads.add_argument('--reads-hong-kong-zh-ratio-comment', required=False, default=0.2, type=float)
    apgreads.add_argument('--reads-hong-kong-zh-ratio-share', required=False, default=0.1, type=float)


    args = ap.parse_args()

    assert os.path.isdir(args.output_path), 'Output folder not found'
    assert os.path.isdir(args.images_path), 'Images folder not found'
    assert os.path.isdir(args.videos_path), 'Videos folder not found'
    return args

if __name__ == "__main__":
    args = get_args()
    t = GeneratorSQL(args)
    t.generate()
    pass