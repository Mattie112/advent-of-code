<?php
/**
 * https://adventofcode.com/2018/day/16
 */
const OPCODES = [
    'addr',
    'addi',
    'mulr',
    'muli',
    'banr',
    'bani',
    'borr',
    'bori',
    'setr',
    'seti',
    'gtir',
    'gtri',
    'gtrr',
    'eqir',
    'eqri',
    'eqrr',
];

$tmp_example = [];
$examples = [];
$program = []; // part 2

if ($file = fopen(__DIR__ . "/day16-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));
        if ($line === "") {
            continue;
        }

        if (preg_match("@Before: \[(\d+), (\d+), (\d+), (\d+)\]@", $line, $matches)) {
            [, $opcode, $a, $b, $c] = $matches;
            $tmp_example['before'] = [(int) $opcode, (int) $a, (int) $b, (int) $c];
        }

        if (preg_match("@After:  \[(\d+), (\d+), (\d+), (\d+)\]@", $line, $matches)) {
            [, $opcode, $a, $b, $c] = $matches;
            $tmp_example['after'] = [(int) $opcode, (int) $a, (int) $b, (int) $c];
            // We kinda assume we always get the input in the correct format
            $examples[] = $tmp_example;
            unset($tmp_example);
        }

        if (preg_match("@(\d+) (\d+) (\d+) (\d+)@", $line, $matches)) {
            [, $opcode, $a, $b, $c] = $matches;
            // Detect if we are working on the examples or on the program (part 2)
            if (isset($tmp_example["before"])) {
                $tmp_example['instruction'] = [(int) $opcode, (int) $a, (int) $b, (int) $c];
            } else {
                // part 2
                $program[] = [(int) $opcode, (int) $a, (int) $b, (int) $c];
            }
        }
    }
}

// So we don't get notices etc
$opcode_lookup = array_fill(0, 16, OPCODES);
$matches_to_count = 0;
foreach ($examples as $example) {
    [$opcode, $a, $b, $c] = $example["instruction"];
    $expected = implode("", $example["after"]);
    $matches = 0;
    $matching_opcodes = [];
    foreach (OPCODES as $tmp_opcode) {
        $value = executeOpCode($tmp_opcode, $a, $b, $c, $example["before"]);
        if ($expected === implode("", $value)) {
            $matches++;
            $matching_opcodes[] = $tmp_opcode;
        }
    }
    if ($matches >= 3) {
        $matches_to_count++;
    }

    // With this we have possible opcode per ID number (still could be multiple possible)
    $opcode_lookup[$opcode] = array_intersect($opcode_lookup[$opcode], $matching_opcodes);
}
echo "Part #1: " . $matches_to_count . " examples behave like 3 or more opcodes" . PHP_EOL;

// Now try to find the opcodes by looking for everything with count = 1 and then removing that opcode from all other arrays
$opcode_result = [];
//$tmp_lookup_opcode = $opcode_lookup;
while (count($opcode_result) < 16) {
    foreach ($opcode_lookup as $i => $opcodes) {
        if (count($opcodes) === 1) {
            // Yes we know for sure this ID belongs to this opcode
            $opcode = reset($opcodes);
            $opcode_result[$i] = $opcode;
            // Now remove this opcode from all lookup candidates
            foreach ($opcode_lookup as $j => $tmp_opcode) {
                $opcode_lookup[$j] = array_filter($tmp_opcode, function ($value) use ($opcode) {
                    return $value !== $opcode;
                });
            }
        }
    }
}

// Now we have the ocrrect opcodes! Let's execut the program from step 1
ksort($opcode_result);

$registers = [0, 0, 0, 0];
foreach ($program as $p) {
    [$opcode, $a, $b, $c] = $p;
    $opcode = $opcode_result[$opcode];
    $registers = executeOpCode($opcode, $a, $b, $c, $registers);
}

echo "Part #2, the register with ID 0 contains: " . $registers[0] . PHP_EOL;

function executeOpCode($opcode, $a, $b, $c, $registers)
{
    switch ($opcode) {
        case "addr":
            $registers[$c] = $registers[$a] + $registers[$b];
            break;
        case "addi":
            $registers[$c] = $registers[$a] + $b;
            break;
        case"mulr":
            $registers[$c] = $registers[$a] * $registers[$b];
            break;
        case "muli":
            $registers[$c] = $registers[$a] * $b;
            break;
        case "banr":
            $registers[$c] = $registers[$a] & $registers[$b];
            break;
        case "bani":
            $registers[$c] = $registers[$a] & $b;
            break;
        case "borr":
            $registers[$c] = $registers[$a] | $registers[$b];
            break;
        case "bori":
            $registers[$c] = $registers[$a] | $b;
            break;
        case "setr":
            $registers[$c] = $registers[$a];
            break;
        case "seti":
            $registers[$c] = $a;
            break;
        case "gtir":
            $registers[$c] = ($a > $registers[$b]) ? 1 : 0;
            break;
        case 'gtri':
            $registers[$c] = ($registers[$a] > $b) ? 1 : 0;
            break;
        case"gtrr":
            $registers[$c] = ($registers[$a] > $registers[$b]) ? 1 : 0;
            break;
        case "eqir":
            $registers[$c] = ($a === $registers[$b]) ? 1 : 0;
            break;
        case "eqri":
            $registers[$c] = ($registers[$a] === $b) ? 1 : 0;
            break;
        case"eqrr":
            $registers[$c] = ($registers[$a] === $registers[$b]) ? 1 : 0;
            break;
    }

    return $registers;
}

