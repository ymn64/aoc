const { readFileSync } = require('fs');

const left = [];
const right = [];

readFileSync('input', 'utf-8')
  .trim()
  .split('\n')
  .forEach((line) => {
    const f = line.split('   ').map(Number);
    left.push(f[0]);
    right.push(f[1]);
  });

// Part 1: 2164381
left.sort();
right.sort();
let sum = 0;
for (let i = 0; i < left.length; i++) {
  sum += Math.abs(left[i] - right[i]);
}
console.log(sum);

// Part 2: 20719933
sum = 0;
for (const x of left) {
  sum += x * right.filter((y) => y === x).length;
}
console.log(sum);
