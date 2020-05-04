# log

## What is this?

A simple logger package.

## How to install

Open a terminal and run the following:

```
$ go get -u gitlab.com/mjwhitta/log
```

## Usage

```
package main

import (
    hl "gitlab.com/mjwhitta/hilighter"
    "gitlab.com/mjwhitta/log"
)

var logger *log.Messenger

func main() {
    var e error

    // Default log functionality (stdout w/o timestamp)
    log.Info("Info message")
    log.Good("Good message")
    log.Err("Error message")

    // Default log functionality + timestamp
    log.Timestamp = true
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

    // CustomLogHandler
    logger.SetLogHandler(
        func(msg log.Message) error {
            switch msg.Type {
            case log.TypeErr:
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
    logger.AddLogHandler(
        func(msg log.Message) error {
            hl.Println("Custom 2")
            return nil
        },
    )
    logger.Info("Info message")
    logger.Good("Good message")
    logger.Err("Error message")

    // Close logger
    if e = logger.Close(); e != nil {
        panic(e)
    }
}
```

## Links

- [Source](https://gitlab.com/mjwhitta/log)
