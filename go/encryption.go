package main
import ("fmt"
	"cmd/asm/internal/flags"
		"os"
		"strings"
		"flag"
)

// this is a simple encryption/decryption program
// using a simple substituion cipher


func encryption_main () {
	if len(os.Args) <= 1 {
		fmt.Println("Hey you haven't provided enough arguments - as reminder include a word and the type e/d")
	}

	cipher := shuffle("abcdefghijklmnopqrstuvqxyz")
	
	wordPtr := flag.String("str", "", "The String that you want encrypted/descrypted.")
	typePtr := flag.String("str", "e", "e for encryption, d for decryption, default to e.")
	
	flags.Parse()

	if *typePtr == "e"{
		encrypt(*wordPtr)
	} else { 
		decrypt(*wordPtr)
	}
}

func encrypt(str string) {
	lowered = strings.ToLower(str)

}

func decrypt(str string) {
	lowered = strings.ToLower(str)

}

func shuffle(word string)  {
	split := strings.Split(word, "")
	rand.Seed(time.Now().UTC().UnixNano())

	shuffled = make([]string, len(split))
	perm := rand.Perm(len(split))

	for i, v := range perm {
		shuffled[v] = split[i]
	}

	return shuffled
}


