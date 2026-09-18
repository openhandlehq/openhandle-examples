# How do I get a competitor's Instagram engagement rate?
# https://openhandle.dev/questions/how-do-i-get-a-competitors-instagram-engagement-rate
# Run: python instagram/engagement_rate.py

import os

from openhandle import OpenHandle

openhandle = OpenHandle(api_key=os.environ["OPENHANDLE_TEST_KEY"])

profile = openhandle.instagram.profile("northstar_forge_test")
account = profile.get().data
page = profile.posts.list()

# Skip posts with hidden likes. None is not zero.
posts = [post for post in page.data if post["metrics"]["likes"] is not None and post["metrics"]["comments"] is not None]
per_post = sum(post["metrics"]["likes"] + post["metrics"]["comments"] for post in posts) / len(posts)
rate = per_post / account["metrics"]["followers"] * 100

print(f"{rate:.2f}% over {len(posts)} posts")
