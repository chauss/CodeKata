import { readInput } from "./GeneralLogic.js";

const calculateItemPriority = (item: string): number => {
  const lowerCaseStartPrio = 1;
  const upperCaseStartPrio = 27;

  const lowerCaseItemPrioOffset = "a".charCodeAt(0);
  const upperCaseItemPrioOffset = "A".charCodeAt(0);

  let itemPrio;
  if (item === item.toLocaleLowerCase()) {
    itemPrio =
      item.charCodeAt(0) - lowerCaseItemPrioOffset + lowerCaseStartPrio;
  } else {
    itemPrio =
      item.charCodeAt(0) - upperCaseItemPrioOffset + upperCaseStartPrio;
  }
  return itemPrio;
};

const findDuplicateItems = (lines: string[]): string[] => {
  const duplicateItems: string[] = [];

  for (const line of lines) {
    const compartmentOne = line.substring(0, line.length / 2);
    const compartmentTwo = line.substring(line.length / 2, line.length);
    const backpackOnlyDuplicateItems = new Set<string>();

    for (const item of compartmentOne) {
      if (compartmentTwo.indexOf(item) >= 0) {
        backpackOnlyDuplicateItems.add(item);
      }
    }
    for (const item of backpackOnlyDuplicateItems) {
      duplicateItems.push(item);
    }
  }
  return duplicateItems;
};

const inputFile = await readInput("input_day3.txt");
const lines = inputFile.split("\n");
const duplicateItems = findDuplicateItems(lines);

let prioSum = 0;

for (const item of duplicateItems) {
  prioSum += calculateItemPriority(item);
}

console.log("DAY3");
console.log("PART1: The sum of the duplicate items priotiry is: " + prioSum);

const elvesTeamBadges: string[] = [];
let counter = 0;

for (let i = 0; i < lines.length; i += 3) {
  const elveOne = lines[i];
  const elveTwo = lines[i + 1];
  const elveThree = lines[i + 2];
}
