<?php

// How do I search X (Twitter) posts by keyword?
// https://openhandle.dev/questions/how-do-i-search-x-twitter-posts-by-keyword
// Run: php twitter/search-posts.php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use OpenHandle\OpenHandle;

$openhandle = new OpenHandle(apiKey: (string) getenv('OPENHANDLE_TEST_KEY'));

foreach ($openhandle->twitter->search->posts->items(q: 'synthetic', sort: 'latest') as $post) {
    echo $post->author?->handle, ' ', $post->text, PHP_EOL;
}
