#!/usr/bin/env sh
# How do I get a TikTok follower count for any username?
# https://openhandle.dev/questions/how-do-i-get-a-tiktok-follower-count-for-any-username
# Run: OPENHANDLE_TEST_KEY=oh_test_... sh tiktok/follower-count.sh
set -eu

curl "https://api.openhandle.dev/v1/tiktok/profiles/@pixel_orchard_tt_test" \
  -H "Authorization: Bearer $OPENHANDLE_TEST_KEY"
