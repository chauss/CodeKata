import { readInput } from "./GeneralLogic.js";

const currentLocationInFileSystem = (
  fileSystem: object,
  pwd: string[]
): object => {
  let currentLocation = fileSystem;
  for (const pathDir of pwd) {
    currentLocation = currentLocation[pathDir];
  }
  return currentLocation;
};

const buildFileSystem = (lines: string[]): object => {
  let pwd: string[] = [];
  const fileSystem = { files: [] };

  for (const line of lines) {
    if (line.trim().length === 0) {
      continue;
    }
    const components = line.split(" ");
    if (components[0] === "$") {
      // command
      if (components[1] === "cd") {
        const dir = components[2];
        if (dir === "/") {
          pwd = [];
        } else if (dir === "..") {
          pwd.pop();
        } else {
          pwd.push(dir);
        }
      } else if (components[1] === "ls") {
        // we don't care about that
      }
    } else if (components[0] === "dir") {
      // dir
      const dirName = components[1];
      let currentLocation = currentLocationInFileSystem(fileSystem, pwd);
      if (!currentLocation[dirName]) {
        currentLocation[dirName] = { files: [] };
      }
    } else {
      // file
      const fileSize = parseInt(components[0]);
      const fileName = components[1];
      let currentLocation = currentLocationInFileSystem(fileSystem, pwd);
      currentLocation["files"].push({ name: fileName, size: fileSize });
    }
  }
  return fileSystem;
};

const calcDirSize = (dirInFileSystem): number => {
  let fileSizeInDir: number = 0;
  if (dirInFileSystem["files"]) {
    for (const file of dirInFileSystem["files"]) {
      fileSizeInDir += file["size"];
    }
  }
  let subDirsSize: number = 0;
  for (const key in dirInFileSystem) {
    if (!dirInFileSystem.hasOwnProperty(key) || key === "files") {
      continue;
    }
    subDirsSize += calcDirSize(dirInFileSystem[key]);
  }
  return subDirsSize + fileSizeInDir;
};

const sumOfDirSizes = (rootDir, maxDirSizeToConsider: number): number => {
  let dirSizeSum: number = 0;

  const currentDirSize = calcDirSize(rootDir);
  if (currentDirSize <= maxDirSizeToConsider) {
    dirSizeSum += currentDirSize;
  }

  for (const key in rootDir) {
    if (!rootDir.hasOwnProperty(key) || key === "files") {
      continue;
    }

    dirSizeSum += sumOfDirSizes(rootDir[key], maxDirSizeToConsider);
  }
  return dirSizeSum;
};

const sizeOfSmallestDirGreaterThan = (
  rootDir,
  minSize: number,
  currentWinningSize: number = 999999999999999
): number => {
  let possibleNewWinningSize: number = currentWinningSize;
  const currentSize = calcDirSize(rootDir);
  if (currentSize >= minSize && currentSize < currentWinningSize) {
    possibleNewWinningSize = currentSize;
  }

  // Check if subDirs are smaller only if the current dir is bigger than the minSize
  if (currentSize > minSize) {
    for (const key in rootDir) {
      if (!rootDir.hasOwnProperty(key) || key === "files") {
        continue;
      }

      possibleNewWinningSize = sizeOfSmallestDirGreaterThan(
        rootDir[key],
        minSize,
        possibleNewWinningSize
      );
    }
  }
  return possibleNewWinningSize;
};

const inputFile = await readInput("input_day7.txt");
const lines = inputFile.split("\n");

const fileSystem = buildFileSystem(lines);

console.log("DAY7");
console.log(
  "PART1: The sum of all dirs that are smaller than 100000 is: " +
    sumOfDirSizes(fileSystem, 100000)
);

const totalDiskSpace: number = 70000000;
const usedDiskSpace: number = calcDirSize(fileSystem);
const requiredDiskSpace: number = 30000000;
const leftDiskSpace: number = totalDiskSpace - usedDiskSpace;
const missingDiskSpace: number = requiredDiskSpace - leftDiskSpace;

console.log(
  "PART2: The size of the smallest dir that has a size of at least " +
    missingDiskSpace +
    " is: " +
    sizeOfSmallestDirGreaterThan(fileSystem, missingDiskSpace)
);
