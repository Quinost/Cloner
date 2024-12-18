package filemanipulator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func DeleteUnnecessary(s Settings) {
	files, _ := os.ReadDir(s.Destination)
	var wg sync.WaitGroup

	for _, file := range files {
		if strings.HasSuffix(file.Name(), s.Extension) {
			path := filepath.Join(s.Destination, file.Name())
			fmt.Printf("Usuwam %v \n", path)

			wg.Add(1)
			go func(name string) {
				defer wg.Done()
				os.Remove((path))
			}(path)
		}
	}

	wg.Wait()
}
