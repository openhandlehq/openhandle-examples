# How do I get replies or a full thread on X (Twitter)?
# https://openhandle.dev/questions/how-do-i-get-replies-or-a-full-thread-on-x
# Run: python twitter/thread_replies.py

import os

from openhandle import OpenHandle

openhandle = OpenHandle(api_key=os.environ["OPENHANDLE_TEST_KEY"])

for reply in openhandle.twitter.post("940100000000000001").comments.items():
    print(reply["author"]["handle"], reply["text"])
