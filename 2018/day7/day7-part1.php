<?php
/**
 * https://adventofcode.com/2018/day/7
 */

$order = "";
$tree = [];

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

// Find parents that are not anyones children
while (count($tree) > 0) {
    $not_a_child = [];
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
            if (!in_array($parent, $not_a_child, true)) {
                $not_a_child[] = $parent;
            }
        }
    }

// First letter (in the alphabet) goes first
    sort($not_a_child);
    $this_child = reset($not_a_child);
    $order .= $this_child;

// If this is the last element make sure to add all children
    if (count($tree) === 1) {
        foreach (reset($tree) as $item) {
            $order .= $item;
        }
    }

    unset($not_a_child[0]);
    unset($tree[$this_child]);
}

echo "Order (part #1): " . $order . PHP_EOL;
