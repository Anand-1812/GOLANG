package addtask

import "os"

func Add(f *os.File, task string) error {
	_, err := f.WriteString(task + "\n")

	return err
}
