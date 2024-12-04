import strutils
import re
import sequtils
import os

proc getPath(): string =
  if commandLineParams().len() <= 0:
    echo "Missing argument to specify the input file"
    quit(1)
  return paramStr(1)

type HandleCallback = proc(line: string) {.closure.}

proc handle(path: string, callback: HandleCallback) =
  for line in lines(path):
    callback(line)

proc part1(path: string) =
  var total = 0
  handle(path) do(line: string):
    let mulMatches = findAll(line, re"mul\(\d+,\d+\)")
      .mapIt(it.split({'(', ')', ','}))
      .mapIt(it.filterIt(it != "mul" and it != ""))
      .mapIt(it.mapIt(it.parseInt))
    for m in mulMatches:
      total += m[0] * m[1]
  echo("Part 1: ", total)

proc part2(path: string) =
  var total = 0
  var enabled = true
  handle(path) do(line: string):
    let mulMatches = findAll(line, re"mul\(\d+,\d+\)|do\(\)|don't\(\)")
      .mapIt(it.split({'(', ')', ','}))
      .mapIt(it.filterIt(it != "mul" and it != ""))
    for m in mulMatches:
      if m.len() == 1:
        enabled =
          case m[0]
          of "do": true
          else: false
      elif enabled:
        total += m[0].parseInt() * m[1].parseInt()

  echo("Part 2: ", total)

when isMainModule:
  let path = getPath()
  part1(path)
  part2(path)
