const { readFileSync } = require("node:fs");
const input = readFileSync(`${__dirname}/02_input.txt`, "utf8").split("\n");

// FIRST PART

function part1(input) {
    let ids = 0;

    line: for (let line of input) {
        line = line.split(":");
        const id = parseInt(line[0].replace("Game ", ""));
        const sets = line[1].split(";");

        for (let set of sets) {
            let cubes = { red: 0, green: 0, blue: 0 };
            set = set.split(",");
            for (let cube of set) {
                if (cube.includes("red")) {
                    cubes.red = cubes.red + parseInt(cube);
                } else if (cube.includes("green")) {
                    cubes.green = cubes.green + parseInt(cube);
                } else if (cube.includes("blue")) {
                    cubes.blue = cubes.blue + parseInt(cube);
                }
            }
            if (cubes.green > 13 || cubes.red > 12 || cubes.blue > 14)
                continue line;
        }
        ids += id;
    }

    return ids;
}

// SECOND PART

function part2(input) {
    let power = 0;

    for (let line of input) {
        line = line.split(":");
        let cubes = { red: 0, green: 0, blue: 0 };
        const sets = line[1].split(";");

        for (let set of sets) {
            set = set.split(",");
            for (let cube of set) {
                if (cube.includes("red") && parseInt(cube) > cubes.red) {
                    cubes.red = parseInt(cube);
                } else if (
                    cube.includes("green") &&
                    parseInt(cube) > cubes.green
                ) {
                    cubes.green = parseInt(cube);
                } else if (
                    cube.includes("blue") &&
                    parseInt(cube) > cubes.blue
                ) {
                    cubes.blue = parseInt(cube);
                }
            }
        }
        power += cubes.red * cubes.green * cubes.blue;
    }

    return power;
}

// OUTPUT

console.log(part1(input));
console.log(part2(input));
