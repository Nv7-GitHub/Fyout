package Widgets

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"os"
)

// Read reads the file
func Read() {
	file, _ := os.Open(path)
	dec := gob.NewDecoder(file)

	var rootS RootSerialized
	dec.Decode(&rootS)

	if rootS == *new(RootSerialized) {
		root = new(Root)
	} else {
		root = rootS.Deserialize(func() {})
	}
	file.Close()
}

// Save saves the data
func Save() {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	enc.Encode(root.Serialize())

	ioutil.WriteFile(path, buff.Bytes(), 0644)
}
