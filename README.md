# winpush - Windows Toast Notifications for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/yourusername/winpush.svg)](https://pkg.go.dev/github.com/yourusername/winpush)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A Go package for creating Windows Toast notifications without visible PowerShell windows.

## Features

- 🛠️ Customizable notifications (title, message, icon)
- 🔘 Action buttons support
- ⏱️ Duration control
- 🔗 Activation arguments for click handling
- 🖼️ Adaptive templates for Windows 10/11
- 🕶️ Silent execution (no pop-up windows)

 ## Requirements

- Windows 10/11
- Go 1.16+
- Notifications enabled for the app in Windows settings
- PowerShell 5.0+ (usually pre-installed)

## Installation

```bash
go get github.com/yourusername/winpush
```

## Quick Start

### Basic Notification
```go
package main

import "github.com/yourusername/winpush"

func main() {
    notifier := winpush.Notificator{
        Title:   "Hello World",
        Message: "This is a toast notification",
    }
    err := notifier.Push()
    if err != nil {
        panic(err)
    }
}
```

### Notification with Actions
```go
notifier := winpush.Notificator{
    Title:   "Download Complete",
    Message: "file.zip has been downloaded successfully",
    Icon:    "C:\\path\\to\\icon.png",
    Duration: "long",
    Actions: []winpush.Actions{
        {
            Content:   "Open",
            Arguments: "open:file.zip",
        },
        {
            Content:   "Show in folder",
            Arguments: "explorer:C:\\downloads",
            Icon:      "folder.ico",
        },
    },
}
err := notifier.Push()
```

## Configuration

### Notificator Structure
```go
type Notificator struct {
    AppID               string      // Application ID
    Title               string      // Main title
    Subtitle            string      // Subtitle
    Message             string      // Notification body
    Icon                string      // Path to icon (PNG, JPG, ICO)
    Actions             []Actions   // List of actions
    ActivationType      string      // Activation type ("protocol" by default)
    ActivationArguments string      // Activation arguments
    Duration            string      // Duration: "short" (default) or "long"
}
```

### Actions Structure
```go
type Actions struct {
    Content        string // Button text
    Arguments      string // Action arguments
    Icon           string // Button icon (optional)
    ActivationType string // Activation type ("protocol")
    Placement      string // Position ("contextMenu")
}
```

## Error Handling

Possible errors:
```go
ErrExecuteToast = errors.New("winPush: error execute toast")
ErrCreateFile   = errors.New("winPush: error create file for xml push")
ErrReadXML      = errors.New("winPush: error read xml for push")
```

Example error handling:
```go
err := notifier.Push()
if err != nil {
    switch {
    case errors.Is(err, winpush.ErrCreateFile):
        log.Fatal("Failed to create temp file")
    case errors.Is(err, winpush.ErrExecuteToast):
        log.Fatal("Failed to execute PowerShell script")
    default:
        log.Fatal("Unknown error:", err)
    }
}
```
