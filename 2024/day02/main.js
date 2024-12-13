const { readFileSync } = require('fs');

const reports = readFileSync('input', 'utf-8')
  .trim()
  .split('\n')
  .map((line) => line.split(' ').map(Number));

function safe(report) {
  d0 = report[1] - report[0];
  for (let i = 1; i < report.length; i++) {
    d = report[i] - report[i - 1];
    if (d == 0 || Math.abs(d) > 3 || d * d0 < 0) {
      return false;
    }
  }
  return true;
}

// Part1: 591
let sum = 0;
for (r of reports) {
  if (safe(r)) {
    sum++;
  }
}
console.log(sum);

// Part2: 621
sum = 0;
for (r of reports) {
  if (safe(r)) {
    sum++;
    continue;
  }
  for (let i = 0; i < r.length; i++) {
    if (safe(r.slice(0, i).concat(r.slice(i + 1)))) {
      sum++;
      break;
    }
  }
}
console.log(sum);
