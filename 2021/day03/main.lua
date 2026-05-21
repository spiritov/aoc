local input = io.lines("input.txt")
local binary, mappedBinary = {}, {}
local gammaRates, epsilonRates = "", ""
-- local oxygenRatings, coRatings = {}, {}

for line in input do
  table.insert(binary, line)
end

-- init empty mappedBinary
for _ = 1, #binary[1] do
  table.insert(mappedBinary, "")
end

-- "rotate" binary
for i = 1, #binary do
  for j = 1, #binary[i] do
    local c = string.sub(binary[i], j, j)
    mappedBinary[j] = mappedBinary[j] .. c
  end
end

-- count zeroes and ones
for _, b in pairs(mappedBinary) do
  local _, zeroCount = string.gsub(b, "0", "")
  local _, oneCount = string.gsub(b, "1", "")

  if oneCount >= zeroCount then
    gammaRates = gammaRates .. "1"
    epsilonRates = epsilonRates .. "0"
  else
    gammaRates = gammaRates .. "0"
    epsilonRates = epsilonRates .. "1"
  end
end

-- part 1
print(tonumber(gammaRates, 2) * tonumber(epsilonRates, 2))

-- for part 2..
-- 1. loop through binary with gammaRates and epsilonRates
-- 2. find oxygenRatings and coRatings using this, til 1 match is left for each
