# How do I search X (Twitter) posts by keyword?
# https://openhandle.dev/questions/how-do-i-search-x-twitter-posts-by-keyword
# Run: python twitter/search_posts.py

import os

from openhandle import OpenHandle

openhandle = OpenHandle(api_key=os.environ["OPENHANDLE_TEST_KEY"])

for post in openhandle.twitter.search.posts.items(q="synthetic", sort="latest"):
    print(post["author"]["handle"], post["text"])
