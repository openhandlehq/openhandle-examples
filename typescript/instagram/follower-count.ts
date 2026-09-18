// How do I get an Instagram follower count without logging in?
// https://openhandle.dev/questions/how-do-i-get-an-instagram-follower-count-without-logging-in
// Run: node instagram/follower-count.ts (Node 24 or newer)

import { OpenHandle } from '@openhandle/sdk';

const openhandle = new OpenHandle({ apiKey: process.env.OPENHANDLE_TEST_KEY! });

const { data, capturedAt } = await openhandle.instagram.profile('northstar_forge_test').get();

console.log(data.handle, data.metrics.followers, capturedAt); // northstar_forge_test 48291 ...
