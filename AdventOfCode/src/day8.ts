import { readInput } from "./GeneralLogic.js";

const parseForest = (lines: string[]): number[][] => {
  const forest: number[][] = [];

  for (const line of lines) {
    if (line.trim().length === 0) {
      continue;
    }
    const trees: string[] = line.split("");
    const treeHeights: number[] = trees.map((tree) => parseInt(tree));
    forest.push(treeHeights);
  }
  return forest;
};

const calcForestBorderTreeCount = (forest: number[][]): number => {
  return (forest.length + forest[0].length) * 2 - 4;
};

const inputFile = await readInput("input_day8.txt");
const lines = inputFile.split("\n");

const forest: number[][] = parseForest(lines);
let visibleTreesCount: number = calcForestBorderTreeCount(forest);

for (let row = 1; row <= forest.length - 2; row++) {
  for (let col = 1; col <= forest[0].length - 2; col++) {
    const currentTreeHeight: number = forest[row][col];
    let isTreeVisibleLeft: boolean = true;
    let isTreeVisibleRight: boolean = true;
    let isTreeVisibleTop: boolean = true;
    let isTreeVisibleBottom: boolean = true;
    for (let tmpCol = 0; tmpCol <= forest[0].length - 1; tmpCol++) {
      if (forest[row][tmpCol] >= currentTreeHeight) {
        if (tmpCol < col) {
          isTreeVisibleLeft = false;
        } else if (tmpCol > col) {
          isTreeVisibleRight = false;
        }
      }
    }
    for (let tmpRow = 0; tmpRow <= forest.length - 1; tmpRow++) {
      if (forest[tmpRow][col] >= currentTreeHeight) {
        if (tmpRow < row) {
          isTreeVisibleTop = false;
        } else if (tmpRow > row) {
          isTreeVisibleBottom = false;
        }
      }
    }
    if (
      isTreeVisibleLeft ||
      isTreeVisibleRight ||
      isTreeVisibleTop ||
      isTreeVisibleBottom
    ) {
      visibleTreesCount++;
    }
  }
}

console.log("DAY8");
console.log(
  "PART1: From outside there are " + visibleTreesCount + " trees visible."
);

const treeScenicScore: number[][] = [];

for (let row = 1; row <= forest.length - 2; row++) {
  treeScenicScore.push([]);
  for (let col = 1; col <= forest[0].length - 2; col++) {
    const currentTreeHeight: number = forest[row][col];
    let sightLeft: number = 0;
    let sightRight: number = 0;
    let sightTop: number = 0;
    let sightBottom: number = 0;

    // left
    for (let tmpCol = col - 1; tmpCol >= 0; tmpCol--) {
      sightLeft++;
      if (forest[row][tmpCol] >= currentTreeHeight) {
        break;
      }
    }
    // right
    for (let tmpCol = col + 1; tmpCol <= forest[0].length - 1; tmpCol++) {
      sightRight++;
      if (forest[row][tmpCol] >= currentTreeHeight) {
        break;
      }
    }
    // top
    for (let tmpRow = row - 1; tmpRow >= 0; tmpRow--) {
      sightTop++;
      if (forest[tmpRow][col] >= currentTreeHeight) {
        break;
      }
    }
    // bottom
    for (let tmpRow = row + 1; tmpRow <= forest.length - 1; tmpRow++) {
      sightBottom++;
      if (forest[tmpRow][col] >= currentTreeHeight) {
        break;
      }
    }
    treeScenicScore[row - 1].push(
      sightLeft * sightRight * sightTop * sightBottom
    );
  }
}

let biggestScenicValue = 0;

for (let row = 1; row <= forest.length - 2; row++) {
  treeScenicScore.push([]);
  for (let col = 1; col <= forest[0].length - 2; col++) {
    if (treeScenicScore[row][col] > biggestScenicValue) {
      biggestScenicValue = treeScenicScore[row][col];
    }
  }
}

console.log("PART2: The biggest scenic score is " + biggestScenicValue);
