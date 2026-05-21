local input = io.lines("input.txt")
local x, y = 0, 0
local cY = 0

for line in input do
  local dir, amt = line:match("(%a+)%s(%d+)")
  if dir == "forward" then
    x = x + amt
    cY = cY + y * amt
  elseif dir == "up" then
    y = y - amt
  elseif dir == "down" then
    y = y + amt
  end
end

print(x * y)
print(cY * x)
