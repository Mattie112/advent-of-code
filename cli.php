<?php
declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use Exception;
use Symfony\Component\Console\Application;

require __DIR__ . '/vendor/autoload.php';

$app = new Application();
$app->add(new Day1());
$app->add(new Day2());
$app->add(new Day3());
$app->add(new Day4());
$app->add(new Day5());
$app->add(new Day6());
$app->add(new Day7());
$app->add(new Day8());
$app->add(new Day9());
try {
    $app->run();
} catch (Exception $e) {
    echo "Whoops!";
}
