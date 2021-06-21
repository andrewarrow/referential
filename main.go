package main

import (
	"fmt"

	"github.com/andrewarrow/cloutcli"
)

func main() {
	// case 1, anonymous, hitting bitclout.com api for small amount of data
	fmt.Println("list messages from the global feed...")

	list := cloutcli.GlobalPosts()

	for _, post := range list {
		fmt.Println(post.Body)
	}

	// case 2, username with read-only access
	// same as above but to show following feed vs global need a "who" is viewing

	// case 3, username with write access
	// to post a message or send coin, cli needs private key i.e seed words

	// case 4, query a copy of the complete badgerdb database (>70 GB)
	// you'll have to spend 24 hours first downloading it or
	// get a copy from someone who already did that

	// case 5, import parts of that complete basgerdb database into a local
	// sqlite database first, then query that sqlite database

}
