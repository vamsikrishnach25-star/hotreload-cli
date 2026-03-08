# HotReload CLI

HotReload CLI is a lightweight command-line tool written in Go that automatically rebuilds and restarts a Go server whenever source files change.

During development, restarting a server after every code change can slow down the workflow. This tool solves that problem by watching the project files and automatically rebuilding and restarting the server whenever changes are detected. It allows developers to focus on writing code instead of manually restarting the application.

---

## Features

- Recursive directory watching
- Automatic rebuild when files change
- Automatic server restart
- Debounce mechanism to prevent multiple rebuilds during rapid file changes
- Real-time server log streaming
- Process tree termination for stubborn processes
- Ignores unnecessary folders such as .git, node_modules, and build artifacts

---

## Project Structure

hotreload-cli
│
├── main.go
├── go.mod
│
├── internal
│   ├── watcher
│   ├── builder
│   └── runner
│
└── testserver
    └── main.go

---

## Running the Demo

Run the hot reload tool using the following command:

go run main.go --root ./testserver \
--build "go build -o server.exe ./testserver" \
--exec "server.exe"

After running the command, edit the file `testserver/main.go`.

The tool will detect the changes, rebuild the project, and restart the server automatically.

---

## Architecture

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

---

## Demo Server

The repository includes a simple HTTP server inside the `testserver` directory that demonstrates the hot reload functionality. The tool monitors this server and automatically rebuilds and restarts it whenever source files are modified.

---

## Technologies Used

- Go
- fsnotify
- os/exec
- Go Standard Library
