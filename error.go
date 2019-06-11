package main

import "log"

func errcheck(_err *error) {
	if *_err != nil {
		log.Print(*_err)
	}
}
