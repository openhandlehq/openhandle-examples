// How do I get a competitor's Instagram engagement rate?
// https://openhandle.dev/questions/how-do-i-get-a-competitors-instagram-engagement-rate
// Run: node instagram/engagement-rate.ts (Node 24 or newer)

import { OpenHandle } from '@openhandle/sdk';

const openhandle = new OpenHandle({ apiKey: process.env.OPENHANDLE_TEST_KEY! });

const profile = openhandle.instagram.profile('northstar_forge_test');
const [{ data: account }, page] = await Promise.all([profile.get(), profile.posts.list()]);

// Skip posts with hidden likes. Null is not zero.
const posts = page.data.filter(post => post.metrics.likes !== null && post.metrics.comments !== null);
const perPost = posts.reduce((sum, post) => sum + post.metrics.likes! + post.metrics.comments!, 0) / posts.length;
const rate = (perPost / account.metrics.followers!) * 100;

console.log(`${rate.toFixed(2)}% over ${posts.length} posts`);
