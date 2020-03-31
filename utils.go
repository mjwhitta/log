package log

import (
	"os"
	"time"

	hl "gitlab.com/mjwhitta/hilighter"
)

func logToFile(msg string, stamp bool, f ...*os.File) {
	var ts = time.Now().Format(time.RFC3339) + ": "

	if stamp {
		hl.Println(ts + msg)
	} else {
		hl.Println(msg)
	}

	if (len(f) == 0) || (f[0] == nil) {
		return
	}

	f[0].WriteString(ts + hl.Plain(msg) + "\n")
}
