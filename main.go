package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"comment_tree/post"
)

func main() {
	var p post.Post
	err := json.NewDecoder(os.Stdin).Decode(&p)
	if err != nil {
		log.Fatal(err)
	}

	p.FillBodies()

	ans, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(ans))
}