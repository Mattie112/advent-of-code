<?php
/** https://adventofcode.com/2019/day/6 */

$raw_universe = [];

if ($file = fopen(__DIR__ . "/day6-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));
        if ($line === "") {
            continue;
        }

            [$parent, $child] = explode(")", $line);
        $raw_universe[$child] = $parent;

    }
}

$orbits = 0;
foreach ($raw_universe as $child => $parent) {
    while ($parent !== "COM") {
        $parent = $raw_universe[$parent];
        $orbits++;
    }
    $orbits++;
}
echo "Part 1:". $orbits.PHP_EOL;
// make sure com is on top

//$value = getDirectOrbits($universe, "COM");

//$bla = generateUniverse($raw_universe);

function generateUniverse(array $raw_universe)
{
    $universe = [];
    $orbits = 0;
    foreach ($raw_universe as $child => $parent) {
        while ($parent !== "COM") {
            $parent = $raw_universe[$parent];
            $universe[$parent] = $raw_universe[$parent];
            $orbits++;
        }
    }
    echo $orbits . PHP_EOL;
    return $universe;
}


// docker-compose run php-cli day6/day6.php
function getDirectOrbits(array $universe, string $planet)
{
    $value = 0;

    if (!isset($universe[$planet])) {
        return 1;
    }

//    $universe = $universe[$planet];
    foreach ($universe as $child) {
        $value += getDirectOrbits($child);
    }
    return $value;
}
