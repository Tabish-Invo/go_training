//custom reader func

// package main

// import (
// 	"fmt"
// 	"io"
// )

// type MyStringData struct {
// 	str       string
// 	readIndex int
// }

// func (myStringData *MyStringData) Read(p []byte) (n int, err error) {

// 	strBytes := []byte(myStringData.str)
// 	if myStringData.readIndex >= len(strBytes) {
// 		return 0, io.EOF
// 	}

// 	nextReadLimit := myStringData.readIndex + len(p)

// 	if nextReadLimit >= len(strBytes) {
// 		nextReadLimit = len(strBytes)
// 		err = io.EOF
// 	}

// 	nextBytes := strBytes[myStringData.readIndex:nextReadLimit]
// 	n = len(nextBytes)

// 	copy(p, nextBytes)

// 	myStringData.readIndex = nextReadLimit

// 	return

// }

// func main() {

// 	// create data source
// 	src := MyStringData{
// 		str: "Hello Amazing World!",
// 	}

// 	// create a packet
// 	p := make([]byte, 3) // slice of length `3`

// 	// read `src` until an error is returned
// 	for {

// 		// read `p` bytes from `src`
// 		n, err := src.Read(p)
// 		fmt.Printf("%d bytes read, data: %s\n", n, p[:n])

// 		// handle error
// 		if err == io.EOF {
// 			fmt.Println("--end-of-file--")
// 			break
// 		} else if err != nil {
// 			fmt.Println("Oops! Some error occured!", err)
// 			break
// 		}
// 	}

// }

// built-in function for read

// package main

// import (
// 	"fmt"
// 	"io"
// 	"strings"
// )

// func main() {

// 	// create data source
// 	src := strings.NewReader("Hello Amazing World!")

// 	// create a packet
// 	p := make([]byte, 3) // slice of length `3`

// 	// read `src` until an error is returned
// 	for {

// 		// read `p` bytes from `src`
// 		n, err := src.Read(p)
// 		fmt.Printf("%d bytes read, data: %s\n", n, p[:n])

// 		// handle error
// 		if err == io.EOF {
// 			fmt.Println("--end-of-file--")
// 			break
// 		} else if err != nil {
// 			fmt.Println("Oops! Some error occured!", err)
// 			break
// 		}
// 	}

// }

//read_all,read_full,limit_reader

// package main

// import (
// 	"fmt"
// 	"io"
// 	"strings"
// )

// func main() {

// 	src := strings.NewReader("This is a string")

// 	//ioutil.ReadAll
// 	// data, _ := ioutil.ReadAll(src)
// 	// fmt.Printf("Length of %s is %d\n", data, len(data))

// 	//io.ReadFull
// 	buff := make([]byte, 10)
// 	str1, err1 := io.ReadFull(src, buff)
// 	fmt.Printf("Bytes read: %d, value: %s, err: %v\n", str1, buff[:str1], err1)
// 	str2, err2 := io.ReadFull(src, buff)
// 	fmt.Printf("Bytes read: %d, value: %s, err: %v\n", str2, buff[:str2], err2)
// 	str3, err3 := io.ReadFull(src, buff)
// 	fmt.Printf("Bytes read: %d, value: %s, err: %v\n", str3, buff[:str3], err3)

// 	////io.LimitReader
// 	src2 := strings.NewReader("Hello Amazing World!")
// 	newSrc := io.LimitReader(src2, 10)
// 	p := make([]byte, 3)

// 	for {

// 		n, err := newSrc.Read(p)
// 		fmt.Printf("%d bytes read, data: %s\n", n, p[:n])

// 		if err == io.EOF {
// 			fmt.Println("--end-of-file--")
// 			break
// 		} else if err != nil {
// 			fmt.Println("Oops! Some error occured!", err)
// 			break
// 		}
// 	}

// }

//write_string

// package main

// import (
// 	"fmt"
// 	"io"
// )

// // SampleStore - sample store type
// type SampleStore struct {
// 	data []byte
// }

// // implement `io.Writer` interface
// func (ss *SampleStore) Write(p []byte) (n int, err error) {

// 	// check if `10` bytes has been written
// 	if len(ss.data) == 10 {
// 		return 0, io.EOF // end of limit error
// 	}

// 	// get remaining capacity of the `ss.data`
// 	remainingCap := 10 - len(ss.data)

// 	// get length of data to write
// 	writeLength := len(p)
// 	if remainingCap <= writeLength {
// 		writeLength = remainingCap
// 		err = io.EOF
// 	}

// 	// append `writeLength` of data from `p` to `ss.data`
// 	ss.data = append(ss.data, p[:writeLength]...)

// 	// set number of bytes written and return
// 	n = writeLength
// 	return
// }

// func main() {

// 	ss := &SampleStore{}

// 	// write 1: "Hello!"
// 	bytesWritten1, err1 := io.WriteString(ss, "Hello!")
// 	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten1, err1)
// 	fmt.Printf("Value of ss.data: %s\n\n", ss.data)

// 	// write 2: " Amazing"
// 	bytesWritten2, err2 := io.WriteString(ss, " Amazing")
// 	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten2, err2)
// 	fmt.Printf("Value of ss.data: %s\n\n", ss.data)

// 	// write 3: " World!"
// 	bytesWritten3, err3 := io.WriteString(ss, " World!")
// 	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten3, err3)
// 	fmt.Printf("Value of ss.data: %s\n\n", ss.data)

// }

// Standard I/O Streams
// The standard I/O streams viz. os.Stdin, os.Stdout and os.Stderr implement the io.Writer and io.StringWriter interfaces.
// Hence we can call their Write or WriteString methods to write some data. The fmt package also provides some functions to write some data to io.Writer objects.
// func Fprint(w io.Writer, a ...interface{}) (n int, err error)
// func Fprintln(w, a) (n int, err error)
// func Fprintf(w, format string, a) (n int, err error)

// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {

// 	io.WriteString(os.Stdout, "Hello World!\n")
// 	os.Stdout.Write([]byte("Hello World!\n"))

// 	fmt.Fprint(os.Stdout, "Hello World!\n")
// 	fmt.Fprintln(os.Stdout, "Hello World!")
// 	fmt.Fprintf(os.Stdout, "%s World!\n", "Hello")

// }

// Closing I/O Operations
// type Closer interface {
//     Close() error
// }

// type ReadCloser interface {
//     Reader
//     Closer
// }
// type WriteCloser interface {
//     Writer
//     Closer
// }
// type ReadWriteCloser interface {
//     Reader
//     Writer
//     Closer
// }

// Transferring Data between streams

//io.copy

// package main

// import (
// 	"io"
// 	"os"
// 	"strings"
// )

// func main() {

// 	// create a string `io.Reader` object
// 	stringReader := strings.NewReader("Hello World! How are you?\n")

// 	// copy data from `stringsReader` to `os.Stdout` (`io.Writer`)
// 	// io.Copy(os.Stdout, stringReader)
// 	io.CopyN(os.Stdout, stringReader, 12)

// }

//io.pipe

// package main

// import (
// 	"fmt"
// 	"io"
// )

// func main() {

// 	// create a pipe
// 	src, dst := io.Pipe()

// 	// start goroutine that writes data to `dst`
// 	go func() {
// 		dst.Write([]byte("DATA_1"))
// 		dst.Write([]byte("DATA_2"))
// 		dst.Close()
// 	}()

// 	packet := make([]byte, 6)

// 	bytesRead1, err1 := src.Read(packet)
// 	fmt.Printf("bytes read: %d, value: %s, err: %v\n", bytesRead1, packet, err1)

// 	bytesRead2, err2 := src.Read(packet)
// 	fmt.Printf("bytes read: %d, value: %s, err: %v\n", bytesRead2, packet, err2)

// 	bytesRead3, err3 := src.Read(packet)
// 	fmt.Printf("bytes read: %d, value: %s, err: %v\n", bytesRead3, packet, err3)

// }

//buffered streams

// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {

// 	// create new buffer
// 	buf := bytes.NewBufferString("Hello World!")

// 	// write some data to the buffer
// 	fmt.Print("bytes written => ")
// 	fmt.Println(buf.WriteString("How are you?"))

// 	// append data from a `io.Reader` to the buffer
// 	strReader := strings.NewReader(" Doing Well? ")
// 	fmt.Print("bytes written => ")
// 	fmt.Println(buf.ReadFrom(strReader))

// 	// read first `12 bytes` from the buffer
// 	fmt.Print("bytes read => ")
// 	fmt.Println(buf.Read(make([]byte, 12)))

// 	// read all `unread bytes` to STDOUT
// 	fmt.Print("bytes read => ")
// 	fmt.Println(buf.WriteTo(os.Stdout))

// 	// read another `10 bytes` from the buffer
// 	fmt.Print("bytes read => ")
// 	fmt.Println(buf.Read(make([]byte, 10))) // EOF

// 	// write some more data to the buffer
// 	fmt.Print("bytes written => ")
// 	fmt.Println(buf.WriteString("Hello!"))

// 	// read another `10 bytes` from the buffer
// 	fmt.Print("bytes read => ")
// 	fmt.Println(buf.Read(make([]byte, 10)))
// }

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// create a buffered `io.Writer` from `Stdout`
	dst := bufio.NewWriter(os.Stdout)

	// write some data to the buffer
	fmt.Println(dst.WriteString("Hello World!"))
	fmt.Println(dst.Write([]byte(" How are you?\n")))

	// write buffered data to underlying `io.Writer`
	fmt.Println(dst.Flush(), "Flushed!")

	// write some more data to the buffer
	fmt.Println(dst.WriteString("Gooood?\n"))
	fmt.Println(dst.Flush(), "Flushed!")

}
