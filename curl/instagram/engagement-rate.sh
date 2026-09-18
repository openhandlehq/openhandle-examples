#!/usr/bin/env sh
# How do I get a competitor's Instagram engagement rate?
# https://openhandle.dev/questions/how-do-i-get-a-competitors-instagram-engagement-rate
# Run: OPENHANDLE_TEST_KEY=oh_test_... sh instagram/engagement-rate.sh
set -eu

# Two reads. Divide (likes + comments) per post by followers. Skip null likes.
curl "https://api.openhandle.dev/v1/instagram/profiles/@northstar_forge_test" \
  -H "Authorization: Bearer $OPENHANDLE_TEST_KEY"
curl "https://api.openhandle.dev/v1/instagram/profiles/@northstar_forge_test/posts" \
  -H "Authorization: Bearer $OPENHANDLE_TEST_KEY"
