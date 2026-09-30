<?php

// How do I track brand mentions across subreddits?
// https://openhandle.dev/questions/how-do-i-track-brand-mentions-across-subreddits
// Run: php reddit/brand-mentions.php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use OpenHandle\OpenHandle;

$openhandle = new OpenHandle(apiKey: (string) getenv('OPENHANDLE_TEST_KEY'));

// Run on a schedule. Store post IDs and skip the ones you have seen.
$page = $openhandle->reddit->search->posts->list(q: 'synthetic', sort: 'new', t: 'day');

foreach ($page->data as $post) {
    echo $post->community?->displayHandle, ' ', $post->metrics->score, ' ', $post->title, PHP_EOL;
}
