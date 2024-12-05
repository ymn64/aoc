local grid = {}

for line in io.lines('input') do
  table.insert(grid, line)
end

local height = #grid
local width = #grid[1]

local function at(x, y)
  return grid[y]:sub(x, x)
end

local function valid(x, y)
  return 0 < x and x <= width and 0 < y and y <= height
end

local function count1(x, y)
  if at(x, y) ~= 'X' then
    return 0
  end
  local c = 0
  for dx = -1, 1 do
    for dy = -1, 1 do
      if valid(x + 3 * dx, y + 3 * dy) then
        local word = ''
        for i = 0, 3 do
          word = word .. at(x + i * dx, y + i * dy)
        end
        if word == 'XMAS' then
          c = c + 1
        end
      end
    end
  end
  return c
end

-- 2397
local sum = 0
for x = 1, width do
  for y = 1, height do
    sum = sum + count1(x, y)
  end
end
print(sum)

local function count2(x, y)
  if at(x, y) ~= 'A' then
    return 0
  end
  local word = at(x - 1, y - 1) .. at(x + 1, y - 1) .. at(x + 1, y + 1) .. at(x - 1, y + 1)
  if word == 'MMSS' or word == 'MSSM' or word == 'SMMS' or word == 'SSMM' then
    return 1
  end
  return 0
end

-- 1824
sum = 0
for x = 2, width - 1 do
  for y = 2, height - 1 do
    sum = sum + count2(x, y)
  end
end
print(sum)

