local input = io.lines("input.txt")
local depths = {}
local ascending, ascendingWindow = 0, 0

-- init depths
for line in input do
  depths:insert(math.tointeger(line))
end

for i = 1, #depths - 1 do
  if depths[i + 1] > depths[i] then
    ascending = ascending + 1
  end

  if i <= #depths - 3 then
    local currWindow = depths[i] + depths[i + 1] + depths[i + 2]
    local nextWindow = depths[i + 1] + depths[i + 2] + depths[i + 3]
    if nextWindow > currWindow then
      ascendingWindow = ascendingWindow + 1
    end
  end
end

print(ascending)
print(ascendingWindow)
