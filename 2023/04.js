const { readFileSync } = require("node:fs");
const input = readFileSync(`${__dirname}/04_input.txt`, "utf8").split("\r\n");

// FIRST PART

function part1(input) {
    let totalPoints = 0;

    for (let line of input) {
        line = line.split(":");
        const list = line[1].split("|");
        const winningNumbers = list[0].split(" ");
        const numbers = list[1].split(" ");
        let points = 0;

        for (let winningNumber of winningNumbers) {
            if (winningNumber !== "" && numbers.includes(winningNumber)) {
                if (points === 0) {
                    points++;
                    continue;
                }
                points *= 2;
            }
        }
        totalPoints += points;
    }

    return totalPoints;
}

// SECOND PART

function part2(input) {
    let totalCards = 0;
    let multipliers = [];
    for (let index in input) {
        multipliers[index] = 1;
    }

    for (let line of input) {
        line = line.split(":");
        const id = parseInt(line[0].replace("Card ", ""));
        const list = line[1].split("|");
        const winningNumbers = list[0].split(" ");
        const numbers = list[1].split(" ");
        let wins = [];

        for (let winningNumber of winningNumbers) {
            if (winningNumber !== "" && numbers.includes(winningNumber)) {
                wins.push(winningNumber);
            }
        }

        let cardsMultiplier = multipliers[id - 1];

        for (let i = 1; i <= wins.length; i++) {
            multipliers[i + id - 1] += cardsMultiplier;
        }

        totalCards += cardsMultiplier;
    }

    return totalCards;
}

// OUTPUT

console.log(part1(input));
console.log(part2(input));
