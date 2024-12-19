package main

import (
	"cloner/filemanipulator"
	"fmt"
	"log"

	"github.com/nsf/termbox-go"
)

func main() {

	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	go func() {
		settings := filemanipulator.ReadSettings()

		filemanipulator.DeleteUnnecessary(settings)

		filemanipulator.Copy(&settings)

		fmt.Println("Kopiowanie zakończone")
		fmt.Scanln()
	}()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				return
			}
		case termbox.EventMouse:
			//ignore mouse events
		}
	}
}

//GOOS=windows GOARCH=amd64 go build -o moja_aplikacja.exe