# How do I track brand mentions across subreddits?
# https://openhandle.dev/questions/how-do-i-track-brand-mentions-across-subreddits
# Run: python reddit/brand_mentions.py

import os

from openhandle import OpenHandle

openhandle = OpenHandle(api_key=os.environ["OPENHANDLE_TEST_KEY"])

# Run on a schedule. Store post IDs and skip the ones you have seen.
page = openhandle.reddit.search.posts.list(q="synthetic", sort="new", t="day")

for post in page.data:
    print(post["community"]["displayHandle"], post["metrics"]["score"], post["title"])
