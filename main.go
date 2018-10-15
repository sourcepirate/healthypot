package main

import (
	"fmt"

	"github.com/sourcepirate/healthypot/medium"
)

func main() {
	tag := medium.New("health")
	archives := tag.GetArchive()
	fmt.Println(archives)
}
