// Can I get X (Twitter) posts by date range?
// https://openhandle.dev/questions/can-i-get-x-posts-by-date-range
// Run: node twitter/posts-since.ts (Node 24 or newer)

import { OpenHandle } from '@openhandle/sdk';

const openhandle = new OpenHandle({ apiKey: process.env.OPENHANDLE_TEST_KEY! });

// Pagination stops once posts are older than `since`. Filter the end date yourself.
const since = '2026-08-01T00:00:00Z';
const until = new Date('2026-09-01T00:00:00Z');

let page = await openhandle.twitter.profile('copperfield_lab_x_test').posts.list({ since });

while (true) {
    for (const post of page.data) {
        if (new Date(post.createdAt) < until) console.log(post.createdAt, post.text);
    }
    const next = await page.next();
    if (!next) break;
    page = next;
}
