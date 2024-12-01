<?php

declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day14 extends Day
{
    public function part1(): int|string
    {
        $input = $this->getInputAsArray(2020, 14, 1);
        $programs = [];

        $mask = null;
        foreach ($input as $line) {
            if (preg_match("@mask = (.*)@", $line, $matches)) {
                $mask = $matches[1];
            }

            if (preg_match("@mem\[(\d+)] = (\d+)@", $line, $matches)) {
                $mem = (int)$matches[1];
                $value = $matches[2];
                $programs[$mem] = $this->applyMask($mask, $value);
            }
        }

        return array_sum($programs);
    }

    public function part2(): int|string
    {
        $input = $this->getInputAsArray(2020, 14, 2);
        $programs = [];

        $mask = null;
        foreach ($input as $line) {
            if (preg_match("@mask = (.*)@", $line, $matches)) {
                $mask = $matches[1];
            }

            // Assuming mem[x] = 123 here
            if (preg_match("@mem\[(\d+)] = (\d+)@", $line, $matches)) {
                $mem = (int)$matches[1];
                $value = $matches[2];
                $tmp_adr = $this->getMemoryAdresses($mask, $mem);
                foreach ($tmp_adr as $a) {
                    $programs[$a] = $value;
                }
            }
        }

        return array_sum($programs);
    }

    // There might be a (PHP) function to do that XOR OR AND or some combi but I dont know it :)
    public function applyMask(string $mask, string $value): string
    {
        // Convert our number (string) to a binary string
        $binary_value = base_convert($value, 10, 2);
        $binary_value = str_pad($binary_value, 36, '0', STR_PAD_LEFT);
        $mask_array = str_split($mask);

        $this->log($binary_value);
        $this->log($mask);

        foreach ($mask_array as $id => $item) {
            if ($item === "X") {
                continue;
            }
            // Overwrite our value with the entry from the mask
            $binary_value[$id] = $item;
        }

        // Convert binary to number
        $num = @base_convert($binary_value, 2, 10);

        $this->log($binary_value . "    == " . $num);
        $this->log("");

        return $num;
    }

    public function getMemoryAdresses(string $mask, int $input_adr): array
    {
        $address = base_convert((string)$input_adr, 10, 2);
        $address = str_pad($address, 36, '0', STR_PAD_LEFT);
        $mask_array = str_split($mask);

        $this->log($address);
        $this->log($mask);

        $addresses[] = ''; // Give the code something to start with
        foreach ($mask_array as $index => $mask_item) {
            // When our item is 0 -> add the original item to each of our possible addresses
            if ($mask_item === "0") {
                foreach ($addresses as $i => $_) {
                    $addresses[$i][$index] = $address[$index];
                }
            }

            // When our item is 0 -> add "1" to each of our possible addresses
            if ($mask_item === "1") {
                foreach ($addresses as $i => $_) {
                    $addresses[$i][$index] = "1";
                }
            }

            // X is special, both add a 0 and a new entry with a 1
            if ($mask_item === "X") {
                foreach ($addresses as $i => $_) {
                    // Clone adres, add a 1 and add this address to the array
                    $address_clone = $addresses[$i];
                    $address_clone .= "1";
                    $addresses[] = $address_clone;

                    // Add 0 to original item from the array
                    $addresses[$i][$index] = "0";
                }
            }
        }

        if ($this->output?->isVerbose()) {
            $this->log("------------");
            foreach ($addresses as $address) {
                $this->log($address . "    == " . @base_convert($address, 2, 10));
            }
            $this->log("------------");
            $this->log("");
        }

        $addresses = array_map(static fn($elem) => base_convert($elem, 2, 10), $addresses);

        return $addresses;
    }
}
