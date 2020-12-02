<?php
declare(strict_types=1);

use mattie112\AdventOfCode\Year2020\Day1;
use mattie112\AdventOfCode\Year2020\Day2;
use Symfony\Component\Console\Application;

require __DIR__ . '/vendor/autoload.php';

$app = new Application();
$app->add(new Day1());
$app->add(new Day2());
try {
    $app->run();
} catch (Exception $e) {
    echo "Whoops!";
}
