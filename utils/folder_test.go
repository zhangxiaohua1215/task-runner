package utils

import "testing"

func TestMkdir(t *testing.T) {
	got := GenPath("123.txt")
		
	t.Log(got)
}
