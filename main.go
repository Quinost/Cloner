package main

import (
	"cloner/filemanipulator"
	"fmt"
	"os"
)

func main() {
	os.Stdin.Close()

	settings := filemanipulator.ReadSettings()

	filemanipulator.DeleteUnnecessary(settings)

	filemanipulator.Copy(&settings)

	fmt.Println("Kopiowanie zakończone")
	fmt.Scanln()
}

//GOOS=windows GOARCH=amd64 go build -o moja_aplikacja.exe
