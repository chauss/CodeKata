import { readInput } from "./GeneralLogic.js";

// A =   Rock   = X (1 Point)
// B =   Paper  = Y (2 Points)
// C = Scissors = Z (3 Points)
// Win = 3 Points, Draw = 3 Points, Lose = 0 Points

const calculatePointsForHand = (hand: string): number => {
  if (hand === "X") {
    return 1;
  } else if (hand === "Y") {
    return 2;
  } else if (hand === "Z") {
    return 3;
  } else {
    throw Error('Hand is illegal: "' + hand + '"');
  }
};

const getMyPointsForGame = (opponentHand: string, myHand: string): number => {
  if (
    (myHand === "X" && opponentHand === "A") ||
    (myHand === "Y" && opponentHand === "B") ||
    (myHand === "Z" && opponentHand === "C")
  ) {
    return 3;
  } else if (
    (myHand === "X" && opponentHand === "C") ||
    (myHand === "Y" && opponentHand === "A") ||
    (myHand === "Z" && opponentHand === "B")
  ) {
    return 6;
  } else {
    return 0;
  }
};

const inputFile = await readInput("input_day2.txt");
const lines = inputFile.split("\n");

let myPointsForStrategie = 0;

for (const line of lines) {
  if (line.length != 3) {
    continue;
  }
  const opponent = line[0];
  const me = line[2];

  myPointsForStrategie +=
    calculatePointsForHand(me) + getMyPointsForGame(opponent, me);
}
console.log("DAY2");
console.log(
  "PART1: The strategie would lead me to " + myPointsForStrategie + " points"
);

// X = lose; Y = draw; Z = win

const whatToChose = (opponentHand: string, goal: string): string => {
  if (opponentHand === "A") {
    if (goal === "X") {
      return "Z";
    } else if (goal === "Y") {
      return "X";
    } else if (goal === "Z") {
      return "Y";
    }
  } else if (opponentHand === "B") {
    if (goal === "X") {
      return "X";
    } else if (goal === "Y") {
      return "Y";
    } else if (goal === "Z") {
      return "Z";
    }
  } else if (opponentHand === "C") {
    if (goal === "X") {
      return "Y";
    } else if (goal === "Y") {
      return "Z";
    } else if (goal === "Z") {
      return "X";
    }
  }
};

let myPointsForStrategiePart2 = 0;
for (const line of lines) {
  if (line.length != 3) {
    continue;
  }
  const opponent = line[0];
  const goal = line[2];
  const me = whatToChose(opponent, goal);

  myPointsForStrategiePart2 +=
    calculatePointsForHand(me) + getMyPointsForGame(opponent, me);
}

console.log(
  "PART2: The real strategie would lead me to " +
    myPointsForStrategiePart2 +
    " points"
);
