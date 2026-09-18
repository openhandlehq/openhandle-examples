#!/usr/bin/env sh
# How do I get a Reddit user's post history?
# https://openhandle.dev/questions/how-do-i-get-a-reddit-users-post-history
# Run: OPENHANDLE_TEST_KEY=oh_test_... sh reddit/user-post-history.sh
set -eu

curl "https://api.openhandle.dev/v1/reddit/profiles/@synthetic_reddit/posts?sort=new" \
  -H "Authorization: Bearer $OPENHANDLE_TEST_KEY"
