package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {

	tm := flag.String("time", "", "time to convert. ex 15:04")
	tzp := flag.String("tzo", "", "Time zone location origin")

	flag.Parse()

	var ll, err = time.LoadLocation(*tzp)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range tz {

		loc, err := time.LoadLocation(v)

		if err != nil {
			log.Fatalln(err)
		}
		t, err := time.ParseInLocation("15:04", *tm, ll)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(t.In(loc).Format("15:04"), t.In(loc).Location())
	}
}
