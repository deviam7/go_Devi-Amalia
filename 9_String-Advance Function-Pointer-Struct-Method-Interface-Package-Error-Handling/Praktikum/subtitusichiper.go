package main

import "fmt"

type Chiper interface {
	Encode() string
	Decode() string
}

type student struct {
	name       string
	nameEncode string
	score      int
}

func (s *student) Encode() string {
	var nameEncode = ""

	for _, r := range s.name {
		if r == 'z' {
			nameEncode += "a"
		} else {
			nameEncode += string(r + 1)
		}
	}

	s.nameEncode = nameEncode

	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""

	for _, r := range s.nameEncode {
		if r == 'a' {
			nameDecode += "z"
		} else {
			nameDecode += string(r - 1)
		}
	}

	return nameDecode
}

func main() {
	var menu int
	var a student

	fmt.Println("[1] Encrypt")
	fmt.Println("[2] Decrypt")
	fmt.Print("Choose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		var c Chiper = &a
		fmt.Println("\nEncode of student's name", a.name, "is:", c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Encoded Student Name: ")
		fmt.Scan(&a.nameEncode)
		var c Chiper = &a
		fmt.Println("\nDecode of student's name", a.nameEncode, "is:", c.Decode())
	}
}
