package main

import (
	"fmt"
	"strings"

	"github.com/andrewarrow/cloutcli"
)

func main() {
	// case 1, anonymous, hitting bitclout.com api for small amount of data
	fmt.Println("list messages from the global feed...")

	list := cloutcli.GlobalPosts()
	fmt.Printf("    Got back %d posts.\n", len(list))
	fmt.Printf("    First one:\n")
	PrintPost(list[0].Body)

	// case 2, username with read-only access
	// same as above but to show following feed vs global need a "who" is viewing

	fmt.Println("\nlist messages from a user's following feed...")
	list = cloutcli.FollowingFeedPosts("andrewarrow")
	fmt.Printf("    Got back %d posts.\n", len(list))
	fmt.Printf("    First one:\n")
	PrintPost(list[0].Body)

	// case 3, query a copy of the complete badgerdb database (>70 GB)
	// you'll have to spend 24 hours first downloading it or
	// get a copy from someone who already did that

	// case 4, import parts of that complete basgerdb database into a local
	// sqlite database first, then query that sqlite database

	// case 5, username with write access
	// to post a message or send coin, cli needs private key i.e seed words

}

func PrintPost(body string) {
	tokens := strings.Split(body, "\n")
	for _, line := range tokens {
		fmt.Println("   ", line)
	}
}
