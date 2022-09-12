package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.Open("activeDomains12.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	temp := ""

	scanner := bufio.NewScanner(file)

	br := 0
	for scanner.Scan() {

		temp2 := strings.Split(scanner.Text(), ".")

		if temp2[0] == temp {

			br++
		} else {
			if br > 4 {
				fmt.Println(temp, br)
			}

			temp = temp2[0]
			br = 1
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	elapsed1 := time.Since(start)
	fmt.Println("read took", elapsed1)
}
