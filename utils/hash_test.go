package utils

import (
	"os"
	"testing"
)

func TestHash(t *testing.T) {
	file, err := os.Open("hash.go")
	if err != nil {
		t.Fatal(err)
	}
	got, err := GetMd5FromFile(file)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(got)
}
