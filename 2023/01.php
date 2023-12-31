<?php

$input = file(__DIR__."/01_input.txt");

// FIRST PART

function part1($input) {
    $calibrations = (int) 0;
    
    foreach ($input as $line) {
        $filteredDigits = array_filter(preg_split("/\D+/", $line));

        $firstDigit = (int) substr(reset($filteredDigits), 0, 1);
        $lastDigit = (int) substr(end($filteredDigits), -1);

        $combination = (int) ($firstDigit . $lastDigit);
        $calibrations += $combination;
    }

    return $calibrations;
}

// SECOND PART

function part2($input) { 
    $calibrations = (int) 0;
    
    foreach ($input as $line) {
        $values = str_replace(["oneight", "twone", "eightwo"], [18, 21, 82], $line);
        $values = str_replace(["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"], [1, 2, 3, 4, 5, 6, 7, 8, 9], $values);

        $filteredDigits = array_filter(preg_split("/\D+/", $values));

        $firstDigit = (int) substr(reset($filteredDigits), 0, 1);
        $lastDigit = (int) substr(end($filteredDigits), -1);

        $combination = (int) ($firstDigit . $lastDigit);
        $calibrations += $combination;
    }

    return $calibrations;
}

// OUTPUT

print(part1($input)."\n");
print(part2($input));
