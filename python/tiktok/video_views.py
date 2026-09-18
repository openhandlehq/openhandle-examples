# How do I get the view count of a TikTok video?
# https://openhandle.dev/questions/how-do-i-get-the-view-count-of-a-tiktok-video
# Run: python tiktok/video_views.py

import os

from openhandle import OpenHandle

openhandle = OpenHandle(api_key=os.environ["OPENHANDLE_TEST_KEY"])

response = openhandle.tiktok.post("920100000000000001").get(freshness="live")

print(response.data["metrics"]["views"], response.data["metrics"]["likes"], response.captured_at)  # 23000 1200 ...
