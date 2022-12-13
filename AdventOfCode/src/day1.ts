import { readInput } from "./GeneralLogic.js";

const inputFile = await readInput("input_day1.txt");
const lines = inputFile.split("\n");

let elvesCalories: number[] = [0];

for (const line of lines) {
  if (line.length == 0) {
    elvesCalories.push(0);
  } else {
    elvesCalories[elvesCalories.length - 1] += parseInt(line);
  }
}

console.log("DAY1");
console.log(
  "PART1: The elve with the most calories is carrying " +
    Math.max(...elvesCalories) +
    " calories"
);

elvesCalories = elvesCalories.sort((a, b) => b - a); // inverse sort

const caloriesOfThreeMostCarrying =
  elvesCalories[0] + elvesCalories[1] + elvesCalories[2];

console.log(
  "PART2: The three most carrying elves carry a total of " +
    caloriesOfThreeMostCarrying +
    " calories"
);
