package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	doorID := "ffykfhsq"
	//doorID := "abc"

	// fmt.Println(puzzle1(doorID))
	fmt.Println(puzzle2(doorID))
}

func puzzle1(input string) string {
	password := strings.Builder{}

	for i := 0; ; i++ {
		value := fmt.Sprintf("%s%d", input, i)
		hash := md5.Sum([]byte(value))
		hexa := fmt.Sprintf("%x", hash)
		if strings.HasPrefix(string(hexa), "00000") {
			password.WriteByte(hexa[5])
		}

		if password.Len() == 8 {
			return password.String()
		}
	}
}

func puzzle2(input string) string {
	password := [8]byte{}
	found := 0
	for i := 0; ; i++ {
		value := fmt.Sprintf("%s%d", input, i)
		hash := md5.Sum([]byte(value))
		hexa := fmt.Sprintf("%x", hash)

		if strings.HasPrefix(string(hexa), "00000") {
			pos, err := strconv.Atoi(string(hexa[5]))

			if err == nil && pos < len(password) && password[pos] == 0 {
				password[pos] = hexa[6]
				found++
			}
		}

		if found == 8 {
			return string(password[:])
		}
	}
}
