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

import "gitlab.com/mjwhitta/log"

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

    // Close logger
    if e = logger.Close(); e != nil {
        panic(e)
    }
}
```

## Links

- [Source](https://gitlab.com/mjwhitta/log)
