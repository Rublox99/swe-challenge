package main

import (
	"enron_indexer/pkg/helpers"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

const mailsDir = "enron_mail_20110402/maildir"

func main() {
	/* HTTP server for profiling */
	go func() {
		fmt.Println("Starting pprof server on :6060")
		http.ListenAndServe(":6060", nil)
	}()

	err := filepath.Walk(mailsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			helpers.ProcessFile(path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error while iterating the files: %v\n", err)
	}
}
