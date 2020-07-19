package main

//Credits to https://github.com/atotto/clipboard
//Credits to https://github.com/denisbrodbeck/machineid
import (
	"fmt"
	"log"
	"os"

	"github.com/atotto/clipboard"
	"github.com/denisbrodbeck/machineid"
)

//Main function
func main() {
	fmt.Println("HWID-Checker")
	uuid, err := machineid.ID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Your HWID is :", uuid)
	clipboard.WriteAll(uuid)
	fmt.Println("Copied to your clipboard!")
	//Prevent auto closing
	b := make([]byte, 1)
	os.Stdin.Read(b)

}
