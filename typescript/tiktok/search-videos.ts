// How do I search TikTok videos by keyword?
// https://openhandle.dev/questions/how-do-i-search-tiktok-videos-by-keyword
// Run: node tiktok/search-videos.ts (Node 24 or newer)

import { OpenHandle } from '@openhandle/sdk';

const openhandle = new OpenHandle({ apiKey: process.env.OPENHANDLE_TEST_KEY! });

let page = await openhandle.tiktok.search.posts.list({ q: 'synthetic' });

while (true) {
    for (const video of page.data) console.log(video.author?.handle, video.metrics.views, video.caption);
    const next = await page.next();
    if (!next) break;
    page = next;
}
