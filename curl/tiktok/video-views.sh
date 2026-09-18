#!/usr/bin/env sh
# How do I get the view count of a TikTok video?
# https://openhandle.dev/questions/how-do-i-get-the-view-count-of-a-tiktok-video
# Run: OPENHANDLE_TEST_KEY=oh_test_... sh tiktok/video-views.sh
set -eu

curl "https://api.openhandle.dev/v1/tiktok/posts/920100000000000001?freshness=live" \
  -H "Authorization: Bearer $OPENHANDLE_TEST_KEY"
