package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// 1. Example 1
	/* data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of the file:", string(data)) */

	// 1. Using absolute file path
	/* data, err := ioutil.ReadFile("/Users/Brian.collins/Workspace/Lessons/GoLang/website-tutorials/learngo/filehandling/test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of the file:", string(data)) */

	// 2. Passing the file path as a command line flag
	/* fptr := flag.String("fpath", "text.txt", "file path to read from")
	flag.Parse() // called before any flag is accessed by the program.
	// fmt.Println("value of fpath is", *fptr)
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data)) */

	// 3. Bundling the text file along with the binary
	/* box := packr.New("fileBox", "../filehandling")
	data, err := box.FindString("test.txt")
	if err != nil {
		fmt.Println("File reading error!", err)
		return
	}
	fmt.Println("Contents of file:", data)
	// run the following commands
	// packr2
	// go install
	*/

	// 4. Reading a file in small chunks
	/* fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
		n, err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b[0:n]))
	} */

	//  5. Reading a file line by line
	// 			1. Open the file
	// 			2. Create a new scanner from the file
	// 			3. Scan the file and read it line by line.
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
