package main

import (
	"fmt"
	"log"
	"os"

	"github.com/denisbrodbeck/machineid"
)

var path = "logs.txt"

//Main function
func main() {
	fmt.Println("HWID-Checker used by Charged Development")

	uuid, err := machineid.ID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Your HWID is :", uuid)
	createFile()
	writeFile(uuid)
}

//File creation
func createFile() {
	// check if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
	}

	fmt.Println("File Created Successfully", path)
}

//Write function
func writeFile(uuid string) {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Write HWID to file
	_, err = file.WriteString("Your HWID is: " + uuid)
	if err != nil {
		fmt.Println(err)
	}

	// Save file changes.
	err = file.Sync()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Check logs.txt for your HWID to copy and paste!")
	fmt.Printf("Press enter to exit...")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}
