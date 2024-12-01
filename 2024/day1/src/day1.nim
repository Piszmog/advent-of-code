import strutils
import sequtils
import os
import std/algorithm

proc getPath(): string =
  if commandLineParams().len() <= 0:
    echo "Missing argument to specify the input file"
    quit(1)
  return paramStr(1)

type HandleCallback = proc(line: string) {.closure}

proc handle(path: string, callback: HandleCallback) =
  for line in lines(path):
    callback(line)

proc getData(path: string): (seq[int], seq[int]) =
  var left: seq[int] = @[]
  var right: seq[int] = @[]
  handle(path) do (line: string):
    let parts = line.splitWhitespace()
    left.add(parts[0].parseInt())
    right.add(parts[1].parseInt())
  
  left.sort()
  right.sort()

  return (left, right)

proc day1(left: seq[int], right: seq[int]) =
  var total = 0
  for i in 0..left.len()-1:
    let dist = right[i] - left[i]
    total += abs(dist)

  echo("Part 1: ", total)

proc day2(left: seq[int], right: seq[int]) =
  var total = 0
  for i in 0..left.len()-1:
    let num = left[i]
    let count = right.count(num)
    let sim = left[i] * count
    total += sim

  echo("Part 2: ", total)

when isMainModule:
  let path = getPath()
  let (left, right) = getData(path)
  day1(left, right)
  day2(left, right)
