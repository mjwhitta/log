# log

[![Yum](https://img.shields.io/badge/-Buy%20me%20a%20cookie-blue?labelColor=grey&logo=cookiecutter&style=for-the-badge)](https://www.buymeacoffee.com/mjwhitta)

[![Go Report Card](https://goreportcard.com/badge/github.com/mjwhitta/log?style=for-the-badge)](https://goreportcard.com/report/github.com/mjwhitta/log)
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
    hl "github.com/mjwhitta/hilighter"
    "github.com/mjwhitta/log"
)

var logger *log.Messenger

func main() {
    var e error

    // Default log functionality (stdout w/o timestamp)
    log.Debug("Debug message")
    log.Info("Info message")
    log.Good("Good message")
    log.Err("Error message")

    // Default log functionality + timestamp
    log.Timestamp = true
    log.Debug("Debug message")
    log.Info("Info message")
    log.Good("Good message")
    log.Err("Error message")

    // Will log to stdout (w/o timestamp)
    logger = log.NewMessenger()
    logger.Info("Info message")
    logger.Good("Good message")
    logger.Err("Error message")

    // Will now log to stdout (w/o timestamp) and file (w/ timestamp)
    if logger, e = log.NewFileMessenger("/tmp/test.log"); e != nil {
        panic(e)
    }
    logger.Info("Info message")
    logger.Good("Good message")
    logger.Err("Error message")

    // Will now log to stdout (w/ timestamp) and file (w/ timestamp)
    logger.Timestamp = true
    logger.Info("Info message")
    logger.Good("Good message")
    logger.Err("Error message")

    // Disable color on stdout
    logger.SetColor(false)
    logger.Info("Info message")
    logger.Good("Good message")
    logger.Err("Error message")

    // Custom MsgHandler
    logger.SetMsgHandler(
        func(msg *log.Message) error {
            switch msg.Type {
            case log.TypeDebug:
                hl.Println("Custom 1 - debug")
            case log.TypeErr, log.TypeErrX:
                hl.Println("Custom 1 - error")
            case log.TypeGood:
                hl.Println("Custom 1 - good")
            case log.TypeInfo:
                hl.Println("Custom 1 - info")
            case log.TypeMsg:
                hl.Println("Custom 1 - message")
            case log.TypeSubInfo:
                hl.Println("Custom 1 - additional info")
            case log.TypeWarn:
                hl.Println("Custom 1 - warning")
            }
            return nil
        },
    )
    logger.AddMsgHandler(
        func(msg *log.Message) error {
            hl.Println("Custom 2")
            return nil
        },
    )
    logger.Debug("Debug message")
    logger.Info("Info message")
    logger.Good("Good message")
    logger.Err("Error message")

    // Close logger
    logger.AddCloseHandler(
        func() error {
            hl.Println("Closed")
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
