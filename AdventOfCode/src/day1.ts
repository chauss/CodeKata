import fs from "fs";

const readInput = async (): Promise<string> => {
  return new Promise((resolve, reject) => {
    fs.readFile("./assets/input_day1.txt", "utf8", (err, data) => {
      if (err) {
        reject(err);
      }
      resolve(data);
    });
  });
};

const file: string = await readInput();
const lines = file.split("\n");

const elvesCalories: number[] = [0];

for (const line of lines) {
  if (line.length == 0) {
    elvesCalories.push(0);
  } else {
    elvesCalories[elvesCalories.length - 1] += parseInt(line);
  }
}

console.log(
  "The elve with the most calories is carrying " +
    Math.max(...elvesCalories) +
    " calories"
);
