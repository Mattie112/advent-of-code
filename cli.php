<?php
declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use Exception;
use Symfony\Component\Console\Application;
use Symfony\Component\Finder\Finder;

require __DIR__ . '/vendor/autoload.php';

$app = new Application();

// I am kinda lazy so I automatically all `Day` commands by checking the entire directory and using the fact that the class name equals the file name ;)
$finder = new Finder();
$finder->files()->in(__DIR__ . "/src/Year2020");
if ($finder->hasResults()) {
    foreach ($finder as $file) {
        $getFilenameWithoutExtension = "mattie112\\AdventOfCode\\Year2020\\" . $file->getFilenameWithoutExtension();
        $app->add(new $getFilenameWithoutExtension());
    }
}

try {
    $app->run();
} catch (Exception $e) {
    echo "Whoops!";
}
