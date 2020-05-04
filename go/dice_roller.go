package main

import ("fmt"
		"bufio"
		"os"
		"strconv"
		"math/rand"
		"time"
		"strings"
)

func main() {
	fmt.Println("Hello! You are now running rice roller!\n")

	// inside functions you can use := to decalre a varaible with implicit type
	// that way you don't need to use var
	// this only works inside functions
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter how many dice you would like rolled!")

	// since the rand.Intn function aways generates the same numbers everytime 
	// you run - we want to generate a random see to get true random numbers
	rand.Seed(time.Now().UnixNano()) // UnixNano is unix time
	min := 1
	max := 6

	for scanner.Scan() {
		// fmt.Println(reflect.TypeOf(scanner.Text()))
		text := scanner.Text()

		num, err := strconv.Atoi(text)
		if err != nil {
			if  strings.ToLower(text) == "quit" {
				os.Exit(0)
			}

			fmt.Println("HEY!! That's not a number! Try again!")
			continue
		} 
		for i := 0; i < num; i++ {
			dice_roll := rand.Intn(max - min + 1) + min 
			//since Intn only does from 0 -> n

			str := fmt.Sprintf("Dice Roll %d: %d", i + 1, dice_roll)
			fmt.Println(str)
		}
		fmt.Println("To exit this program, type quit.")


	}

}
