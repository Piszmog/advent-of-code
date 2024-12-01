import read/file
import strutils
import os

proc part1(path: string) =
  var total = 0
  file.handle(path) do (line: string):
    var firstNumber = 0
    var lastNumber = 0
    for char in line:
      if char.isDigit():
        let val = ord(char) - ord('0')
        if firstNumber == 0:
          firstNumber = val * 10
        lastNumber = val
    total += firstNumber + lastNumber

  echo "Total: ", total

type 
  Number = enum
    one = (1, "one"),
    two = (2, "two"),
    three = (3, "three"),
    four = (4, "four"),
    five = (5, "five"),
    six = (6, "six"),
    seven = (7, "seven"),
    eight = (8, "eight"),
    nine = (9, "nine")

proc findAll*(str: string, substr: string): seq[int] = 
  var indexes: seq[int] = @[]
  var start = 0
  while true:
    let index = str.find(substr, start)
    if index == -1:
      break
    indexes.add(index)
    start = index + substr.len()
  return indexes

proc part2(path: string) =
  var total = 0
  file.handle(path) do (line: string):
    var firstNumber = 0
    var firstNumberIndex = -1
    var lastNumber = 0
    var lastNumberIndex = -1

    var index = 0
    for char in line:
      if char.isDigit():
        let val = ord(char) - ord('0')
        if firstNumber == 0:
          firstNumber = val * 10
          firstNumberIndex = index
        lastNumber = val
        lastNumberIndex = index
      index += 1

    for num in Number:
      let indexes = line.findAll($num)
      if indexes.len() == 0:
        continue
      let lowIndex = indexes.low()
      let highIndex = indexes.high()
      if indexes[lowIndex] < firstNumberIndex or firstNumberIndex == -1:
        firstNumber = num.ord * 10
        firstNumberIndex = indexes[lowIndex]
      if indexes[highIndex] > lastNumberIndex:
        lastNumber = num.ord
        lastNumberIndex = indexes[highIndex]

    total += firstNumber + lastNumber

  echo "Total: ", total

when isMainModule:
  if commandLineParams().len() <= 0:
    echo "Missing argument to specify the input file"
    quit(1)

  let path = paramStr(1)
  part1(path)
  part2(path)

