package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/aaaasmile/gpxtidy/gpx"
)

func main() {

	var dir = flag.String("dir", "", "Data dierctory")
	var source = flag.String("source", "", "Source gpx file")
	var target = flag.String("target", "", `Destination gpx`)
	var cmd = flag.String("cmd", "", `Commands available: remext`)
	var ver = flag.Bool("ver", false, "Program version")
	flag.Parse()
	if *ver {
		log.Println("gpxtidy - Version is 0.001.00 20220101")
		return
	}

	if *dir == "" || *source == "" || *target == "" || *cmd == "" {
		fmt.Println(`Process gpx file for removing parts.
		For example use something like this:
		gpx -cmd remext -dir .\data -source S1_12Ipertrail_2022-v221227.GPX  -target noext_S1_20211227_iper.gpx`)
		return
	}

	start := time.Now()
	c := gpx.NewCommander(*cmd, *dir, *source, *target)
	if err := c.Run(); err != nil {
		log.Fatalln("Something went wrong: ", err)
	}
	log.Printf("Time elapsed: %s\n", time.Since(start).String())
	log.Println("That's all folks!")
}
