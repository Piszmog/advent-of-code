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

when isMainModule:
  let path = getPath()

