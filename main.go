package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var s string
	switch len(os.Args) {
	case 1: // no args supplied
		// check stdin
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			b, err := io.ReadAll(os.Stdin)
			if err != nil {
				log.Fatalf(err.Error())
			}

			s = strings.TrimSuffix(string(b), "\n")
		} else { // not stdin
			log.Fatalf("must supply an epoch time argument or from stdin")
		}

	default:
		s = os.Args[1]
	}

	t, err := UnixEpochToTime(s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", t.Format(time.RFC3339))
}

// UnixEpochToTime converts a variety of unix epochs to go times
func UnixEpochToTime(s string) (time.Time, error) {
	// check if number
	_, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, nil
	}

	var secStr string
	var nsecStr string

	// handle long epochs, microsecond resolution
	if len(s) > 10 {
		secStr = s[:10]
		nsecStr = s[10:]
	} else {
		secStr = s
		nsecStr = "0"
	}

	sec, err := strconv.ParseInt(secStr, 10, 64)
	if err != nil {
		return time.Time{}, nil
	}

	nsec, err := strconv.ParseInt(nsecStr, 10, 64)
	if err != nil {
		return time.Time{}, nil
	}

	t := time.Unix(sec, nsec)
	return t, nil
}
