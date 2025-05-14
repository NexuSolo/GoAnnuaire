package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	commands := map[string]func(){
		"add": func() {
			dico := make(map[string]string)
			add(dico, "Alice", "123")
		},
		"addMissing": func() {
			dico := make(map[string]string)
			add(dico, "", "0")
		},
		"searchFound": func() {
			dico := map[string]string{"Bob": "456"}
			search(dico, "Bob")
		},
		"searchNotFound": func() {
			dico := map[string]string{"Bob": "456"}
			search(dico, "Alice")
		},
		"searchEmpty": func() {
			dico := make(map[string]string)
			search(dico, "")
		},
	}
	for name, cmd := range commands {
		t.Run(name, func(t *testing.T) {
			cmd()
		})
	}
}

func TestAdd(t *testing.T) {
	dico := make(map[string]string)
	add(dico, "Alice", "123")
	// Output: Added: Alice 123
}

func TestAddMissing(t *testing.T) {
	dico := make(map[string]string)
	add(dico, "", "0")
	// Output: Please provide a name and a telephone number
}

func TestSearchFound(t *testing.T) {
	dico := map[string]string{"Bob": "456"}
	search(dico, "Bob")
	// Output: Found: Bob => 456
}

func TestSearchNotFound(t *testing.T) {
	dico := map[string]string{"Bob": "456"}
	search(dico, "Alice")
	// Output: Not found: Alice
}

func TestSearchEmpty(t *testing.T) {
	dico := make(map[string]string)
	search(dico, "")
	// Output: Please provide a name
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = stdOut
	buf.ReadFrom(r)
	return buf.String()
}

func TestListFlexible(t *testing.T) {
	dico := map[string]string{"A": "1", "B": "2"}
	output := captureOutput(func() {
		list(dico)
	})
	out1 := "A => 1\nB => 2\n"
	out2 := "B => 2\nA => 1\n"
	if output != out1 && output != out2 {
		t.Errorf("Sortie inattendue : %q", output)
	}
}

func TestListEmpty(t *testing.T) {
	dico := make(map[string]string)
	list(dico)
	// Output: No entries found
}

func TestDeleteEntry(t *testing.T) {
	dico := map[string]string{"A": "1"}
	deleteEntry(dico, "A")
	// Output: Deleting: A
}

func TestDeleteEntryNotFound(t *testing.T) {
	dico := map[string]string{"A": "1"}
	deleteEntry(dico, "B")
	// Output: Not found: B
}

func TestDeleteEntryEmpty(t *testing.T) {
	dico := make(map[string]string)
	deleteEntry(dico, "")
	// Output: Please provide a name
}

func TestExistTrue(t *testing.T) {
	dico := map[string]string{"A": "1"}
	exist(dico, "A")
	// Output: It exists
}

func TestExistFalse(t *testing.T) {
	dico := map[string]string{"A": "1"}
	exist(dico, "B")
	// Output: It doesn't exist
}

func TestExistEmpty(t *testing.T) {
	dico := make(map[string]string)
	exist(dico, "")
	// Output: Please provide a name
}

func TestModify(t *testing.T) {
	dico := map[string]string{"A": "1"}
	modify(dico, "A", "999")
	// Output: Modified: A 999
}

func TestModifyNotFound(t *testing.T) {
	dico := map[string]string{"A": "1"}
	modify(dico, "B", "888")
	// Output: Not found: B
}

func TestModifyNoTel(t *testing.T) {
	dico := map[string]string{"A": "1"}
	modify(dico, "A", "0")
	// Output: Please provide a new telephone number
}
