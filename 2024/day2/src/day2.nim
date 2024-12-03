import strutils
import sequtils
import os
import options

proc getPath(): string =
  if commandLineParams().len() <= 0:
    echo "Missing argument to specify the input file"
    quit(1)
  return paramStr(1)

type HandleCallback = proc(line: string) {.closure.}

proc handle(path: string, callback: HandleCallback) =
  for line in lines(path):
    callback(line)

proc getData(path: string): seq[seq[int]] =
  var data: seq[seq[int]] = @[]
  handle(path) do(line: string):
    let parts = line.split().mapIt(it.parseInt)
    data.add(parts)
  return data

proc isSafe(data: seq[int], skip: int = -1): bool =
  var increasing: Option[bool] = none(bool)
  var isSafe = true

  var i = if skip == 0: 1 else: 0
  while i < data.len() - 1:
    if skip == i:
      i += 1

    let right =
      if skip == i + 1:
        i + 2
      else:
        i + 1

    if right > data.len() - 1:
      break

    let diff = data[right] - data[i]
    let inc = if diff <= 0: false else: true

    if increasing.isNone():
      increasing = some(inc)
    elif increasing.isSome() and increasing.get() != inc:
      isSafe = false
      break

    if abs(diff) > 3 or diff == 0:
      isSafe = false
      break
    i += 1
  return isSafe

proc part1(data: seq[seq[int]]) =
  var count = 0
  for nums in data:
    let isSafe = isSafe(nums)
    if isSafe:
      count += 1
  echo("Part 1: ", count)

proc tryIsSafe(data: seq[int]): bool =
  var i = -1
  while i < data.len():
    let isSafe = isSafe(data, i)
    if isSafe:
      return true
    i += 1
  return false

proc part2(data: seq[seq[int]]) =
  var count = 0
  for nums in data:
    let isSafe = tryIsSafe(nums)
    if isSafe:
      count += 1
  echo("Part 2: ", count)

when isMainModule:
  let path = getPath()
  let data = getData(path)
  part1(data)
  part2(data)
