package main
import ("fmt"
		"os"
		"strings"
		"flag"
		"math/rand"
		"time"
)

// this is a simple encryption/decryption program
// using a simple substituion cipher


func main () {
	if len(os.Args) <= 1 {
		fmt.Println("Hey you haven't provided enough arguments - as reminder include a word and the type e/d")
	}

	regular_alpha := "abcdefghijklmnopqrstuvqxyz"
	cipher := shuffle(regular_alpha)

	fmt.Println("Your Cipher is: ", cipher)
	
	regular_map := make(map[rune]int)
	cipher_map := make(map[rune]int)

	for pos, char := range cipher{
		cipher_map[char] = pos
	}

	for pos, char := range regular_alpha{
		regular_map[char] = pos
	}

	
	
	wordPtr := flag.String("str", "", "The String that you want encrypted/descrypted.")
	typePtr := flag.String("type", "e", "e for encryption, d for decryption, default to e.")
	
	flag.Parse()

	if *typePtr == "e"{
		simple_cipher(*wordPtr, cipher_map, regular_alpha)
	} else { 
		simple_cipher(*wordPtr, regular_map, cipher)
	}

}

func simple_cipher(str string, regular map[rune]int, cipher string) {
	lowered := strings.ToLower(str)
	var translated strings.Builder
	
	for _, char := range lowered {
		corresponding_letter := cipher[regular[char]]
		translated.WriteString(string(corresponding_letter))
	}

	fmt.Printf("Your translated string is: %s \n", translated.String())

}

func shuffle(word string) string{
	split := strings.Split(word, "")
	rand.Seed(time.Now().UTC().UnixNano())

	shuffled := make([]string, len(split))
	perm := rand.Perm(len(split))

	for i, v := range perm {
		shuffled[v] = split[i]
	}

	shuffled_string := strings.Join(shuffled,"")
	return shuffled_string
}


