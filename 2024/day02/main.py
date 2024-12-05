reports = []
lines = open('input').readlines()
for line in lines:
	levels = list(map(int, line.split()))
	reports.append(levels)


def safe(report):
	d0 = report[1] - report[0]
	for i in range(1, len(report)):
		d = report[i] - report[i - 1]
		if d == 0 or abs(d) > 3 or d * d0 < 0:
			return False
	return True


# Part 1: 591
print(sum(1 for r in reports if safe(r)))


# Part 2: 621
sum = 0
for r in reports:
	if safe(r):
		sum += 1
		continue
	for i in range(len(r)):
		if safe(r[:i] + r[i + 1 :]):
			sum += 1
			break
print(sum)
