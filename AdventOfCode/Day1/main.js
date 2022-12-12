const fs = require("fs");

const readInput = async () => {
	return new Promise(
		(resolve, reject) => {
			fs.readFile('./input.txt', 'utf8', (err, data) => {
				if (err) {
				  reject(err);
				}
				resolve(data);
			  });
		}
	)
};

console.log(await readInput());