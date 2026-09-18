# Openhandle examples

Runnable answers to the questions developers ask about public Instagram, TikTok, X, and Reddit data. Every example is one question, one request, in TypeScript, Python, Go, and cURL. The long answer for each one lives on [openhandle.dev/questions](https://openhandle.dev/questions).

Every example uses a Test account, so it runs on a free Test key and is never charged. Create a key at [app.openhandle.dev](https://app.openhandle.dev) and export it:

```sh
export OPENHANDLE_TEST_KEY=oh_test_...
```

## Run

```sh
# TypeScript, Node 24 or newer
cd typescript && npm install && node instagram/follower-count.ts

# Python 3.10 or newer
cd python && pip install -r requirements.txt && python instagram/follower_count.py

# Go 1.27 or newer
cd go && go run ./instagram/followercount

# cURL
sh curl/instagram/follower-count.sh
```

## Questions

| Question | Answer | TypeScript | Python | Go | cURL |
|---|---|---|---|---|---|
| How do I get an Instagram follower count without logging in? | [https://openhandle.dev/questions/how-do-i-get-an-instagram-follower-count-without-logging-in](https://openhandle.dev/questions/how-do-i-get-an-instagram-follower-count-without-logging-in) | [`typescript/instagram/follower-count.ts`](typescript/instagram/follower-count.ts) | [`python/instagram/follower_count.py`](python/instagram/follower_count.py) | [`go/instagram/followercount`](go/instagram/followercount/main.go) | [`curl/instagram/follower-count.sh`](curl/instagram/follower-count.sh) |
| How do I get comments on an Instagram post that is not mine? | [https://openhandle.dev/questions/how-do-i-get-comments-on-an-instagram-post-that-is-not-mine](https://openhandle.dev/questions/how-do-i-get-comments-on-an-instagram-post-that-is-not-mine) | [`typescript/instagram/post-comments.ts`](typescript/instagram/post-comments.ts) | [`python/instagram/post_comments.py`](python/instagram/post_comments.py) | [`go/instagram/postcomments`](go/instagram/postcomments/main.go) | [`curl/instagram/post-comments.sh`](curl/instagram/post-comments.sh) |
| How do I get Instagram hashtag posts past the 30 hashtag limit? | [https://openhandle.dev/questions/how-do-i-get-instagram-hashtag-posts-past-the-30-hashtag-limit](https://openhandle.dev/questions/how-do-i-get-instagram-hashtag-posts-past-the-30-hashtag-limit) | [`typescript/instagram/hashtag-posts.ts`](typescript/instagram/hashtag-posts.ts) | [`python/instagram/hashtag_posts.py`](python/instagram/hashtag_posts.py) | [`go/instagram/hashtagposts`](go/instagram/hashtagposts/main.go) | [`curl/instagram/hashtag-posts.sh`](curl/instagram/hashtag-posts.sh) |
| How do I get a competitor's Instagram engagement rate? | [https://openhandle.dev/questions/how-do-i-get-a-competitors-instagram-engagement-rate](https://openhandle.dev/questions/how-do-i-get-a-competitors-instagram-engagement-rate) | [`typescript/instagram/engagement-rate.ts`](typescript/instagram/engagement-rate.ts) | [`python/instagram/engagement_rate.py`](python/instagram/engagement_rate.py) | [`go/instagram/engagementrate`](go/instagram/engagementrate/main.go) | [`curl/instagram/engagement-rate.sh`](curl/instagram/engagement-rate.sh) |
| How do I get a TikTok follower count for any username? | [https://openhandle.dev/questions/how-do-i-get-a-tiktok-follower-count-for-any-username](https://openhandle.dev/questions/how-do-i-get-a-tiktok-follower-count-for-any-username) | [`typescript/tiktok/follower-count.ts`](typescript/tiktok/follower-count.ts) | [`python/tiktok/follower_count.py`](python/tiktok/follower_count.py) | [`go/tiktok/followercount`](go/tiktok/followercount/main.go) | [`curl/tiktok/follower-count.sh`](curl/tiktok/follower-count.sh) |
| How do I get the view count of a TikTok video? | [https://openhandle.dev/questions/how-do-i-get-the-view-count-of-a-tiktok-video](https://openhandle.dev/questions/how-do-i-get-the-view-count-of-a-tiktok-video) | [`typescript/tiktok/video-views.ts`](typescript/tiktok/video-views.ts) | [`python/tiktok/video_views.py`](python/tiktok/video_views.py) | [`go/tiktok/videoviews`](go/tiktok/videoviews/main.go) | [`curl/tiktok/video-views.sh`](curl/tiktok/video-views.sh) |
| How do I search TikTok videos by keyword? | [https://openhandle.dev/questions/how-do-i-search-tiktok-videos-by-keyword](https://openhandle.dev/questions/how-do-i-search-tiktok-videos-by-keyword) | [`typescript/tiktok/search-videos.ts`](typescript/tiktok/search-videos.ts) | [`python/tiktok/search_videos.py`](python/tiktok/search_videos.py) | [`go/tiktok/searchvideos`](go/tiktok/searchvideos/main.go) | [`curl/tiktok/search-videos.sh`](curl/tiktok/search-videos.sh) |
| How do I search X (Twitter) posts by keyword? | [https://openhandle.dev/questions/how-do-i-search-x-twitter-posts-by-keyword](https://openhandle.dev/questions/how-do-i-search-x-twitter-posts-by-keyword) | [`typescript/twitter/search-posts.ts`](typescript/twitter/search-posts.ts) | [`python/twitter/search_posts.py`](python/twitter/search_posts.py) | [`go/twitter/searchposts`](go/twitter/searchposts/main.go) | [`curl/twitter/search-posts.sh`](curl/twitter/search-posts.sh) |
| How do I get replies or a full thread on X (Twitter)? | [https://openhandle.dev/questions/how-do-i-get-replies-or-a-full-thread-on-x](https://openhandle.dev/questions/how-do-i-get-replies-or-a-full-thread-on-x) | [`typescript/twitter/thread-replies.ts`](typescript/twitter/thread-replies.ts) | [`python/twitter/thread_replies.py`](python/twitter/thread_replies.py) | [`go/twitter/threadreplies`](go/twitter/threadreplies/main.go) | [`curl/twitter/thread-replies.sh`](curl/twitter/thread-replies.sh) |
| Can I get X (Twitter) posts by date range? | [https://openhandle.dev/questions/can-i-get-x-posts-by-date-range](https://openhandle.dev/questions/can-i-get-x-posts-by-date-range) | [`typescript/twitter/posts-since.ts`](typescript/twitter/posts-since.ts) | [`python/twitter/posts_since.py`](python/twitter/posts_since.py) | [`go/twitter/postssince`](go/twitter/postssince/main.go) | [`curl/twitter/posts-since.sh`](curl/twitter/posts-since.sh) |
| How do I track brand mentions across subreddits? | [https://openhandle.dev/questions/how-do-i-track-brand-mentions-across-subreddits](https://openhandle.dev/questions/how-do-i-track-brand-mentions-across-subreddits) | [`typescript/reddit/brand-mentions.ts`](typescript/reddit/brand-mentions.ts) | [`python/reddit/brand_mentions.py`](python/reddit/brand_mentions.py) | [`go/reddit/brandmentions`](go/reddit/brandmentions/main.go) | [`curl/reddit/brand-mentions.sh`](curl/reddit/brand-mentions.sh) |
| How do I get a Reddit user's post history? | [https://openhandle.dev/questions/how-do-i-get-a-reddit-users-post-history](https://openhandle.dev/questions/how-do-i-get-a-reddit-users-post-history) | [`typescript/reddit/user-post-history.ts`](typescript/reddit/user-post-history.ts) | [`python/reddit/user_post_history.py`](python/reddit/user_post_history.py) | [`go/reddit/userposthistory`](go/reddit/userposthistory/main.go) | [`curl/reddit/user-post-history.sh`](curl/reddit/user-post-history.sh) |

## Rules the examples follow

- A number the platform hides is `null`, never `0`. The examples skip nulls instead of counting them.
- Store `capturedAt` with every number. It is when the platform showed it.
- Store the `id` from the first response. Handles change, IDs do not.
- Private accounts return `PROFILE_PRIVATE`. Deleted ones return `PROFILE_NOT_FOUND`. Store those answers, do not retry them.

This repository is synchronized from the Openhandle monorepo by `openhandle-sdk-sync[bot]`. Open issues here; changes land through the sync.

## License

MIT. See [LICENSE](LICENSE).
