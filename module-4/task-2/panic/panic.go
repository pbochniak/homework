package panic

import "fmt"

func iWillPanic() {
	panic("something")
}

func recoveryPanic() {
	if err := recover(); err != nil {
		fmt.Printf("recovery from panic - %q\n", err)
	}
}

func catchPanic() {
	defer recoveryPanic()
	iWillPanic()
}
