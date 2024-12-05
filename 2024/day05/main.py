from functools import cmp_to_key

chunks = open('input').read().split('\n\n')
chunks = [chunk.splitlines() for chunk in chunks]

rules = {}
for line in chunks[0]:
	l, r = line.split('|')
	if l not in rules:
		rules[l] = []
	rules[l].append(r)

updates = [line.split(',') for line in chunks[1]]


def cmp(a, b):
	return -1 if b in rules[a] else 1


def check(update):
	for i in range(1, len(update)):
		if cmp(update[i - 1], update[i]) > 0:
			return False
	return True


# Part 1: 4609
sum = 0
for u in updates:
	if check(u):
		mid = len(u) // 2
		sum += int(u[mid])
print(sum)

# Part 2: 5723
sum = 0
for u in updates:
	if not check(u):
		u.sort(key=cmp_to_key(cmp))
		mid = len(u) // 2
		sum += int(u[mid])
print(sum)
