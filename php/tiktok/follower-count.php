<?php

// How do I get a TikTok follower count for any username?
// https://openhandle.dev/questions/how-do-i-get-a-tiktok-follower-count-for-any-username
// Run: php tiktok/follower-count.php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use OpenHandle\OpenHandle;

$openhandle = new OpenHandle(apiKey: (string) getenv('OPENHANDLE_TEST_KEY'));

$response = $openhandle->tiktok->profile('pixel_orchard_tt_test')->get();

echo $response->data->handle, ' ', $response->data->metrics->followers, ' ', $response->data->metrics->likes, PHP_EOL; // pixel_orchard_tt_test 75120 9814220
