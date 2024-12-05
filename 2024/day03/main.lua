local function prod(mem)
  local sum = 0
  for x, y in string.gmatch(mem, 'mul%((%d+),(%d+)%)') do
    sum = sum + tonumber(x) * tonumber(y)
  end
  return sum
end

local file = io.open('input', 'r')
local mem = file:read('*a')
file:close()

-- Part 1: 173517243
print(prod(mem))

-- Part 2: 100450138
local sum = 0
while true do
  local start = string.find(mem, "don't%(%)")
  if not start then
    break
  end
  sum = sum + prod(mem:sub(1, start - 1))
  local end_pos = string.find(mem, 'do%(%)', start)
  if not end_pos then
    break
  end
  mem = mem:sub(end_pos + 3)
end
print(sum)
