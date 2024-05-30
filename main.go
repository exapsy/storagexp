package main

import (
	"fmt"
	"os"
	"time"
)

type Repo interface {
	Write(key string, value string) error
	Read(key string) (string, error)
	Delete(key string) error
}

type MemoryRepo struct {
	data map[string]string
}

func (m *MemoryRepo) Write(key string, value string) error {
	m.data[key] = value
	return nil
}

func (m *MemoryRepo) Read(key string) (string, error) {
	return m.data[key], nil
}

func (m *MemoryRepo) Delete(key string) error {
	delete(m.data, key)
	return nil
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		data: make(map[string]string),
	}
}

type FileRepo struct {
}

func (f *FileRepo) Write(key string, value string) error {
	file, err := os.Create(key)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(value)
	if err != nil {
		return err
	}

	return nil
}

func (f *FileRepo) Read(key string) (string, error) {
	file, err := os.Open(key)
	if err != nil {
		return "", err
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return "", err
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func (f *FileRepo) Delete(key string) error {
	err := os.Remove(key)
	if err != nil {
		return err
	}

	return nil
}

func NewFileRepo() *FileRepo {
	return &FileRepo{}
}

func main() {
	fmt.Println("Memory Repo")
	testRepo(NewMemoryRepo())

	fmt.Println("File Repo")
	testRepo(NewFileRepo())
}

var val string = "value"

func testRepo(repo Repo) {
	var totalReadTime time.Duration
	var totalWriteTime time.Duration

	for i := 0; i < 1000000; i++ {
		now := time.Now()
		repo.Write("key", val)
		totalWriteTime += time.Since(now)

		now = time.Now()
		repo.Read("key")
		totalReadTime += time.Since(now)

		repo.Delete("key")
	}

	fmt.Printf("\tWrite Time taken: %v\n", totalReadTime)
	fmt.Printf("\tRead Time taken: %v\n", totalWriteTime)
	fmt.Println()
	fmt.Printf("\tTotal Time: %v\n", totalReadTime+totalWriteTime)
	fmt.Printf("\tAverage Time: %v\n", (totalReadTime+totalWriteTime)/2)
	fmt.Println("--------------------------------")
}
