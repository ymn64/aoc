left = []
right = []

for line in open('input').read().splitlines():
	f = line.split('   ')
	left.append(int(f[0]))
	right.append(int(f[1]))

# Part 1: 2164381
left.sort()
right.sort()
print(sum(abs(left[i] - right[i]) for i in range(len(left))))


# Part 2: 20719933
print(sum(x * right.count(x) for x in left))
