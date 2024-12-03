<?php

$input = file(__DIR__."/01_input.txt");

// FIRST PART

function part1($input) {
    $column1 = [];
    $column2 = [];
    $score = (int) 0;
    
    foreach ($input as $line) {
        $filteredDigits = array_filter(preg_split("/\D+/", $line));
        array_push($column1, $filteredDigits[0]);
        array_push($column2, $filteredDigits[1]);
    }

    sort($column1);
    sort($column2);

    foreach ($column1 as $key=>$value) {
        $score += (int) abs($value - $column2[$key]);
    }

    return $score;
}

// SECOND PART

function part2($input) {
    $column1 = [];
    $column2 = [];
    $score = (int) 0;
    
    foreach ($input as $line) {
        $filteredDigits = array_filter(preg_split("/\D+/", $line));
        array_push($column1, $filteredDigits[0]);
        array_push($column2, $filteredDigits[1]);
    }

    sort($column1);
    sort($column2);

    foreach ($column1 as $key=>$value) {
        $score += (int) $value * count(array_keys($column2, $value));
    }

    return $score;
}

// OUTPUT

print(part1($input)."\n");
print(part2($input)."\n");
