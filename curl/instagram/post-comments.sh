#!/usr/bin/env sh
# How do I get comments on an Instagram post that is not mine?
# https://openhandle.dev/questions/how-do-i-get-comments-on-an-instagram-post-that-is-not-mine
# Run: OPENHANDLE_TEST_KEY=oh_test_... sh instagram/post-comments.sh
set -eu

# Pass meta.cursors.next back as ?cursor= for the next page.
curl "https://api.openhandle.dev/v1/instagram/posts/910100000001/comments" \
  -H "Authorization: Bearer $OPENHANDLE_TEST_KEY"
