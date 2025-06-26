package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var p = fmt.Println

func main() {

	var myNewReader = bufio.NewReader(os.Stdin)

	var sum int

	for {
		var input, inputErr = myNewReader.ReadString(' ')
		var trimmedInput = strings.TrimSpace(input)

		if trimmedInput == "" {
			continue
		}

		var num, convertErr = strconv.Atoi(trimmedInput)

		if convertErr == nil {
			sum += num
		} else {
			p(convertErr)
		}

		if inputErr == io.EOF {
			break
		}

		if inputErr != nil {
			p("error reading stdin", inputErr)
		}
	}

	p(sum)
}
