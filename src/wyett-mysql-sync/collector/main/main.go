package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/23 9:21
 * @description: main
 */

// code
type exitCode struct {
	Code int
}

func handleExit() {
	if e := recover(); e != nil {
		if exit, ok := e.(exitCode); ok {
			if exit.Code != 0 {
				fmt.Fprintln(os.Stderr, "wyett-mysql-sync failed at", time.Now().Format("January 2, "+
					"2006 at 3:04pm (MST)"))
			} else {
				fmt.Fprintln(os.Stderr, "Stopped syncing at", time.Now().Format("January 2, 2006 at 3:04pm (MST)"))
			}
			os.Exit(exit.Code)
		}
		panic(e)
	}
}

func main() {
	var err error
	defer handleExit()

	configuration := flag.String("--config", "configuration", "config file name")
	version := flag.Bool("version", false, "check version")

	// 1. check if no configuration file
	if *configuration == "" && *version == true {
		crash("config file does not exist", exitCode{Code: -1})
	}

}

func crash(msg string, code exitCode) {
	log.Fatalf("%s with fatal code %v", msg, code)
}

func usage() error {
	fmt.Printf("--help")
}
