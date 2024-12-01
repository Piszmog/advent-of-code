type HandleCallback = proc(line: string) {.closure}

proc handle(path: string, callback: HandleCallback) =
  for line in lines(path):
    callback(line)

export handle
