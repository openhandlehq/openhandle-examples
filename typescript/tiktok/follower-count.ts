// How do I get a TikTok follower count for any username?
// https://openhandle.dev/questions/how-do-i-get-a-tiktok-follower-count-for-any-username
// Run: node tiktok/follower-count.ts (Node 24 or newer)

import { OpenHandle } from '@openhandle/sdk';

const openhandle = new OpenHandle({ apiKey: process.env.OPENHANDLE_TEST_KEY! });

const { data } = await openhandle.tiktok.profile('pixel_orchard_tt_test').get();

console.log(data.handle, data.metrics.followers, data.metrics.likes); // pixel_orchard_tt_test 75120 9814220
