package advent

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetLines() chan string {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	ch := make(chan string)
	go func() {
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)
		err = scanner.Err()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()
	return ch
}

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return i
}
