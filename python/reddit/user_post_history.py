# How do I get a Reddit user's post history?
# https://openhandle.dev/questions/how-do-i-get-a-reddit-users-post-history
# Run: python reddit/user_post_history.py

import os

from openhandle import OpenHandle

openhandle = OpenHandle(api_key=os.environ["OPENHANDLE_TEST_KEY"])

for post in openhandle.reddit.profile("synthetic_reddit").posts.items(sort="new"):
    print(post["community"]["displayHandle"], post["createdAt"], post["title"])
