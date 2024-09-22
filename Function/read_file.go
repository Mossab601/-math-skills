package function

import (
	"os"
	"strconv"
	"strings"
)


type Numbers struct {
	Nums []int
}

func (n *Numbers) ReadFromFile(filename string) error {

	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	content := string(data)

	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		n.Nums = append(n.Nums, num)
	}

	return nil
}