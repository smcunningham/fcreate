package main

import (
	"fmt"
	"os"
)

// Append to a file
func main() {
	// Opens an existing file in append and write only mode
	f, err := os.OpenFile("/Users/stevan.cunningham/fs/bytes/lines.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	newLine := "File handling made easy."
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")
}

// // Write strings line by line to a file
// func main() {
// 	// Creates a file if it doesn't already exist
// 	f, err := os.Create("/Users/stevan.cunningham/fs/bytes/lines.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		f.Close()
// 		return
// 	}

// 	// Creates a slice of strings and adds them to the file line by line
// 	d := []string{"Welcome to the world of Go.", "Go is a compiled language.", "It is easy to learn Go."}
// 	for _, v := range d {
// 		fmt.Fprintln(f, v)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 	}

// 	err = f.Close()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("file written successfully")
// }

// // Write bytes to a file
// func main() {
// 	f, err := os.Create("/User/stevan.cunningham/fs/bytes/bytes.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
// 	n2, err := f.Write(d2)
// 	if err != nil {
// 		fmt.Println(err)
// 		f.Close()
// 		return
// 	}
// 	fmt.Println(n2, "bytes written successfully")
// 	err = f.Close()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }

// // Write string to a file
// func main() {
// 	f, err := os.Create("test.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	l, err := f.WriteString("Hello World")
// 	if err != nil {
// 		fmt.Println(err)
// 		f.Close()
// 		return
// 	}
// 	fmt.Println(l, "bytes written successfully")
// 	err = f.Close()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }
