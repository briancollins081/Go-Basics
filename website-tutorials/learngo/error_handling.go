package main

import (
	"fmt"
	"path/filepath"
)

// 2. Asserting the underlying struct type and getting more information from the struct fields
type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e PathError) Error() string { // now PathError implements error interface
	return e.Op + " " + e.Path + " " + e.Err.Error()
}

func main() {
	// 1. Example
	/* f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "was opened successfully!") */

	// 2. Asserting the underlying struct type and getting more information from the struct fields
	/* f, err := os.Open("test.txt")
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Failed to open file at path", pErr.Path)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(f.Name(), "opened successfully") */

	// 3. Asserting the underlying struct type and getting more information using methods
	/* addr, err := net.LookupHost("golangbot123.com")
	if err != nil {
		if dnsErr, ok := err.(*net.DNSError); ok { //type assertion
			if dnsErr.Timeout() {
				fmt.Println("operation timed out")
				return
			}
			if dnsErr.Temporary() {
				fmt.Println("temporary error")
				return
			}
			fmt.Println("Generic DNS Error", err)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(addr) */

	// 4. Direct comparison
	files, err := filepath.Glob("[")
	if err != nil {
		if err == filepath.ErrBadPattern {
			fmt.Println("Bad pattern error:", err)
			return
		}
		fmt.Println("Generic pattern error")
		return
	}
	fmt.Println("matched files", files)
}
