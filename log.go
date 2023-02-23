package main

import (
	"fmt"
	"strconv"
)

func R(msg string) {
	report(msg, 30)
}

func E(msg string) {
	report(msg, 31)
}

func S(msg string) {
	report(msg, 32)
}

func I(msg string) {
	report(". "+msg, -1)
}

func D(msg string) {
	report("$$ "+msg, -1)
}

func report(msg string, color int) {
	if msg != "" && len(msg) >= 1 {
		if color <= 0 {
			fmt.Println(msg)
		} else {
			fmt.Println("\033[" + strconv.Itoa(color) + "m" + msg + "\033[0m")
		}
	}
}
