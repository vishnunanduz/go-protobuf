package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	person := &Person{
		Firstname: "Vishnujith",
		Lastname:  "Nandakumar",
	}

	serializedPerson, err := proto.Marshal(person)

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("person.data", serializedPerson, 0644)
	if err != nil {
		log.Fatal(err)
	}

	persons := &Person{}
	err = proto.Unmarshal(serializedPerson, persons)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(persons.GetFirstname())
	fmt.Println(persons.GetLastname())

}
