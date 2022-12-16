import { readInput } from "./GeneralLogic.js";

const inputFile = await readInput("input_day4.txt");
const lines = inputFile.split("\n").filter((line) => line.length > 0);

let fullyContainingAssignmentCount = 0;
let partlyOverlappingAssignmentCount = 0;

for (const line of lines) {
  const assignments = line.trim().split(",");
  const elveOneSections = assignments[0];
  const elveTwoSections = assignments[1];

  const elveOneSectionsSplit = elveOneSections.split("-");
  const elveTwoSectionsSplit = elveTwoSections.split("-");

  const elveOneStartSection = parseInt(elveOneSectionsSplit[0]);
  const elveOneEndSection = parseInt(elveOneSectionsSplit[1]);

  const elveTwoStartSection = parseInt(elveTwoSectionsSplit[0]);
  const elveTwoEndSection = parseInt(elveTwoSectionsSplit[1]);

  if (
    (elveOneStartSection >= elveTwoStartSection &&
      elveOneEndSection <= elveTwoEndSection) ||
    (elveTwoStartSection >= elveOneStartSection &&
      elveTwoEndSection <= elveOneEndSection)
  ) {
    fullyContainingAssignmentCount++;
  }
  if (
    (elveTwoStartSection >= elveOneStartSection &&
      elveTwoStartSection <= elveOneEndSection) ||
    (elveTwoEndSection >= elveOneStartSection &&
      elveTwoEndSection <= elveOneEndSection) ||
    (elveOneStartSection >= elveTwoStartSection &&
      elveOneStartSection <= elveTwoEndSection) ||
    (elveOneEndSection >= elveTwoStartSection &&
      elveOneEndSection <= elveTwoEndSection)
  ) {
    partlyOverlappingAssignmentCount++;
  }
}

console.log("DAY4");
console.log(
  "PART1: There are " +
    fullyContainingAssignmentCount +
    " assignment pairs in which one fully contains the other"
);
console.log(
  "PART2: There are " +
    partlyOverlappingAssignmentCount +
    " assignment pairs in which one fully contains the other"
);
