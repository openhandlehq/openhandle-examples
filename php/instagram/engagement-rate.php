<?php

// How do I get a competitor's Instagram engagement rate?
// https://openhandle.dev/questions/how-do-i-get-a-competitors-instagram-engagement-rate
// Run: php instagram/engagement-rate.php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use OpenHandle\OpenHandle;

$openhandle = new OpenHandle(apiKey: (string) getenv('OPENHANDLE_TEST_KEY'));

$profile = $openhandle->instagram->profile('northstar_forge_test');
$account = $profile->get()->data;
$page = $profile->posts->list();

// Skip posts with hidden likes. Null is not zero.
$posts = array_filter($page->data, static fn ($post): bool => $post->metrics->likes !== null && $post->metrics->comments !== null);
$interactions = array_sum(array_map(static fn ($post): int => (int) $post->metrics->likes + (int) $post->metrics->comments, $posts));
$rate = $interactions / count($posts) / (int) $account->metrics->followers * 100;

printf("%.2f%% over %d posts\n", $rate, count($posts));
