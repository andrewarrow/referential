package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/andrewarrow/cloutcli"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("")
		fmt.Println("This is a referential implementation of a simple app")
		fmt.Println("that uses cloutcli.")
		fmt.Println("")
		fmt.Println("Case 1. bitclout.com api with no username")
		fmt.Println("Case 2. bitclout.com api with username, READ-ONLY access")
		fmt.Println("Case 3. query a copy of the complete badgerdb database (>70 GB)")
		fmt.Println("Case 4. import badgerdb database to sqlite, then query sqlite")
		fmt.Println("Case 5. bitclout.com api with username, WRITE access")
		fmt.Println("")
		fmt.Println("./referential --case1")
		fmt.Println("./referential --case2")
		fmt.Println("./referential --case3")
		fmt.Println("")
		return
	}

	flavor := os.Args[1]

	if flavor == "--case1" {
		// case 1, anonymous, hitting bitclout.com api for small amount of data
		fmt.Println("list messages from the global feed...")
		list := cloutcli.GlobalPosts()
		fmt.Printf("    Got back %d posts.\n", len(list))
		fmt.Printf("    First one:\n")
		PrintPost(list[0].Body)
	} else if flavor == "--case2" {
		// case 2, username with read-only access
		// same as above but to show following feed vs global need a "who" is viewing
		fmt.Println("\nlist messages from a user's following feed...")
		list := cloutcli.FollowingFeedPosts("andrewarrow")
		fmt.Printf("    Got back %d posts.\n", len(list))
		fmt.Printf("    First one:\n")
		PrintPost(list[0].Body)
	} else if flavor == "--case3" {
		// case 3, query a copy of the complete badgerdb database (>70 GB)
		// you'll have to spend 24 hours first downloading it or
		// get a copy from someone who already did that
		cloutcli.PrintAllPostsFromBadger("../acopy/badgerdb")
	}

	// case 4, import parts of that complete badgerdb database into a local
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
