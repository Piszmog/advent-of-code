import strutils
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

proc day1(path: string) =
  var left: seq[int] = @[]
  var right: seq[int] = @[]
  handle(path) do (line: string):
    let parts = line.splitWhitespace()
    left.add(parts[0].parseInt())
    right.add(parts[1].parseInt())
  
  left.sort()
  right.sort()

  var total = 0
  for i in 0..left.len()-1:
    let dist = right[i] - left[i]
    total += abs(dist)

  echo(total)

when isMainModule:
  let path = getPath()
  day1(path)
