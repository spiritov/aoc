local input = io.lines("input.txt")
local binary, transposedBinary = {}, {}
local mcb, lcb -- most / least common bits

for line in input do
  table.insert(binary, line)
end

local function findMcb(bt)
  -- init transposed binary with empty values, to concat later
  local tbt = {}
  for _ = 1, #bt[1] do
    table.insert(tbt, "")
  end
  -- transpose binary
  for i = 1, #bt do
    for j = 1, #bt[i] do
      local c = string.sub(bt[i], j, j)
      tbt[j] = tbt[j] .. c
    end
  end

  -- count zeroes and ones
  local mcb = ""
  local lcb = ""
  for _, b in pairs(tbt) do
    local _, zeroCount = string.gsub(b, "0", "")
    local _, oneCount = string.gsub(b, "1", "")
    if oneCount >= zeroCount then
      mcb = mcb .. "1"
      lcb = lcb .. "0"
    else
      mcb = mcb .. "0"
      lcb = lcb .. "1"
    end
  end
  return mcb, lcb
end

mcb, lcb = findMcb(binary) -- gamma rate, epsilon rate
print(tonumber(mcb, 2) * tonumber(lcb, 2))

local function findMatch(matches, bitIndex, useMcb)
  -- return if finished
  if #matches == 1 then
    return matches[1]
  end

  local mcb, lcb = findMcb(matches)
  -- set specific bit to match against
  local currBit
  if useMcb then
    currBit = mcb:sub(bitIndex, bitIndex)
  else
    currBit = lcb:sub(bitIndex, bitIndex)
  end

  local currMatches = {}
  for _, b in pairs(matches) do
    -- insert match
    if (b:sub(bitIndex, bitIndex) == currBit) then
      table.insert(currMatches, b)
    end
  end

  -- else continue
  return findMatch(currMatches, bitIndex + 1, useMcb)
end

local oxygenRating = (findMatch(binary, 1, true))
local coRating = (findMatch(binary, 1, false))
print(tonumber(oxygenRating, 2) * tonumber(coRating, 2))
