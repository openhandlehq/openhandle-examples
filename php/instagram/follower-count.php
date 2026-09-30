<?php

// How do I get an Instagram follower count without logging in?
// https://openhandle.dev/questions/how-do-i-get-an-instagram-follower-count-without-logging-in
// Run: php instagram/follower-count.php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use OpenHandle\OpenHandle;

$openhandle = new OpenHandle(apiKey: (string) getenv('OPENHANDLE_TEST_KEY'));

$response = $openhandle->instagram->profile('northstar_forge_test')->get();

echo $response->data->handle, ' ', $response->data->metrics->followers, ' ', $response->capturedAt?->format(DATE_ATOM), PHP_EOL; // northstar_forge_test 48291 ...
