const { readFileSync } = require('fs');

const chunks = readFileSync('input', 'utf-8')
  .trim()
  .split('\n\n')
  .map((chunk) => chunk.split('\n'));

const rules = {};
chunks[0].forEach((line) => {
  const [l, r] = line.split('|');
  if (!rules[l]) {
    rules[l] = [];
  }
  rules[l].push(r);
});

const updates = chunks[1].map((line) => line.split(','));

function cmp(a, b) {
  return rules[a].includes(b) ? -1 : 1;
}

function check(u) {
  for (let i = 1; i < u.length; i++) {
    if (cmp(u[i - 1], u[i]) > 0) {
      return false;
    }
  }
  return true;
}

// Part 1: 4609
let sum = 0;
updates.forEach((u) => {
  if (check(u)) {
    const mid = Math.floor(u.length / 2);
    sum += parseInt(u[mid], 10);
  }
});
console.log(sum);

// Part 1: 5723
sum = 0;
updates.forEach((u) => {
  if (!check(u)) {
    u.sort(cmp);
    const mid = Math.floor(u.length / 2);
    sum += parseInt(u[mid], 10);
  }
});
console.log(sum);
