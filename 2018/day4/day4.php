<?php
/**
 * https://adventofcode.com/2018/day/4
 */

const BEGIN_SHIFT = 1;
const FALL_ASLEEP = 2;
const WAKE_UP = 3;

$guard_action_list = [];
$sorted_lines = [];

if ($file = fopen(__DIR__ . "/day4-input.txt", "rb")) {
    while (!feof($file)) {
        $line = trim(fgets($file));

        if (preg_match("@\[(.+)\]@", $line, $matches)) {
            // PHP seems to convert timestamps from before 1970 to a negative timestamp... ah well why not!
            $timestamp = DateTime::createFromFormat("Y-m-d H:i", $matches[1])->getTimestamp();
        }
        $sorted_lines[$timestamp] = $line;
    }
    fclose($file);
}

ksort($sorted_lines);
$guard_id = null;

// Now that it is sorted parse the data from the lines
foreach ($sorted_lines as $timestamp => $line) {
    if (preg_match("@#(\d+)@", $line, $matches)) {
        $guard_id = (int) $matches[1];
        $guard_action_list[$guard_id][] = ["timestamp" => $timestamp, "guard_id" => $guard_id, "action" => BEGIN_SHIFT];
    } else {
        // Everything else aka sleep/wake
        if (preg_match("@\] ([^#G\s]+)@", $line, $matches)) {
            if ($matches[1] === "wakes") {
                $guard_action_list[$guard_id][] = ["timestamp" => $timestamp, "guard_id" => $guard_id, "action" => WAKE_UP];
            } else {
                $guard_action_list[$guard_id][] = ["timestamp" => $timestamp, "guard_id" => $guard_id, "action" => FALL_ASLEEP];
            }
        }
    }
}

// Now count the guard with the most hours asleep
$sleep_count = [];
$highest_value_per_minute = [];
foreach ($guard_action_list as $guard_id => $guard_actions) {
    $sleep_count[$guard_id] = [];
    $sleep_start = 0;
    foreach ($guard_actions as $guard_action) {
        if ($guard_action["action"] === FALL_ASLEEP) {
            $sleep_start = $guard_action["timestamp"];
        }
        if ($guard_action["action"] === WAKE_UP) {
            $minutes_asleep = ($guard_action["timestamp"] - $sleep_start) / 60;
            $start_minute = date("i", $sleep_start);
            $end_minute = date("i", $guard_action["timestamp"]) - 1;
            for ($i = $start_minute; $i < $end_minute; $i++) {
                $sleep_count[$guard_id][$i]++;
            }
        }
    }
}


$sleep_count = [];
$sleep_minutes = [];
foreach ($guard_action_list as $guard_id => $guard_actions) {
    $sleep_count[$guard_id] = 0;
    $sleep_start = 0;
    foreach ($guard_actions as $guard_action) {
        if ($guard_action["action"] === FALL_ASLEEP) {
            $sleep_start = $guard_action["timestamp"];
        }
        if ($guard_action["action"] === WAKE_UP) {
            $sleep_count[$guard_id] = $sleep_count[$guard_id] + (($guard_action["timestamp"] - $sleep_start) / 60);

            $end_minute = date("i", $guard_action["timestamp"]);
            $start_minute = date("i", $sleep_start);
            for ($i = $start_minute; $i < $end_minute; $i++) {
                $sleep_minutes[$guard_id][$i]++;
            }
        }
    }
}

$top_sleeper_id = 0;
$top_sleeper_minute = 0;
$top_sleeper_amount = 0;
foreach ($sleep_minutes as $guard_id => $elem) {
    foreach ($elem as $minute => $amount) {
        if ($amount > $top_sleeper_amount) {
            $top_sleeper_amount = $amount;
            $top_sleeper_id = $guard_id;
            $top_sleeper_minute = $minute;
        }
    }
}

arsort($sleep_count);
$sleepy_guard = key($sleep_count);
$sleepy_guard_schedule = $sleep_minutes[$sleepy_guard];
arsort($sleepy_guard_schedule);
$minutes_asleep = reset($sleepy_guard_schedule);
array_flip($sleepy_guard_schedule);
$sleepy_minute = key($sleepy_guard_schedule);

echo "Guard #" . $sleepy_guard . " sleeps the most at xx:" . $sleepy_minute . "(" . $minutes_asleep . " minutes)" . " (answer = " . $sleepy_guard * $sleepy_minute . ")" . PHP_EOL;
echo "Top sleeper: #" . $top_sleeper_id . " at minute xx:" . $top_sleeper_minute . " for " . $top_sleeper_amount . " times in total" . " (answer = " . $top_sleeper_id * $top_sleeper_minute . ")" . PHP_EOL;
