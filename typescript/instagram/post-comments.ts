// How do I get comments on an Instagram post that is not mine?
// https://openhandle.dev/questions/how-do-i-get-comments-on-an-instagram-post-that-is-not-mine
// Run: node instagram/post-comments.ts (Node 24 or newer)

import { OpenHandle } from '@openhandle/sdk';

const openhandle = new OpenHandle({ apiKey: process.env.OPENHANDLE_TEST_KEY! });

let page = await openhandle.instagram.post('910100000001').comments.list();

while (true) {
    for (const comment of page.data) console.log(comment.author?.handle, comment.text);
    const next = await page.next();
    if (!next) break;
    page = next;
}
