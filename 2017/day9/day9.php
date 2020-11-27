<?php
/**
 * http://adventofcode.com/2017/day/9
 */

// < .. > => garbage ALL characters will be ignored (even '<')
// { .. } => group (can be nested)
// ! within garbage will cancel the next character
// , separates groups


// Read the file and store in array
if ($file = fopen(__DIR__ . "/day9-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));

        if ($line === "") {
            continue;
        }

        $input = str_split($line);
        $garbage = false;
        $score = 0;
        $depth = 0;
        $garbage_count = 0;

        $count = count($input);
        for ($i = 0; $i < $count; $i++) {
            $char = $input[$i];

            if ($char === "!") {
                // Skip  the next char
                $i++;
                continue;
            }

            if ($garbage && $char !== ">") {
                $garbage_count++;
                continue;
            }

            if ($char === "<") {
                $garbage = true;
                continue;
            }

            if ($char === ">") {
                $garbage = false;
                continue;
            }

            if ($char === "{") {
                $depth++;
                $score += $depth;
                continue;
            }

            if ($char === "}") {
                $depth--;
                continue;
            }

            // And the ',' well we simply ignore it it is useless to us
        }
        echo "Total group score: " . $score . PHP_EOL;
        echo "Total garbage count: " . $garbage_count . PHP_EOL;
    }
    fclose($file);
}


