package main

import "fmt"

func main() {
	//starts with new line ends with new line
	fmt.Print("hello world", 1, 2)

	//variadic function similar to print
	fmt.Println("hello world")

	//same as println but it returns
	x := fmt.Sprintln("hello sprint")
	fmt.Println(x)

	//string interpolation
	// fmt.Printf() //used for interpolation

	// %v for the value
	// fmt.Sprintf() is similar to fmt.Printf() but it returns to a standard output

	// struct
	s := struct {
		name string
		age  int
	}{"John", 26}

	// channel
	c := make(chan int)

	// map
	m := map[string]int{
		"one": 1,
		"two": 2,
	}

	fmt.Printf("%v \n", true)           // boolean values //true
	fmt.Printf("%v \n", 132)            // signed and unsigned integers //132
	fmt.Printf("%v \n", 198.454)        // floating point numbers //198.454
	fmt.Printf("%v \n", 3+7i)           // complex numbers //(3+7i)
	fmt.Printf("%v \n", "Hello World!") // strings //Hello World!
	fmt.Printf("%v \n", []int{1, 2, 3}) // slices and arrays //[1 2 3]
	fmt.Printf("%v \n", m)              // maps //map[one:1 two:2]
	fmt.Printf("%v \n", s)              // structs //{John 26}
	fmt.Printf("%v \n", c)              // channels //0xc000062060
	fmt.Printf("%v \n", &s)             // pointers //&{John 26}

	fmt.Printf("%v \n", []int{1, 2, 3}) // print as in the code '#' flag
	fmt.Printf("%+v \n", s)             // display the fields of struct '+' flag

	fmt.Printf("[1] %f\n", 1225.04532684676)     // defalt width and preicision
	fmt.Printf("[2] %.11f\n", -1225.04532684676) // defalt width
	fmt.Printf("[3] %20.12f\n", 1225.04532684676)
	fmt.Printf("[4] %20f\n", 1225.04532684676)  // defalt precision
	fmt.Printf("[5] %20.f\n", 1225.04532684676) // zero precision
	fmt.Printf("[6] %.f\n", 1225.04532684676)   // defalt width and zero precision
	fmt.Printf("[7] %#.f\n", 1225.04532684676)  // always show decimal point

	fmt.Printf("[8] %+.2F\n", 1225.04532684676) // alias of %f (with sign)

	//can use %#v to format value
	//can use %+v to get fields name from struct
	//%T is used to get type of the variable
	//%t to convert boolean value to string
	// %d converts hexa/octal decimal to base10 format also used to display rune value. "+"," " is used for signed value or spcar %+d,% d
	//%b to convert into binary format
	//%x hexadecimal format
	//%o octal format
	//%f decimal format
	//%e scientific notation //can be used when a floating point number is too large to display
	//%g scientific notation with significant fields
	//%s used for string format
	// %q used for escaped string format
	// %p pointer address format

}
