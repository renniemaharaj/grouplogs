# log

A lightweight, thread-safe Go logging package with support for structured log entries, real-time subscriptions via channels, and a grouping mechanism for broadcasting logs from multiple sources. Ideal for building reactive systems and piping logs into websockets or monitoring services.

## Features

- Thread-safe subscriber management using `sync.Mutex`
- Per-subscriber channels (`chan Line`)
- Grouping of multiple loggers via `logger.Group()`
- Real-time log broadcasting
- Log chaining: `logger.New()... .Info("starting...").Debug("ready")`
- Runtime caller metadata: file, line number, and function name

---

## Installation

```bash
go get github.com/renniemaharaj/grouplogs

import "github.com/renniemaharaj/grouplogs/pkg/logger"
```

---

## Building a logger

```go
l := logger.New().
	Prefix("Primary").
	Debugging(true).
	JSONMode(false).
	Subscribable(true).
	MaxLines(100).
	STDOUT(true).
	Rotate()
```

---

## Logging methods

```go
l.Info("This is an information").
	Success("This is a success").
	Warning("This is a warning").
	Debug("Is debugging enabled").
	Error("Oh, no. This is an error")
```

---

## Grouping loggers

```go
group := logger.Group()

l1 := logger.New().Prefix("L1").Subscribable(true).Rotate()
l2 := logger.New().Prefix("L2").Subscribable(true).Rotate()

group.Join(l1)
group.Join(l2)

// Now any logs from l1 or l2 will be piped into group.Delegate
```

---

## Real-time log piping to WebSocket

```go
func logHandler(con *websocket.Conn) {
	// Create a group to pipe logs
	group := logger.Group()

	// Add one or more loggers
	l := logger.New().Prefix("WS").Subscribable(true).Rotate()
	group.Join(l)

	// Example logs
	l.Info("WebSocket log stream initialized")

	for {
		select {
		case entry := <-group.Delegate:
			logArr := &[]logger.Line{entry}

			logBytes, err := json.Marshal(logArr)
			if err != nil {
				break
			}

			if err := con.WriteMessage(websocket.TextMessage, logBytes); err != nil {
				group.Remove(l)
				l.Warning("WebSocket connection closed")
				return
			}
		}
	}
}
```

---

## Use Cases

- Centralized log collection from multiple services
- Real-time debug dashboards
- WebSocket log streaming to browser clients
- Embedded monitoring in custom Go applications

---

## License

MIT
