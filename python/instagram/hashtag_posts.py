# How do I get Instagram hashtag posts past the 30 hashtag limit?
# https://openhandle.dev/questions/how-do-i-get-instagram-hashtag-posts-past-the-30-hashtag-limit
# Run: python instagram/hashtag_posts.py

import os

from openhandle import OpenHandle

openhandle = OpenHandle(api_key=os.environ["OPENHANDLE_TEST_KEY"])

for post in openhandle.instagram.hashtag("syntheticvolume").posts.items(sort="recent"):
    print(post["id"], post["author"]["handle"], post["metrics"]["likes"])
