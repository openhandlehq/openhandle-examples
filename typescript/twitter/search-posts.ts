// How do I search X (Twitter) posts by keyword?
// https://openhandle.dev/questions/how-do-i-search-x-twitter-posts-by-keyword
// Run: node twitter/search-posts.ts (Node 24 or newer)

import { OpenHandle } from '@openhandle/sdk';

const openhandle = new OpenHandle({ apiKey: process.env.OPENHANDLE_TEST_KEY! });

let page = await openhandle.twitter.search.posts.list({ q: 'synthetic', sort: 'latest' });

while (true) {
    for (const post of page.data) console.log(post.author?.handle, post.text);
    const next = await page.next();
    if (!next) break;
    page = next;
}
