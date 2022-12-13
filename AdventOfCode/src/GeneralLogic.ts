import fs from "fs";

export const readInput = async (filename: string): Promise<string> => {
  return new Promise((resolve, reject) => {
    fs.readFile("./assets/" + filename, "utf8", (err, data) => {
      if (err) {
        reject(err);
      }
      resolve(data);
    });
  });
};
