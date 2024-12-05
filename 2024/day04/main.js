const { readFileSync } = require('fs');

const grid = readFileSync('input', 'utf-8').trim().split('\n');
const height = grid.length;
const width = grid[0].length;

function valid(x, y) {
  return 0 <= x && x < width && 0 <= y && y < height;
}

function count1(x, y) {
  if (grid[y][x] !== 'X') {
    return 0;
  }
  let c = 0;
  for (let dx = -1; dx <= 1; dx++) {
    for (let dy = -1; dy <= 1; dy++) {
      if (!valid(x + 3 * dx, y + 3 * dy)) {
        continue;
      }
      let word = '';
      for (let i = 0; i < 4; i++) {
        word += grid[y + i * dy][x + i * dx];
      }
      if (word === 'XMAS') {
        c++;
      }
    }
  }
  return c;
}

let sum = 0;
for (x = 0; x < width; x++) {
  for (y = 0; y < height; y++) {
    sum += count1(x, y);
  }
}
console.log(sum);

function count2(x, y) {
  if (grid[y][x] != 'A') {
    return 0;
  }
  const word = grid[y - 1][x - 1] + grid[y - 1][x + 1] + grid[y + 1][x + 1] + grid[y + 1][x - 1];
  if (['MMSS', 'MSSM', 'SMMS', 'SSMM'].includes(word)) {
    return 1;
  }
  return 0;
}

sum = 0;
for (x = 1; x < width - 1; x++) {
  for (y = 1; y < height - 1; y++) {
    sum += count2(x, y);
  }
}
console.log(sum);
