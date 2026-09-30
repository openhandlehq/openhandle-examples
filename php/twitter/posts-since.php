<?php

// Can I get X (Twitter) posts by date range?
// https://openhandle.dev/questions/can-i-get-x-posts-by-date-range
// Run: php twitter/posts-since.php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use OpenHandle\OpenHandle;

$openhandle = new OpenHandle(apiKey: (string) getenv('OPENHANDLE_TEST_KEY'));

// Pagination stops once posts are older than `since`. Filter the end date yourself.
$since = new DateTimeImmutable('2026-08-01T00:00:00Z');
$until = new DateTimeImmutable('2026-09-01T00:00:00Z');

foreach ($openhandle->twitter->profile('copperfield_lab_x_test')->posts->items(since: $since) as $post) {
    if (new DateTimeImmutable((string) $post->createdAt) < $until) {
        echo $post->createdAt, ' ', $post->text, PHP_EOL;
    }
}
