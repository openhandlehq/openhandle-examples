# How do I get comments on an Instagram post that is not mine?
# https://openhandle.dev/questions/how-do-i-get-comments-on-an-instagram-post-that-is-not-mine
# Run: python instagram/post_comments.py

import os

from openhandle import OpenHandle

openhandle = OpenHandle(api_key=os.environ["OPENHANDLE_TEST_KEY"])

# items() pages lazily, one request per page.
for comment in openhandle.instagram.post("910100000001").comments.items():
    print(comment["author"]["handle"], comment["text"])
