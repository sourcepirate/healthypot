package honeypot

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/sourcepirate/healthypot/medium"
)

func PopulatePot(tag string) {
	filename := "data/" + tag + ".json"
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot crate/open a file")
		panic(err)
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	driver := medium.New(tag)
	archive := driver.GetArchive()
	js, err := json.Marshal(archive)
	if err != nil {
		log.Fatal("Cannot write to the file")
		panic(err)
	}
	_, err = writer.WriteString(string(js))

	if err != nil {
		log.Fatal("Cannot write to the file")
		panic(err)
	}

	writer.Flush()
}

func Download() {
	for _, tag := range TAGS {
		PopulatePot(tag)
	}
}
