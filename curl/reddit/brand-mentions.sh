#!/usr/bin/env sh
# How do I track brand mentions across subreddits?
# https://openhandle.dev/questions/how-do-i-track-brand-mentions-across-subreddits
# Run: OPENHANDLE_TEST_KEY=oh_test_... sh reddit/brand-mentions.sh
set -eu

curl "https://api.openhandle.dev/v1/reddit/search/posts?q=synthetic&sort=new&t=day" \
  -H "Authorization: Bearer $OPENHANDLE_TEST_KEY"
