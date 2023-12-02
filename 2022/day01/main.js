const fs = require("fs");

const calories = fs
  .readFileSync("input", "utf8")
  .trim()
  .split("\n\n")
  .map((chunk) =>
    chunk
      .split("\n")
      .map((x) => Number.parseInt(x))
      .reduce((x, y) => x + y)
  )
  .sort()
  .reverse();

console.log(calories[0]);
console.log(calories[0] + calories[1] + calories[2]);
