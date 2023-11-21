package infra

import (
	"bufio"
	"fmt"
	"os"
)

func ReadUsernamesFromFile(filename string) []string {
	readFile, err := os.Open(filename)
	usernames := make([]string, 0)
	defer readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		usernames = append(usernames, fileScanner.Text())
	}
	return usernames
}
