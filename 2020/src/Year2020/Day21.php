<?php

declare(strict_types=1);

namespace mattie112\AdventOfCode\Year2020;

use mattie112\AdventOfCode\Day;

class Day21 extends Day
{
    public function part1(): int|string
    {
        $input = $this->getInputAsArray(2020, 21, 1);

        $recepies = [];
        $possible_allergens_for_ingredient = [];
        $recipies_with = [];

        foreach ($input as $recipe_id => $line) {
            [$ingredients, $allergens] = explode(" (contains ", $line);
            $allergens = str_replace(")", "", $allergens);
            $ingredients = explode(" ", $ingredients);
            $allergens = explode(", ", $allergens);
            $recepies[] = $ingredients;
            foreach ($allergens as $aller) {
                $recipies_with[$aller][] = $recipe_id;
            }
            foreach ($ingredients as $ingr) {
                if (!isset($possible_allergens_for_ingredient[$ingr])) {
                    $possible_allergens_for_ingredient[$ingr] = [];
                }
                $possible_allergens_for_ingredient[$ingr] = array_merge($possible_allergens_for_ingredient[$ingr], $allergens);
            }
        }

        $safe = [];

        foreach ($possible_allergens_for_ingredient as $ingredient => $possible_allergens) {
            $impossible = [];
            $c = $ingredient;
            foreach ($possible_allergens as $possible_allergen) {
                foreach ($recipies_with[$possible_allergen] as $i => $recipe_id) {
                    if (!in_array($ingredient, $recepies[$recipe_id], true)) {
                        // Found a recipe that contains the ingredient but not the allergen so it must be allergen-free
                        $impossible[] = $possible_allergen;
                        break;
                    }
                }
            }
            $a = array_diff($possible_allergens, $impossible);
            if (empty($a)) {
                $safe[] = $ingredient;
            }
        }

        $safe_count = 0;
        foreach ($safe as $ingredient) {
            foreach ($recepies as $ingr_recep) {
                if (in_array($ingredient, $ingr_recep, true)) {
                    $safe_count++;
                }
            }
        }

        return $safe_count;
    }

    public function part2(): int|string
    {
        $input = $this->getInputAsArray(2020, 21, 2);

        return 0;
    }

}
