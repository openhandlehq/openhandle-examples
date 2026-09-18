#!/usr/bin/env sh
# How do I get an Instagram follower count without logging in?
# https://openhandle.dev/questions/how-do-i-get-an-instagram-follower-count-without-logging-in
# Run: OPENHANDLE_TEST_KEY=oh_test_... sh instagram/follower-count.sh
set -eu

curl "https://api.openhandle.dev/v1/instagram/profiles/@northstar_forge_test" \
  -H "Authorization: Bearer $OPENHANDLE_TEST_KEY"
