local reports = {}

for line in io.lines('input') do
  local levels = {}
  for level in line:gmatch('%S+') do
    table.insert(levels, tonumber(level))
  end
  table.insert(reports, levels)
end

local function safe(report)
  local d0 = report[2] - report[1]
  for i = 2, #report do
    local d = report[i] - report[i - 1]
    if d == 0 or math.abs(d) > 3 or d * d0 < 0 then
      return false
    end
  end
  return true
end

-- Part 1: 591
local sum = 0
for _, r in ipairs(reports) do
  if safe(r) then
    sum = sum + 1
  end
end
print(sum)

-- Part 2: 621
local sum2 = 0
for _, r in ipairs(reports) do
  if safe(r) then
    sum2 = sum2 + 1
  else
    for i = 1, #r do
      local modified = {}
      for j, v in ipairs(r) do
        modified[j] = v
      end
      table.remove(modified, i)
      if safe(modified) then
        sum2 = sum2 + 1
        break
      end
    end
  end
end
print(sum2)
