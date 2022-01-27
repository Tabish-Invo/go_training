// const (
// 	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
// 	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
// 	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
// 	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
// 	// The remaining values may be or'ed in to control behavior.
// 	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
// 	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
// 	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
// 	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
// 	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
// )

package main

import (
	"log"
	"os"
)

func main() {

	//write to file

	file, err := os.OpenFile("info.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0744)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(file, "Hello World!\n")
	fmt.Fprint(file, "How Are You\n")

	//methods to write strings to file

	// n, err := file.Write([]byte("Hello World!"))
	// n, err := file.WriteString("Hello World!")
	// n, err := io.WriteString(file, "Hello World!")

	//read from file

	// file, err := os.OpenFile("info.txt", os.O_RDONLY, 0744)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// data, _ := ioutil.ReadAll(file)
	// fmt.Printf("Data : %s", data)

	//change ownership and mod of file

	// file, _ := os.OpenFile("info.txt", os.O_WRONLY, 0744)

	// adminGroup, _ := user.LookupGroup("admin")
	// adminGroupId, _ := strconv.ParseInt(adminGroup.Gid, 10, 64)

	// fmt.Println("Admin Group ID ==> ", adminGroupId)

	// userGroup, _ := user.Lookup("tabishamin")
	// userGroupId, _ := strconv.ParseInt(userGroup.Gid, 10, 64)

	// fmt.Println("User Group ID ==> ", userGroupId)

	// file.Chown(int(userGroupId), int(adminGroupId))

	// file.Chmod(0755)

	//closing a file

	// file.Close()

	// data, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	fmt.Println("Error while reading contents of file is : ", err)
	// } else {
	// 	fmt.Printf("Data : %s", data)
	// }

	// get file stat of `info.txt` file in the CWD
	// fileinfo, err := os.Stat("info.txt")

	// // log error
	// if err != nil {
	// 	log.Fatal("Error:", err)
	// }

	// // print file information
	// fmt.Println("file name:", fileinfo.Name())
	// fmt.Println("file size (bytes):", fileinfo.Size())
	// fmt.Println("file mod. time:", fileinfo.ModTime().Format("2 Jan 2006"))
	// fmt.Println("file a directory:", fileinfo.IsDir())

	// get file stat of `info.txt` file in the CWD
	// fileinfo, err := os.Stat("info.txt")

	// // log error
	// if err != nil {
	// 	log.Fatal("Error:", err)
	// }

	// // get `os.FileMode`object from the `fileInfo`
	// fileMode := fileinfo.Mode()

	// // print file-mode information
	// fmt.Println("file a directory:", fileMode.IsDir())
	// fmt.Println("file is regular:", fileMode.IsRegular())
	// fmt.Printf("file permission: %b(binary), %o(octal)\n", fileMode.Perm(), fileMode.Perm())
	// fmt.Println("file-mode:", fileMode.String())

	// rename `info.txt` to `readme.txt` (in CWD)
	// if err := os.Rename("info.txt", "readme.txt"); err != nil {
	// 	log.Fatal("Error: ", err)
	// }

	// // move `error.log` to `logs` directory (in CWD)
	// if err := os.Rename("error.log", "logs/error.log"); err != nil {
	// 	log.Fatal("Error: ", err)

	// }

	// func Remove(path string) error
	// // remove a nested sub directory (in CWD)
	// if err := os.RemoveAll("sub/nested"); err != nil {
	// 	log.Fatal("Error: ", err)
	}

}
