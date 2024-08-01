package utils

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestWriteToFile(t *testing.T) {
	filename := "testfile.txt"
	content := "Hello, World!"

	err := WriteToFile(filename, content)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	defer os.Remove(filename)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("expected no error reading the file, got %v", err)
	}

	if string(data) != content {
		t.Errorf("expected file content %s, got %s", content, string(data))
	}
}
