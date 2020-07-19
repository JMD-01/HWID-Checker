package main

//Credits to https://github.com/atotto/clipboard/blob/master/clipboard_windows.go
//Credits to https://github.com/denisbrodbeck/machineid

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/denisbrodbeck/machineid"
)

var (
	pasteCmdArgs = "pbpaste"
	copyCmdArgs  = "pbcopy"
)

func getCopyCommand() *exec.Cmd {
	return exec.Command(copyCmdArgs)
}

//Main function
func main() {
	fmt.Println("HWID-Checker")
	uuid, err := machineid.ID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Your HWID is :", uuid)
	writeAll(uuid)
	fmt.Println("Copied to your clipboard!")
	//Prevent auto closing
	b := make([]byte, 1)
	os.Stdin.Read(b)

}

func writeAll(text string) error {
	copyCmd := getCopyCommand()
	in, err := copyCmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := copyCmd.Start(); err != nil {
		return err
	}
	if _, err := in.Write([]byte(text)); err != nil {
		return err
	}
	if err := in.Close(); err != nil {
		return err
	}
	return copyCmd.Wait()
}
