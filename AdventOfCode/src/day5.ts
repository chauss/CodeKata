import { readInput } from "./GeneralLogic.js";

const movementRegex = new RegExp(
  "move (?<amount>\\d*) from (?<from>\\d*) to (?<to>\\d*)"
);

const readStorageInfo = (lines: string[]) => {
  const storageLines: string[] = [];
  for (const line of lines) {
    if (line.trim().length === 0) {
      break;
    }
    storageLines.push(line);
  }
  const storageNumbersLine = storageLines[storageLines.length - 1].trim();
  const storagePlaces: number = parseInt(
    storageNumbersLine[storageNumbersLine.length - 1]
  );

  storageLines.pop();
  return { storageLines, storagePlaces };
};

const buildCurrentStorage = (
  storageLines: string[],
  storagePlaces: number
): string[][] => {
  const storage: string[][] = [];

  for (let i = 0; i < storagePlaces; i++) {
    storage.push([]);
  }

  for (let i = storageLines.length - 1; i >= 0; i--) {
    for (let j = 0; j < storagePlaces; j++) {
      const currentLocation = j * 4;

      const item = storageLines[i][currentLocation + 1];
      if (item.trim().length === 1) {
        storage[j].push(item);
      }
    }
  }
  return storage;
};

const readMovements = (lines: string[]): string[] => {
  let isAfterStorage = false;
  const movementLines: string[] = [];

  for (const line of lines) {
    if (line.trim().length === 0) {
      isAfterStorage = true;
    }
    if (isAfterStorage && line.trim().length !== 0) {
      movementLines.push(line);
    }
  }

  return movementLines;
};

const parseMovement = (movementLine: string) => {
  const movement = movementRegex.exec(movementLine).groups;
  const amount = parseInt(movement.amount);
  const from = parseInt(movement.from) - 1;
  const to = parseInt(movement.to) - 1;

  return { amount, from, to };
};

const getTopItemsOfStorage = (storage: string[][]): string => {
  let result = "";
  for (let i = 0; i < storagePlaces; i++) {
    result += storage[i][storage[i].length - 1];
  }
  return result;
};

const inputFile = await readInput("input_day5.txt");
const lines = inputFile.split("\n");

const { storageLines, storagePlaces } = readStorageInfo(lines);
let storage = buildCurrentStorage(storageLines, storagePlaces);
const movementLines = readMovements(lines);

for (const line of movementLines) {
  const { amount, from, to } = parseMovement(line);

  for (let i = 0; i < amount; i++) {
    const item = storage[from].pop();
    storage[to].push(item);
  }
}

console.log("DAY5");
console.log(
  "PART1: After all the movements the top items of the storage are: " +
    getTopItemsOfStorage(storage)
);

storage = buildCurrentStorage(storageLines, storagePlaces);
for (const line of movementLines) {
  const { amount, from, to } = parseMovement(line);

  const movedItems: string[] = [];
  for (let i = 0; i < amount; i++) {
    movedItems.push(storage[from].pop());
  }
  while (movedItems.length > 0) {
    storage[to].push(movedItems.pop());
  }
}

console.log(
  "PART2: After all the movements the top items of the storage are: " +
    getTopItemsOfStorage(storage)
);
