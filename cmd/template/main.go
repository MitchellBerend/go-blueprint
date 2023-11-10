// Package template provides utility functions that
// help with the templating of created files.
package template

// MakeHTTPRoutes returns a byte slice that represents 
// the default cmd/api/main.go file template.
func MainTemplate() []byte {
	return []byte(`package main

import (
	"{{.ProjectName}}/internal/server"
)

func main() {

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
`)
}

// WebsocketTemplate returns a slice that represents the logic of the websocket
// handler. This function expects some variables set in the template. See echo
// router for more details.
//
// w http.ResponseWriter
// r *http.Request
func websocketTemplate() []byte {
  return []byte(`
	errorMessage := []byte("This is another message not PING")
	socket, err := websocket.Accept(w, r, nil)

	if err != nil {
		log.Print("could not open websocket")
		_, _ = w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx := r.Context()
	for {
		msgType, socketBytes, err := socket.Read(ctx)

		if err != nil {
			log.Print("could not read from websocket")
			return
		}

		if string(socketBytes) == "PING" {
			if err := socket.Write(ctx, msgType, []byte("PONG")); err != nil {
				log.Print("could not write to socket")
				return
			}
			continue
		}

		if err := socket.Write(ctx, msgType, errorMessage); err != nil {
			log.Print("could not write to socket")
			return
		}
	}
`)
}

// MakeHTTPRoutes returns a byte slice that represents 
// the default Makefile.
func MakeTemplate() []byte {
	return []byte(
		`
# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./...

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@echo "Watching..."
	@air

.PHONY: all build run test clean
		`)
}


func AirTomlTemplate() []byte {
	return []byte(
		`
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = []
  bin = "./tmp/main"
  cmd = "make run"
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]
  exclude_file = []
  exclude_regex = ["_test.go"]
  exclude_unchanged = false
  follow_symlink = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "tpl", "tmpl", "html"]
  include_file = []
  kill_delay = "0s"
  log = "build-errors.log"
  poll = false
  poll_interval = 0
  post_cmd = []
  pre_cmd = []
  rerun = false
  rerun_delay = 500
  send_interrupt = false
  stop_on_error = false

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  main_only = false
  time = false

[misc]
  clean_on_exit = false

[screen]
  clear_on_rebuild = false
  keep_scroll = true
        `)
}


// ReadmeTemplate returns a byte slice that represents 
// the default README.md file template.
func ReadmeTemplate() []byte {
	return []byte(
		`
# Project {{.ProjectName}}

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

run all make commands with clean tests
` + "```bash" + `
make all build
` + "```" + `

build the application
` + "```bash" + `
make build
` + "```" + `

run the application
` + "```bash" + `
make run
` + "```" + `

run the test suite
` + "```bash" + `
make test
` + "```" + `

clean up binary from the last build
` + "```bash" + `
make clean
` + "```" + `
	`)
}
