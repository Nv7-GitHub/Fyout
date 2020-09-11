package Widgets

import (
	"compress/gzip"
	"encoding/gob"
	"os"
)

// Read reads the file
func Read() {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	file, err := gzip.NewReader(f)
	if err != nil {
		panic(err)
	}

	dec := gob.NewDecoder(file)

	var rootS RootSerialized
	err = dec.Decode(&rootS)
	if err != nil {
		panic(err)
	}

	if rootS == *new(RootSerialized) {
		root = new(Root)
	} else {
		root = rootS.Deserialize(func() {})
	}

	file.Close()
}

// Save saves the data
func Save() {
	if root.(*Root).Child != nil {
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			panic(err)
		}

		file := gzip.NewWriter(f)

		enc := gob.NewEncoder(file)

		err = enc.Encode(root.Serialize())
		if err != nil {
			panic(err)
		}

		file.Close()
	}
}
