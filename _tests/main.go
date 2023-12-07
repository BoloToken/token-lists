package tests

import (
	"fmt"
	"os"
	"path"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		os.Exit(1)
	}

	dir := path.Join(wd, "tokenslist")
	fmt.Println(dir)
}
