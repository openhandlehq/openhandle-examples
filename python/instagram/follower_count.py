# How do I get an Instagram follower count without logging in?
# https://openhandle.dev/questions/how-do-i-get-an-instagram-follower-count-without-logging-in
# Run: python instagram/follower_count.py

import os

from openhandle import OpenHandle

openhandle = OpenHandle(api_key=os.environ["OPENHANDLE_TEST_KEY"])

response = openhandle.instagram.profile("northstar_forge_test").get()

print(response.data["handle"], response.data["metrics"]["followers"], response.captured_at)  # northstar_forge_test 48291 ...
