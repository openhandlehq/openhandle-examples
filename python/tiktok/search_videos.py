# How do I search TikTok videos by keyword?
# https://openhandle.dev/questions/how-do-i-search-tiktok-videos-by-keyword
# Run: python tiktok/search_videos.py

import os

from openhandle import OpenHandle

openhandle = OpenHandle(api_key=os.environ["OPENHANDLE_TEST_KEY"])

for video in openhandle.tiktok.search.posts.items(q="synthetic"):
    print(video["author"]["handle"], video["metrics"]["views"], video["text"])
