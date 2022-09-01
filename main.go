package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/adhocore/gronx"
)

const expr = "0 * 9 LW JAN-OCT 1-5 2000/10"

var (
	usage = func() {
		log.Printf("usage:\n  %s [options] \"{cron expression}\"\noptions:\n", os.Args[0])
		flag.PrintDefaults()
	}
	inTimeStr  string
	cronString string
)

func main() {
	flag.Usage = usage
	//	flag.StringVar(&inTimeStr, "t", "", `whole or partial RFC3339 time value (i.e. "2006-01-02T15:04:05Z07:00") against which the cron expression is evaluated, now if not present`)
	flag.StringVar(&cronString, "c", "", `0 * 9 LW JAN-OCT 1-5 2000/10`)
	flag.Parse()

	gron := gronx.New()

	b, err := gron.IsDue(cronString)
	if err != nil {
		log.Printf("Err: %v\n", err)
	}
	if b {
		log.Printf("Mached: %s at %s\n", cronString, time.Now().String())
	}

}
