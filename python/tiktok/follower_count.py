# How do I get a TikTok follower count for any username?
# https://openhandle.dev/questions/how-do-i-get-a-tiktok-follower-count-for-any-username
# Run: python tiktok/follower_count.py

import os

from openhandle import OpenHandle

openhandle = OpenHandle(api_key=os.environ["OPENHANDLE_TEST_KEY"])

response = openhandle.tiktok.profile("pixel_orchard_tt_test").get()

print(response.data["handle"], response.data["metrics"]["followers"], response.data["metrics"]["likes"])  # pixel_orchard_tt_test 75120 9814220
