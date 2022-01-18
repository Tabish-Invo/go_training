package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// simple struct
type Student struct {
	FirstName, lastName string
	Email               string
	Age                 int
	Height              float64
	IsMale              bool
}

//map to json
type Student_ map[string]interface{}

//Abstract Data Types

type Profile struct {
	Username  string
	followers int
	Grade     map[string]string
}

type Stud struct {
	FirstName, lastName string
	Age                 int
	Profile             //Profile
	Languages           []string
}

//pointer

type ProfileI interface {
	Follow()
}

type Profile_ struct {
	Username  string
	Followers int
}

func (p *Profile_) Follow() {
	p.Followers++
}

type Stud_ struct {
	FirstName, lastName string
	Age                 int
	Primary             ProfileI
	Secondary           ProfileI
}

//Data type conversion

type Age int

func (p Profile_) MarshalJSON() ([]byte, error) {
	//return JSON value
	return []byte(fmt.Sprintf(`{"f_count" : "%d"}`, p.Followers)), nil
}

func (a Age) MarshalText() ([]byte, error) {
	// return string value
	return []byte(fmt.Sprintf(`{"age" : "%d"}`, int(a))), nil
}

type Student_1 struct {
	FirstName, lastName string
	Age                 Age
	Profile             Profile_
}

// structure tags
type Profile_tag struct {
	Username  string `json:"uname"`
	Followers int    `json:"followers,omitempty,string"`
}

type Student_tag struct {
	FirstName string      `json:"fname"`           // `fname` as field name
	LastName  string      `json:"lname,omitempty"` // discard if value is empty
	Email     string      `json:"-"`               // always discard
	Age       int         `json:"-,"`              // `-` as field name
	IsMale    bool        `json:",string"`         // keep original field name, coerce to a string
	Profile   Profile_tag `json:""`                // no effect
}

type Student_type map[string]interface{}

//Encoding

type Person_buf struct {
	Name string
	Age  int
}

func main() {

	jhon := Student{
		FirstName: "Jhon",
		lastName:  "Doe",
		// Email:     "jhon@doe.com",
		Age:    24,
		Height: 5.8,
		IsMale: true,
	}

	//convert to json
	jhonJSON, _ := json.Marshal(jhon)
	fmt.Printf(string(jhonJSON))

	//map to json
	jhon2 := Student_{
		"FirstName": "Jhon",
		"lastName":  "Doe",
		// Email:     "jhon@doe.com",
		"Age":    24,
		"Height": 5.8,
		"IsMale": true,
	}

	jhon2JSON, _ := json.Marshal(jhon2)
	fmt.Printf(string(jhon2JSON))

	//Abstract Datatypes

	var jhon3 Stud
	jhon3 = Stud{
		FirstName: "Jhon",
		lastName:  "Doe",
		Age:       24,
		Profile: Profile{
			Username:  "jhondoe",
			followers: 1234,
			Grade:     map[string]string{"Math": "A", "Science": "A+"},
		},
		Languages: []string{"English", "Urdu"},
	}

	jhon3JSON, err := json.MarshalIndent(jhon3, "", " ")

	fmt.Println(string(jhon3JSON), err)

	john4 := &Stud_{
		FirstName: "John",
		lastName:  "Doe",
		Age:       21,
		Primary: &Profile_{
			Username:  "johndoe91",
			Followers: 1975,
		},
	}

	john5 := john4

	// follow `john`
	john5.Primary.Follow()

	john4JSON, _ := json.MarshalIndent(john4, "", "  ")

	fmt.Println(string(john4JSON))

	john5JSON, _ := json.MarshalIndent(john5, "", "  ")

	fmt.Println(string(john5JSON))

	john6 := &Student_1{
		FirstName: "John",
		lastName:  "Doe",
		Age:       21,
		Profile: Profile_{
			Username:  "johndoe91",
			Followers: 1975,
		}}

	john6JSON, _ := json.MarshalIndent(john6, "", "  ")

	fmt.Println(string(john6JSON))

	// define `john` struct (pointer)
	john7 := &Student_tag{
		FirstName: "John",
		LastName:  "", // empty
		Age:       21,
		Email:     "john@doe.com",
		Profile: Profile_tag{
			Username:  "johndoe91",
			Followers: 1975,
		},
	}

	// encode `john` as JSON
	john7JSON, _ := json.MarshalIndent(john7, "", "  ")

	// print JSON string
	fmt.Println(string(john7JSON))

	// decdding json
	data := []byte(`
		{
			"FirstName": "John",
			"Age": 21,
			"Username": "johndoe91",
			"Grades": null,
			"Languages": [
			  "English",
			  "French"
			]
		}`)

	// check if `data` is valid JSON
	isValid := json.Valid(data)
	fmt.Println(isValid)

	//Unmarshal

	data_ := []byte(`
	{
		"FirstName": "John",
		"lastName": "Doe",
		"Age": 21,
		"HeightInMeters": 175,
		"Username": "johndoe91"
	}`)

	var john Student

	fmt.Printf("Error: %v\n", json.Unmarshal(data_, &john))

	fmt.Printf("%#v\n", john)

	data_2 := []byte(`
	{
		"id": 123,
		"fname": "John",
		"height": 1.75,
		"male": true,
		"languages": null,
		"subjects": [ "Math", "Science" ],
		"profile": {
			"uname": "johndoe91",
			"f_count": 1975
		}
	}`)

	var john8 Student_type
	fmt.Printf("Error: %v\n", json.Unmarshal(data_2, &john))
	fmt.Printf("%#v\n\n", john8)
	i := 1
	for k, v := range john8 {
		fmt.Printf("%d: key (`%T`)`%v`, value (`%T`)`%#v`\n", i, k, k, v, v)
		i++
	}

	//Encoding

	// create a buffer to hold JSON data
	buf := new(bytes.Buffer)
	// create JSON encoder for `buf`
	bufEncoder := json.NewEncoder(buf)

	// encode JSON from `Person` structs
	bufEncoder.Encode(Person_buf{"Ross Geller", 28})
	bufEncoder.Encode(Person_buf{"Monica Geller", 27})
	bufEncoder.Encode(Person_buf{"Jack Geller", 56})

	// print contents of the `buf`
	fmt.Println(buf) // calls `buf.String()` method

	jsonStream := strings.NewReader(`
	{"Name":"Ross Geller","Age":28}
	{"Name":"Monica Geller","Age":27}
	{"Name":"Jack Geller","Age":56}
	`)

	// create JSON decoder using `jsonStream`
	decoder := json.NewDecoder(jsonStream)

	// create `Person` structs to hold decoded data
	var ross, monica Person_buf

	// decode JSON from `decoder` one line at a time
	decoder.Decode(&ross)
	decoder.Decode(&monica)

	// see value of the `ross` and `monica`
	fmt.Printf("ross: %#v\n", ross)
	fmt.Printf("monica: %#v\n", monica)

}
