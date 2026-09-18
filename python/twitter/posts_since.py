# Can I get X (Twitter) posts by date range?
# https://openhandle.dev/questions/can-i-get-x-posts-by-date-range
# Run: python twitter/posts_since.py

import os

from openhandle import OpenHandle

openhandle = OpenHandle(api_key=os.environ["OPENHANDLE_TEST_KEY"])

from datetime import datetime, timezone

# Pagination stops once posts are older than `since`. Filter the end date yourself.
since = datetime(2026, 8, 1, tzinfo=timezone.utc)
until = datetime(2026, 9, 1, tzinfo=timezone.utc)

for post in openhandle.twitter.profile("copperfield_lab_x_test").posts.items(since=since):
    if datetime.fromisoformat(post["createdAt"].replace("Z", "+00:00")) < until:
        print(post["createdAt"], post["text"])
