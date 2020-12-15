USE users;

INSERT INTO user (timestamp,uid,name,gender,email,phone,dept,grade,language,region,role,prefer_tags,obtained_credits) VALUES
  (1506328859000, '0', 'user0', 'male', 'email0', 'phone0', 'dept13', 'grade1', 'zh', 'Beijing', 'role2', 'tags24', '42'),
  (1506328859001, '1', 'user1', 'female', 'email1', 'phone1', 'dept5', 'grade1', 'en', 'Beijing', 'role2', 'tags7', '22'),
  (1506328859002, '2', 'user2', 'male', 'email2', 'phone2', 'dept4', 'grade4', 'en', 'Beijing', 'role2', 'tags46', '62'),
  (1506328859003, '3', 'user3', 'female', 'email3', 'phone3', 'dept15', 'grade4', 'zh', 'Beijing', 'role1', 'tags0', '2'),
  (1506328859004, '4', 'user4', 'male', 'email4', 'phone4', 'dept15', 'grade4', 'en', 'Hong Kong', 'role2', 'tags18', '63');

INSERT INTO user_read (timestamp,uid,aid,read_or_not,read_time_length,read_sequence,agree_or_not,comment_or_not,share_or_not,comment_detail) VALUES
  (1506332297000, 1, 4, "1", "55", "2", "0", "1", "0", "comments to this article: (88,40)"),
  (1506332307000, 5, 4, "1", "42", "0", "0", "0", "0", "comments to this article: (13,99)"),
  (1506332317000, 5, 3, "1", "98", "0", "1", "1", "0", "comments to this article: (27,12)"),
  (1506332327000, 4, 1, "1", "97", "2", "0", "0", "1", "comments to this article: (31,37)"),
  (1506332327000, 4, 3, "1", "97", "2", "0", "1", "1", "comments to this article: (31,37)"),
  (1506332337000, 5, 5, "1", "25", "3", "0", "0", "0", "comments to this article: (66,53)");
