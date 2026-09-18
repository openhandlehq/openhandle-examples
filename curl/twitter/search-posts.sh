#!/usr/bin/env sh
# How do I search X (Twitter) posts by keyword?
# https://openhandle.dev/questions/how-do-i-search-x-twitter-posts-by-keyword
# Run: OPENHANDLE_TEST_KEY=oh_test_... sh twitter/search-posts.sh
set -eu

curl "https://api.openhandle.dev/v1/twitter/search/posts?q=synthetic&sort=latest" \
  -H "Authorization: Bearer $OPENHANDLE_TEST_KEY"
