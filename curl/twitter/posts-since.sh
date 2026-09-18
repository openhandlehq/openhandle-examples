#!/usr/bin/env sh
# Can I get X (Twitter) posts by date range?
# https://openhandle.dev/questions/can-i-get-x-posts-by-date-range
# Run: OPENHANDLE_TEST_KEY=oh_test_... sh twitter/posts-since.sh
set -eu

# Pagination stops once posts are older than since.
curl "https://api.openhandle.dev/v1/twitter/profiles/@copperfield_lab_x_test/posts?since=2026-08-01T00:00:00Z" \
  -H "Authorization: Bearer $OPENHANDLE_TEST_KEY"
