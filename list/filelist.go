package list

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

//FileList uses files to store lists
type FileList struct {
	Root string
}

//GetList gets a list
func (fl *FileList) GetList(name string) (*List, error) {
	filename := name + ".txt"
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	items := make([]string, 0)
	for scanner.Scan() {
		items = append(items, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &List{ListName: name, Items: items}, nil
}

//AddItem adds an item to the list
func (fl *FileList) AddItem(name string, item string) error {
	filename, _ := filepath.Abs(name + ".txt")
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()
	fi, _ := file.Stat()
	if fi.Size() > int64(0) {
		_, _ = file.WriteString("\n")
	}
	if _, err := file.WriteString(item); err != nil {
		return err
	}
	return nil
}
func (fl *FileList) writeList(filename string, lst *List) error {
	err := os.Truncate(filename, 0)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, i := range lst.Items {
		if _, err := file.WriteString(i + "\n"); err != nil {
			return err
		}
	}
	return nil
}
func (fl *FileList) RemoveItem(name string, item int) error {
	filename := name + ".txt"
	lst, err := fl.GetList(name)
	if err != nil {
		return err
	}
	lst.Items = append(lst.Items[:item], lst.Items[item+1:]...)
	for _, i := range lst.Items {
		fmt.Println(i)
	}
	return fl.writeList(filename, lst)
}
