// How do I get the view count of a TikTok video?
// https://openhandle.dev/questions/how-do-i-get-the-view-count-of-a-tiktok-video
// Run: node tiktok/video-views.ts (Node 24 or newer)

import { OpenHandle } from '@openhandle/sdk';

const openhandle = new OpenHandle({ apiKey: process.env.OPENHANDLE_TEST_KEY! });

const { data, capturedAt } = await openhandle.tiktok.post('920100000000000001').get({ freshness: 'live' });

console.log(data.metrics.views, data.metrics.likes, capturedAt); // 23000 1200 ...
