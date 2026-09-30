<?php

// How do I get a Reddit user's post history?
// https://openhandle.dev/questions/how-do-i-get-a-reddit-users-post-history
// Run: php reddit/user-post-history.php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use OpenHandle\OpenHandle;

$openhandle = new OpenHandle(apiKey: (string) getenv('OPENHANDLE_TEST_KEY'));

foreach ($openhandle->reddit->profile('synthetic_reddit')->posts->items(sort: 'new') as $post) {
    echo $post->community?->displayHandle, ' ', $post->createdAt, ' ', $post->title, PHP_EOL;
}
