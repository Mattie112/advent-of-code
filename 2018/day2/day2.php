<?php
/**
 * https://adventofcode.com/2018/day/2
 */

$two_letters = 0;
$three_letters = 0;
$sorted_letters = [];
$unsorted_letters = [];

if ($file = fopen(__DIR__ . "/day2-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));

        $letters = str_split($line);
        $letter_count = [];
        foreach ($letters as $letter) {
            if (!isset($letter_count[$letter])) {
                $letter_count[$letter] = 1;
            } else {
                $letter_count[$letter]++;
            }
        }

        // Array flip will 'merge' counts that already exist
        $letter_count = array_flip($letter_count);
        if (isset($letter_count[2])) {
            $two_letters++;
        }
        if (isset($letter_count[3])) {
            $three_letters++;
        }

        // Part 2
        // this was first sorted but not sure why I even dit it that way, will probaly clean this up later as it now contains some strange looking code...
        $sorted_letters[] = $letters;
    }
    fclose($file);
}

$part2 = null;

// Loop throug the array and for each element try to find an element that matches all letters (they are already sorted) but differ one letter
// When found loop through those 2 arrays again and remove the duplicate letter, then that is your answer
// Array with sorted letters
foreach ($sorted_letters as $sorted_key => $sorted_letter) {
    // Separate letters
    foreach ($sorted_letter as $letters_to_check) {
        $correct_count = 0;
        foreach ($sorted_letters as $letters) {
            foreach ($letters as $key => $letter) {
                if ($letter === $sorted_letter[$key]) {
                    $correct_count++;
                    if ($correct_count === count($sorted_letter) - 1 && implode("", $sorted_letter) !== implode("", $letters)) {
                        echo "We have found the matching string" . PHP_EOL;
                        echo "STRING1: " . implode("", $letters) . PHP_EOL;
                        echo "STRING2: " . implode("", $sorted_letter) . PHP_EOL;

                        // Now find the letter that does not match
                        foreach ($letters as $letter_key => $letter_match) {
                            if ($sorted_letter[$letter_key] !== $letter_match) {
                                // Found the incorrect letter
                                unset($letters[$letter_key]);
                                $part2 = implode("", $letters);
                                echo $part2.PHP_EOL;
                                break;
                            }
                        }

                        break 4;
                    }
                }
            }
            $correct_count = 0;
        }
    }
}

echo "Part 1: " . ($two_letters * $three_letters) . PHP_EOL;
echo "Part 2: " . $part2 . PHP_EOL;

