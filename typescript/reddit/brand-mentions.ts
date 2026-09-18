// How do I track brand mentions across subreddits?
// https://openhandle.dev/questions/how-do-i-track-brand-mentions-across-subreddits
// Run: node reddit/brand-mentions.ts (Node 24 or newer)

import { OpenHandle } from '@openhandle/sdk';

const openhandle = new OpenHandle({ apiKey: process.env.OPENHANDLE_TEST_KEY! });

// Run on a schedule. Store post IDs and skip the ones you have seen.
const page = await openhandle.reddit.search.posts.list({ q: 'synthetic', sort: 'new', t: 'day' });

for (const post of page.data) console.log(post.community?.displayHandle, post.metrics.score, post.title);
