package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// example of working with files in an infinite loop
	for {
		file, err := os.Open("./README.md")

		if err != nil {
			fmt.Println(err)
			return
		}

		func() {
			defer file.Close()
			// some actions with the file can be placed here
		}()
		fmt.Println(file.Close())
		time.Sleep(time.Second * 1)
	}
}
