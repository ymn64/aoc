local left = {}
local right = {}

for line in io.lines('input') do
  local f = {}
  for number in string.gmatch(line, '%S+') do
    table.insert(f, tonumber(number))
  end
  table.insert(left, f[1])
  table.insert(right, f[2])
end

-- Part 1: 2164381
table.sort(left)
table.sort(right)
local sum = 0
for i = 1, #left do
  sum = sum + math.abs(left[i] - right[i])
end
print(sum)

local function freq(x, s)
  local f = 0
  for _, y in pairs(s) do
    if x == y then
      f = f + 1
    end
  end
  return f
end

-- Part 2: 20719933
sum = 0
for _, x in pairs(left) do
  sum = sum + x * freq(x, right)
end
print(sum)
