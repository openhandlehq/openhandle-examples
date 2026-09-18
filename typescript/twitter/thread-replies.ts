// How do I get replies or a full thread on X (Twitter)?
// https://openhandle.dev/questions/how-do-i-get-replies-or-a-full-thread-on-x
// Run: node twitter/thread-replies.ts (Node 24 or newer)

import { OpenHandle } from '@openhandle/sdk';

const openhandle = new OpenHandle({ apiKey: process.env.OPENHANDLE_TEST_KEY! });

let page = await openhandle.twitter.post('940100000000000001').comments.list();

while (true) {
    for (const reply of page.data) console.log(reply.author?.handle, reply.text);
    const next = await page.next();
    if (!next) break;
    page = next;
}
