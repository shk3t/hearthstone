return {
  run = {
    -- go = "go build -o build/hearthstone cmd/main.go; ./build/hearthstone"
    go = "go run cmd/main.go",
  },
  debug = {
    go = {
      {
        type = "delve",
        request = "launch",
        name = "Default",
        program = "cmd/main.go",
      },
    },
  },
}
