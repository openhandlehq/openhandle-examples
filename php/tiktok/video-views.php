<?php

// How do I get the view count of a TikTok video?
// https://openhandle.dev/questions/how-do-i-get-the-view-count-of-a-tiktok-video
// Run: php tiktok/video-views.php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use OpenHandle\OpenHandle;

$openhandle = new OpenHandle(apiKey: (string) getenv('OPENHANDLE_TEST_KEY'));

$response = $openhandle->tiktok->post('920100000000000001')->get(freshness: 'live');

echo $response->data->metrics->views, ' ', $response->data->metrics->likes, ' ', $response->capturedAt?->format(DATE_ATOM), PHP_EOL; // 23000 1200 ...
