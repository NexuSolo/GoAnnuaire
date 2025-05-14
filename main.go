package main

import (
	"flag"
	"fmt"
)

func main() {
	var action string
	var name string
	var tel string

	flag.StringVar(&action, "action", "default", "Action to perform")
	flag.StringVar(&name, "name", "", "Name of the person")
	flag.StringVar(&tel, "tel", "0", "Telephone number")

	flag.Parse()

	dico := make(map[string]string)
	dico["one"] = "1"
	dico["two"] = "2"
	dico["three"] = "3"

	switch action {
	case "add":
		add(dico, name, tel)
	case "search":
		search(dico, name)
	case "list":
		list(dico)
	case "delete":
		deleteEntry(dico, name)
	case "exist":
		exist(dico, name)
	case "modify":
		modify(dico, name, tel)
	}

}

func add(dico map[string]string, name string, tel string) {
	if name == "" || tel == "0" {
		fmt.Println("Please provide a name and a telephone number")
		return
	}
	dico[name] = tel
	fmt.Println("Added:", name, tel)
}

func search(dico map[string]string, name string) {
	if name == "" {
		fmt.Println("Please provide a name")
		return
	}
	if val, ok := dico[name]; ok {
		fmt.Printf("Found: %s => %s\n", name, val)
	} else {
		fmt.Println("Not found:", name)
	}
}

func list(dico map[string]string) {
	for k, v := range dico {
		fmt.Printf("%s => %s\n", k, v)
	}
	if len(dico) == 0 {
		fmt.Println("No entries found")
	}
}

func deleteEntry(dico map[string]string, name string) {
	if name == "" {
		fmt.Println("Please provide a name")
		return
	}
	if _, ok := dico[name]; ok {
		fmt.Println("Deleting:", name)
		delete(dico, name)
	} else {
		fmt.Println("Not found:", name)
	}
}

func exist(dico map[string]string, name string) {
	if name == "" {
		fmt.Println("Please provide a name")
		return
	}
	if _, ok := dico[name]; ok {
		fmt.Println("It exists")
	} else {
		fmt.Println("It doesn't exist")
	}

}

func modify(dico map[string]string, name string, tel string) {
	if tel == "0" {
		fmt.Println("Please provide a new telephone number")
		return
	}
	if _, ok := dico[name]; ok {
		dico[name] = tel
		fmt.Println("Modified:", name, tel)
	} else {
		fmt.Println("Not found:", name)
	}

}
