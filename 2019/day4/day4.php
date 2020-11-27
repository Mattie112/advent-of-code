<?php
/** https://adventofcode.com/2019/day/4 */

$input_start = 156218;
$input_end = 652527;

$counter1 = 0;
$counter2 = 0;
for ($i = $input_start; $i <= $input_end; $i++) {
    if (check($i)) {
        $counter1++;
    }
    if (check2($i)) {
        $counter2++;
    }
}

echo "Part 1: " . $counter1 . PHP_EOL;
echo "Part 2: " . $counter2 . PHP_EOL;

function check(int $input): bool
{
    $number = (string)$input;
    // Assuming 6 digits
    // Assuming in range of input
    // Fist check never decreacing digits
    $digits = str_split($number);
    $prev_digit = -1;
    foreach ($digits as $digit) {
        $digit = (int)$digit;
        if ($digit < $prev_digit) {
            return false;
        }
        $prev_digit = $digit;
    }

    // Now check for at least a single double digit
    if (!preg_match("~(11|22|33|44|55|66|77|88|99|00)~", $number)) {
        return false;
    }

    return true;
}

function check2(int $input)
{
    if (!check($input)) {
        return false;
    }
    $number = (string)$input;
    // Additional check for step 2 check for triple (or more) digits
    $matches = [];
    $correct_count = 0;
    if (preg_match_all("~(11|22|33|44|55|66|77|88|99|00)~", $number, $matches)) {
        foreach ($matches[0] as $match) {
            $start = strpos($number, $match);
            $number_to_check = substr($match, 0, 1);
            $digits = str_split($number);
            if (isset($digits[$start - 1]) && $digits[$start - 1] === $number_to_check) {
                continue;
            }
            if (isset($digits[$start + 2]) && $digits[$start + 2] === $number_to_check) {
                continue;
            }
            $correct_count++;
        }
    }
    return $correct_count >= 1;
}
