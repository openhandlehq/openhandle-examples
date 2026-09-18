#!/usr/bin/env sh
# How do I get replies or a full thread on X (Twitter)?
# https://openhandle.dev/questions/how-do-i-get-replies-or-a-full-thread-on-x
# Run: OPENHANDLE_TEST_KEY=oh_test_... sh twitter/thread-replies.sh
set -eu

curl "https://api.openhandle.dev/v1/twitter/posts/940100000000000001/comments" \
  -H "Authorization: Bearer $OPENHANDLE_TEST_KEY"
