from re import compile

with open('input', 'r') as file:
	mem = file.read()


def prod(mem):
	sum = 0
	matches = compile(r'mul\(\d+,\d+\)').findall(mem)
	for match in matches:
		x, y = map(int, match[4:-1].split(','))
		sum += x * y
	return sum


# Part 1: 173517243
print(prod(mem))

# Part 2: 100450138
sum = 0
while True:
	start = mem.find("don't()")
	if start == -1:
		break
	sum += prod(mem[:start])
	end = mem.find('do()', start)
	if end == -1:
		break
	mem = mem[end + 4 :]
print(sum)
