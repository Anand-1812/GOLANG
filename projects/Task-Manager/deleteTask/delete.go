package deletetask

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
)

func Delete(f string, idxStr string) error {
	idx, err := strconv.Atoi(idxStr)
	if err != nil {
		return fmt.Errorf("Invalid index")
	}

	if idx < 1 {
		return fmt.Errorf("index must be >= 1")
	}

	file, err := os.Open(f)
	if err != nil {
		return err
	}

	defer file.Close()


}
