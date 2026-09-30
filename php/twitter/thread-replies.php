<?php

// How do I get replies or a full thread on X (Twitter)?
// https://openhandle.dev/questions/how-do-i-get-replies-or-a-full-thread-on-x
// Run: php twitter/thread-replies.php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use OpenHandle\OpenHandle;

$openhandle = new OpenHandle(apiKey: (string) getenv('OPENHANDLE_TEST_KEY'));

foreach ($openhandle->twitter->post('940100000000000001')->comments->items() as $reply) {
    echo $reply->author?->handle, ' ', $reply->text, PHP_EOL;
}
