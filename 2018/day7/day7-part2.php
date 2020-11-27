<?php
/**
 * https://adventofcode.com/2018/day/7
 */

$order = "";
$tree = [];
// For part #2
$letters = array_merge([0 => 0], range('A', 'Z'));
$letters = array_flip($letters);
unset($letters[0]);
// Setup part #2
$letter_time = 60;
$worker_amount = 5;

if ($file = fopen(__DIR__ . "/day7-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));

        if (preg_match("@Step (.) .* step (.)@", $line, $matches)) {
            $parent = $matches[1];
            $child = $matches[2];

            $tree[$parent][] = $child;
        }
    }
    fclose($file);
}

// False = idle
// True = working
// int = seconds busy
$workers = [];
for ($i = 1; $i <= $worker_amount; $i++) {
    $workers[$i] = [false, 0, ""];
}
$total_seconds = 0;

$todo = [];
$in_progress = [];
// Find parents that are not anyones children
while (count($tree) > 0) {
    // Only count the time once no matter how many workers
    $counted = false;
    foreach ($workers as $id => $elem) {
        if ($workers[$id][0] === true) {
            if (!$counted) {
                // Only add work time for a single worker as it does not matter how many are working in paralel
                $total_seconds++;
                $counted = true;
            }
            $workers[$id][1]--;
            if ($workers[$id][1] === 0) {
                $workers[$id][0] = false;
                echo "Worker " . $id . " is free again (finished " . $workers[$id][2] . ")" . PHP_EOL;
                $order .= $workers[$id][2];

                // When done don't forget to wait for the last job to finish
                if (count($tree) === 1) {
                    foreach (reset($tree) as $item) {
                        $order .= $item;
                        $total_seconds += $letter_time + $letters[$item];
                    }
                }
                echo $order . PHP_EOL;
                unset($tree[$workers[$id][2]]);
            }
        }
    }

    // Only continue when there are free workers
    $free_worker = false;
    foreach ($workers as $id => [$status, $seconds, $letter]) {
        if ($status === false) {
            $free_worker = true;
            break;
        }
    }

    if (!$free_worker) {
        continue;
    }

    $not_children = [];
    foreach ($tree as $parent => $children) {
        $is_child = false;
        foreach ($tree as $subparent => $subchildren) {
            if ($subparent === $parent) {
                continue;
            }

            foreach ($subchildren as $subchild) {
                if ($subchild === $parent) {
                    $is_child = true;
                    break;
                }
            }
        }
        if (!$is_child) {
//            echo $parent . " is not a child" . PHP_EOL;
            if (!in_array($parent, $not_children, true)) {
                $not_children[] = $parent;
            }
        }
    }

// First letter (in the alphabet) goes first
    sort($not_children);
    foreach ($not_children as $elem) {
        if (!isset($todo[$elem])) {
            $todo[$elem] = true;
        }
    }

    foreach ($todo as $not_a_child => $_tmp) {
        if (!isset($in_progress[$not_a_child])) {
            foreach ($workers as $id => [$status, $seconds]) {
                if ($status === false) {
                    echo "Assigning job " . $not_a_child . " to worker " . $id . " for " . ($letters[$not_a_child] + $letter_time) . " seconds" . PHP_EOL;
                    // We have a free worker!!! yeah! let's assign a job!
                    $workers[$id] = [true, $letters[$not_a_child] + $letter_time, $not_a_child,];
                    $todo = array_diff($todo, [$not_a_child]);
                    $in_progress[$not_a_child] = true;
                    break;
                }
            }
        }
    }
}

echo "Order (part #2): " . $order . PHP_EOL;
echo $total_seconds . PHP_EOL;
