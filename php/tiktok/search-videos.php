<?php

// How do I search TikTok videos by keyword?
// https://openhandle.dev/questions/how-do-i-search-tiktok-videos-by-keyword
// Run: php tiktok/search-videos.php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use OpenHandle\OpenHandle;

$openhandle = new OpenHandle(apiKey: (string) getenv('OPENHANDLE_TEST_KEY'));

foreach ($openhandle->tiktok->search->posts->items(q: 'synthetic') as $video) {
    echo $video->author?->handle, ' ', $video->metrics->views, ' ', $video->text, PHP_EOL;
}
