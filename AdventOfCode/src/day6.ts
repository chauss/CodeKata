import { readInput } from "./GeneralLogic.js";

const stringHasDuplicates = (str: string): boolean => {
  return /(.).*\1/.test(str);
};

const getMessageStartIndex = (
  message: string,
  uniqueRequiredLength: number
): number => {
  for (let i = uniqueRequiredLength; i < message.length; i++) {
    const possibleStart = message.substring(i - uniqueRequiredLength, i);

    if (!stringHasDuplicates(possibleStart)) {
      return i;
    }
  }
  return NaN;
};

const inputFile = await readInput("input_day6.txt");

console.log("DAY6");
console.log(
  "PART1: There need to be " +
    getMessageStartIndex(inputFile, 4) +
    " characters processed until the start of the first packet."
);
console.log(
  "PART2: There need to be " +
    getMessageStartIndex(inputFile, 14) +
    " characters processed until the start of the first message."
);
