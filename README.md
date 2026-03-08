# HotReload CLI

A lightweight Go CLI tool that automatically rebuilds and restarts a Go server whenever source files change.

This tool improves developer productivity by eliminating the need to manually stop, rebuild, and restart the server after each code modification.

---

## Features

* Recursive directory watching
* Automatic rebuild on file changes
* Automatic server restart
* Debounce mechanism to avoid multiple rebuilds
* Real-time server log streaming
* Process tree termination for stubborn processes
* Ignore unnecessary folders (`.git`, `node_modules`, build artifacts)

---

## Project Structure

```
hotreload-cli/
│
├── main.go
├── go.mod
│
├── internal/
│   ├── watcher/   # filesystem watcher
│   ├── builder/   # build command execution
│   └── runner/    # process management
│
└── testserver/
    └── main.go    # sample HTTP server
```

---

## Running the Demo

Run the hot reload tool:

```
go run main.go --root ./testserver \
--build "go build -o server.exe ./testserver" \
--exec "server.exe"
```

Then edit `testserver/main.go`.

The server will automatically rebuild and restart.

---

## Architecture

```
CLI
 ↓
Argument Parsing
 ↓
File Watcher (fsnotify)
 ↓
Debounce Layer
 ↓
Build Engine
 ↓
Process Manager
 ↓
Server Execution
```

---

## Demo Server

The repository includes a simple HTTP server inside `testserver/` used to demonstrate the hot reload functionality.

---

## Technologies Used

* Go
* fsnotify
* os/exec
* Go standard library
