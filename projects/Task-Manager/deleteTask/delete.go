package deletetask

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Delete(f string, idxStr string) error {
	idx, err := strconv.Atoi(idxStr)
	if err != nil {
		return fmt.Errorf("invalid index")
	}

	if idx < 1 {
		return fmt.Errorf("index must be >= 1")
	}

	file, err := os.Open(f)
	if err != nil {
		return err
	}

	defer file.Close()

	var tasks []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks = append(tasks, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if idx > len(tasks) {
		return fmt.Errorf("task %d does not exist", idx)
	}

	tasks = append(tasks[:idx-1], tasks[idx:]...)

	var output strings.Builder
	for _, task := range tasks {
		output.WriteString(task + "\n")
	}

	return os.WriteFile(f, []byte(output.String()), 0644)
}
