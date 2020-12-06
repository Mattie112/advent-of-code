<?php
declare(strict_types=1);

use mattie112\AdventOfCode\Year2020\Day1;
use mattie112\AdventOfCode\Year2020\Day2;
use mattie112\AdventOfCode\Year2020\Day3;
use mattie112\AdventOfCode\Year2020\Day4;
use mattie112\AdventOfCode\Year2020\Day5;
use mattie112\AdventOfCode\Year2020\Day6;
use Symfony\Component\Console\Application;

require __DIR__ . '/vendor/autoload.php';

$app = new Application();
$app->add(new Day1());
$app->add(new Day2());
$app->add(new Day3());
$app->add(new Day4());
$app->add(new Day5());
$app->add(new Day6());
try {
    $app->run();
} catch (Exception $e) {
    echo "Whoops!";
}
