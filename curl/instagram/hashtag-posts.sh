#!/usr/bin/env sh
# How do I get Instagram hashtag posts past the 30 hashtag limit?
# https://openhandle.dev/questions/how-do-i-get-instagram-hashtag-posts-past-the-30-hashtag-limit
# Run: OPENHANDLE_TEST_KEY=oh_test_... sh instagram/hashtag-posts.sh
set -eu

curl "https://api.openhandle.dev/v1/instagram/hashtags/syntheticvolume/posts?sort=recent" \
  -H "Authorization: Bearer $OPENHANDLE_TEST_KEY"
