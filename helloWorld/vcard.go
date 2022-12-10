package main

import (
	"fmt"
	"time"
)

type Address struct {
	Street           string
	HouseNumber      uint32
	HouseNumberAddOn string
	POBox            string
	ZipCode          string
	City             string
	Country          string
}

type VCard struct {
	name string
	address map[string]*Address
	Birth time.Time
	image string
}

func main() {
	addr1 := &Address{"一德路", 196, "", "", "510000", "Guangzhou", "China"}
	addr2 := &Address{"Heideland", 28, "", "", "2640", "Mortsel", "België"}

	addrs := make(map[string]*Address)

	addrs["youth"] = addr1
	addrs["now"] = addr2
	birthday := time.Date(1998,11,10,0,0,0,0,time.Local)
	photo := "MyDocuments/MyPhotos/photo1.jpg"
	vcard := &VCard{"Mars",addrs, birthday, photo }


	fmt.Printf("Here is the full VCard: %v\n", vcard)
	fmt.Printf("My Addresses are:\n %v\n %v", addr1, addr2)
}