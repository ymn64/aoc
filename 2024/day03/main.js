const { readFileSync } = require('fs');

let mem = readFileSync('input', 'utf-8');

function prod(mem) {
  let sum = 0;
  const matches = mem.match(/mul\(\d+,\d+\)/g);
  for (const match of matches) {
    const [x, y] = match.slice(4, -1).split(',').map(Number);
    sum += x * y;
  }
  return sum;
}

// Part 1: 173517243
console.log(prod(mem));

// Part 2: 100450138
let sum = 0;
while (true) {
  const start = mem.indexOf("don't()");
  if (start === -1) {
    break;
  }
  sum += prod(mem.slice(0, start));
  const end = mem.indexOf('do()', start);
  if (end === -1) {
    break;
  }
  mem = mem.slice(end + 4);
}
console.log(sum);
