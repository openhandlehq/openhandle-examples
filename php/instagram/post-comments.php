<?php

// How do I get comments on an Instagram post that is not mine?
// https://openhandle.dev/questions/how-do-i-get-comments-on-an-instagram-post-that-is-not-mine
// Run: php instagram/post-comments.php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use OpenHandle\OpenHandle;

$openhandle = new OpenHandle(apiKey: (string) getenv('OPENHANDLE_TEST_KEY'));

// items() pages lazily, one request per page.
foreach ($openhandle->instagram->post('910100000001')->comments->items() as $comment) {
    echo $comment->author?->handle, ' ', $comment->text, PHP_EOL;
}
