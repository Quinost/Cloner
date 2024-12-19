package filemanipulator

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func Copy(s *Settings) {
	filesInDirectory, _ := os.ReadDir(s.From)

	var filesToCopy []os.DirEntry
	for _, file := range filesInDirectory {
		if strings.HasSuffix(file.Name(), s.Extension) {
			filesToCopy = append(filesToCopy, file)
		}
	}

	var totalSize int64
	for _, file := range filesToCopy {
		info, _ := file.Info()
		totalSize += info.Size()
	}

	progressChan := make(chan int64, len(filesToCopy))
	var currentProgress int64

	go func() {
		for progress := range progressChan {
			currentProgress += progress
			percentage := float64(currentProgress) / float64(totalSize) * 100
			fmt.Printf("\r[%.2f%%] Kopiowanie: %d/%d MB",
				percentage,
				currentProgress/1024/1024,
				totalSize/1024/1024)
		}
	}()

	var wg sync.WaitGroup

	for _, file := range filesToCopy {
		fromFilePath := filepath.Join(s.From, file.Name())
		destinationFilePath := filepath.Join(s.Destination, file.Name())

		wg.Add(1)
		go func(fromPath, destPath string) {
			defer wg.Done()

			src, err := os.Open(fromPath)
			if err != nil {
				fmt.Printf("Błąd otwierania pliku %s: %v\n", fromPath, err)
				return
			}
			defer src.Close()

			dst, err := os.Create(destPath)
			if err != nil {
				fmt.Printf("Błąd tworzenia pliku %s: %v\n", destPath, err)
				return
			}
			defer dst.Close()

			buf := make([]byte, 32*1024)
			var copiedSize int64
			for {
				srcBytesNr, readErr := src.Read(buf)

				if srcBytesNr > 0 {
					dstBytesNr, writeErr := dst.Write(buf[:srcBytesNr])
					if dstBytesNr > 0 {
						copiedSize += int64(dstBytesNr)
						progressChan <- int64(dstBytesNr)
					}
					if writeErr != nil {
						err = writeErr
						break
					}
				}
				if readErr == io.EOF {
					break
				}
				if readErr != nil {
					err = readErr
					break
				}
			}

			if err != nil {
				fmt.Printf("Błąd kopiowania %s: %v\n", filepath.Base(fromPath), err)
			}
		}(fromFilePath, destinationFilePath)
	}

	wg.Wait()
	close(progressChan)
	fmt.Println("")
}
