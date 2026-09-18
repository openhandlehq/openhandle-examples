#!/usr/bin/env sh
# How do I search TikTok videos by keyword?
# https://openhandle.dev/questions/how-do-i-search-tiktok-videos-by-keyword
# Run: OPENHANDLE_TEST_KEY=oh_test_... sh tiktok/search-videos.sh
set -eu

curl "https://api.openhandle.dev/v1/tiktok/search/posts?q=synthetic" \
  -H "Authorization: Bearer $OPENHANDLE_TEST_KEY"
