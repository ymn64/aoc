grid = open('input').read().splitlines()

height = len(grid)
width = len(grid[0])


def valid(x, y):
	return 0 <= x and x < width and 0 <= y and y < height


def count1(x, y):
	if grid[y][x] != 'X':
		return 0
	c = 0
	for dx in [-1, 0, 1]:
		for dy in [-1, 0, 1]:
			if not valid(x + 3 * dx, y + 3 * dy):
				continue
			word = ''
			for i in range(4):
				word += grid[y + i * dy][x + i * dx]
			if word == 'XMAS':
				c += 1
	return c


# 2397
sum = 0
for x in range(width):
	for y in range(height):
		sum += count1(x, y)
print(sum)


def count2(x, y):
	if grid[y][x] != 'A':
		return 0
	word = grid[y - 1][x - 1] + grid[y - 1][x + 1] + grid[y + 1][x + 1] + grid[y + 1][x - 1]
	if word in ['MMSS', 'MSSM', 'SMMS', 'SSMM']:
		return 1
	return 0


# 1824
sum = 0
for x in range(1, width - 1):
	for y in range(1, height - 1):
		sum += count2(x, y)
print(sum)
