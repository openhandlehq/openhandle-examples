// How do I get Instagram hashtag posts past the 30 hashtag limit?
// https://openhandle.dev/questions/how-do-i-get-instagram-hashtag-posts-past-the-30-hashtag-limit
// Run: node instagram/hashtag-posts.ts (Node 24 or newer)

import { OpenHandle } from '@openhandle/sdk';

const openhandle = new OpenHandle({ apiKey: process.env.OPENHANDLE_TEST_KEY! });

let page = await openhandle.instagram.hashtag('syntheticvolume').posts.list({ sort: 'recent' });

while (true) {
    for (const post of page.data) console.log(post.id, post.author?.handle, post.metrics.likes);
    const next = await page.next();
    if (!next) break;
    page = next;
}
