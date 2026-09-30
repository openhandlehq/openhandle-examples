<?php

// How do I get Instagram hashtag posts past the 30 hashtag limit?
// https://openhandle.dev/questions/how-do-i-get-instagram-hashtag-posts-past-the-30-hashtag-limit
// Run: php instagram/hashtag-posts.php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use OpenHandle\OpenHandle;

$openhandle = new OpenHandle(apiKey: (string) getenv('OPENHANDLE_TEST_KEY'));

foreach ($openhandle->instagram->hashtag('syntheticvolume')->posts->items(sort: 'recent') as $post) {
    echo $post->id, ' ', $post->author?->handle, ' ', $post->metrics->likes, PHP_EOL;
}
