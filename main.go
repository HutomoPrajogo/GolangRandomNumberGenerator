package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Random Number Generator")
	fmt.Println("---------------------")

	for {
		fmt.Print("Maximum Number: ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Invalid Number")
		}

		seed := time.Now().Unix() % 10
		rand.Seed(int64(seed))
		fmt.Println(rand.Intn(num))

	}
}
