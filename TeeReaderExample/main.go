package main

// This program downloads a gzip file from a URL and tracks the download progress
// Based on YT video https://youtu.be/fJHNhL1FUEs?si=EvqKAMRGf1t5dZ9i

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// Counter tracks the total number of bytes processed
type Counter struct {
	Total uint64
}

// Write implements io.Writer interface to count bytes processed
// Reports download progress in MB every 32KB chunk
func (c *Counter) Write(p []byte) (int, error) {
	c.Total += uint64(len(p))
	progress := float64(c.Total) / 1024 / 1024
	fmt.Println("Downloaded", progress, "MB")
	return len(p), nil
}

func main() {
	// Fetch the remote file via HTTP GET request
	res, err := http.Get("http://storage.googleapis.com/books/ngrams/books/20200217/eng/5-00020-of-19423.gz")
	if err != nil {
		log.Fatalf("Failed to fetch file: %v", err)
	}
	defer res.Body.Close()

	// Check if the request was successful
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Bad status: %s", res.Status)
	}

	// Verify content type (optional but recommended)
	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/x-gzip") && !strings.Contains(contentType, "application/octet-stream") {
		log.Printf("Warning: Unexpected content type: %s", contentType)
	}

	// Create or overwrite local file with 0600 permissions
	local, err := os.OpenFile("downloaded.txt", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("Failed to create local file: %v", err)
	}
	defer local.Close()

	// Create buffered reader for better performance
	bufferedReader := bufio.NewReader(res.Body)

	// Create the gzip reader
	dec, err := gzip.NewReader(bufferedReader)
	if err != nil {
		log.Fatalf("Failed to create gzip reader: %v", err)
	}
	defer dec.Close()

	// Create counter for progress tracking
	counter := &Counter{}

	// Stream and decompress file contents while tracking progress
	if _, err := io.Copy(local, io.TeeReader(dec, counter)); err != nil {
		log.Fatalf("Failed to copy data: %v", err)
	}
}