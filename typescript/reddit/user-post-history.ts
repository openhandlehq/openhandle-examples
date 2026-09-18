// How do I get a Reddit user's post history?
// https://openhandle.dev/questions/how-do-i-get-a-reddit-users-post-history
// Run: node reddit/user-post-history.ts (Node 24 or newer)

import { OpenHandle } from '@openhandle/sdk';

const openhandle = new OpenHandle({ apiKey: process.env.OPENHANDLE_TEST_KEY! });

let page = await openhandle.reddit.profile('synthetic_reddit').posts.list({ sort: 'new' });

while (true) {
    for (const post of page.data) console.log(post.community?.displayHandle, post.createdAt, post.title);
    const next = await page.next();
    if (!next) break;
    page = next;
}
