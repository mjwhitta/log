# log

[![Yum](https://img.shields.io/badge/-Buy%20me%20a%20cookie-blue?labelColor=grey&logo=cookiecutter&style=for-the-badge)](https://www.buymeacoffee.com/mjwhitta)

[![Go Report Card](https://goreportcard.com/badge/github.com/mjwhitta/log?style=for-the-badge)](https://goreportcard.com/report/github.com/mjwhitta/log)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/mjwhitta/log/ci.yaml?style=for-the-badge)](https://github.com/mjwhitta/log/actions)
![License](https://img.shields.io/github/license/mjwhitta/log?style=for-the-badge)

## What is this?

A simple, extensible logger package. Maybe you want a logger that logs
to a file and also logs to a websocket. Maybe you want to log
everything to a file but only portions of certain messages to STDOUT.
Maybe you don't want to log to STDOUT at all, but rather to a TUI.
Using the `CustomLogHandler` functionality of this package, you can do
anything you want.

## How to install

Open a terminal and run the following:

```
$ go get -u github.com/mjwhitta/log
```

## Usage

```
package main

import (
    "fmt"

    "github.com/mjwhitta/log"
)

var logger *log.Messenger

func main() {
    var e error

    // Default log functionality (stdout w/o timestamp)
    log.Debug("Debug message")
    log.Info("Info message")
    log.SubInfo("SubInfo message")
    log.Good("Good message")
    log.Warn("Warn message")
    log.Err("Error message")

    // Default log functionality + timestamp
    log.Timestamp = true
    log.Debug("Debug message")
    log.Info("Info message")
    log.SubInfo("SubInfo message")
    log.Good("Good message")
    log.Warn("Warn message")
    log.Err("Error message")

    // Will log to stdout (w/o timestamp)
    logger = log.NewMessenger()
    logger.Debug("Debug message")
    logger.Info("Info message")
    logger.SubInfo("SubInfo message")
    logger.Good("Good message")
    logger.Warn("Warn message")
    logger.Err("Error message")

    // Will now log to stdout (w/o timestamp) and file (w/ timestamp)
    if logger, e = log.NewFileMessenger("/tmp/test.log"); e != nil {
        panic(e)
    }

    logger.Debug("Debug message")
    logger.Info("Info message")
    logger.SubInfo("SubInfo message")
    logger.Good("Good message")
    logger.Warn("Warn message")
    logger.Err("Error message")

    // Will now log to stdout (w/ timestamp) and file (w/ timestamp)
    log.Prefixes[log.TypeDebug] = "[DEBU]"
    log.Prefixes[log.TypeErr] = "[ERRO]"
    log.Prefixes[log.TypeGood] = "[GOOD]"
    log.Prefixes[log.TypeInfo] = "[INFO]"
    log.Prefixes[log.TypeSubInfo] = "[SUBI]"
    log.Prefixes[log.TypeWarn] = "[WARN]"
    logger.Timestamp = true
    logger.Debug("Debug message")
    logger.Info("Info message")
    logger.SubInfo("SubInfo message")
    logger.Good("Good message")
    logger.Warn("Warn message")
    logger.Err("Error message")

    // Disable color on stdout
    logger.SetColor(false)
    logger.Debug("Debug message")
    logger.Info("Info message")
    logger.SubInfo("SubInfo message")
    logger.Good("Good message")
    logger.Warn("Warn message")
    logger.Err("Error message")

    // Custom MsgHandler
    logger.SetMsgHandler(
        func(msg *log.Message) error {
            switch msg.Type {
            case log.TypeDebug:
                fmt.Println("Custom 1 - debug")
            case log.TypeErr, log.TypeErrX:
                fmt.Println("Custom 1 - error")
            case log.TypeGood:
                fmt.Println("Custom 1 - good")
            case log.TypeInfo:
                fmt.Println("Custom 1 - info")
            case log.TypeMsg:
                fmt.Println("Custom 1 - message")
            case log.TypeSubInfo:
                fmt.Println("Custom 1 - additional info")
            case log.TypeWarn:
                fmt.Println("Custom 1 - warning")
            }
            return nil
        },
    )
    logger.AddMsgHandler(
        func(msg *log.Message) error {
            fmt.Println("Custom 2")
            return nil
        },
    )
    logger.Debug("Debug message")
    logger.Info("Info message")
    logger.SubInfo("SubInfo message")
    logger.Good("Good message")
    logger.Warn("Warn message")
    logger.Err("Error message")

    // Close logger
    logger.AddCloseHandler(
        func() error {
            fmt.Println("Closed")
            return nil
        },
    )
    if e = logger.Close(); e != nil {
        panic(e)
    }
}
```

## Links

- [Source](https://github.com/mjwhitta/log)
