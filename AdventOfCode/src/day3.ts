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
const lines = inputFile.split("\n").filter((line) => line.length > 0);
const duplicateItems = findDuplicateItems(lines);

let prioSum = 0;

for (const item of duplicateItems) {
  prioSum += calculateItemPriority(item);
}

console.log("DAY3");
console.log("PART1: The sum of the duplicate items priotiry is: " + prioSum);

const elvesTeamBadges: string[] = [];

for (let i = 0; i < lines.length; i += 3) {
  const elveOneItems = lines[i];
  const elveTwoItems = lines[i + 1];
  const elveThreeItems = lines[i + 2];

  const commonItems = new Set<string>();

  for (const item of elveOneItems) {
    commonItems.add(item);
  }
  for (const item of commonItems) {
    if (elveTwoItems.indexOf(item) < 0 || elveThreeItems.indexOf(item) < 0) {
      commonItems.delete(item);
    }
  }
  const [first] = commonItems;
  elvesTeamBadges.push(first);
}

let teamPrioSum = 0;
for (const item of elvesTeamBadges) {
  console.log(item);
  teamPrioSum += calculateItemPriority(item);
}

console.log("PART2: The sum of the team batches priority is: " + teamPrioSum);
