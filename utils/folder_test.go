package utils

import "testing"

func TestMkdir(t *testing.T) {
	got := GenResultFilePath(10)
		
	t.Log(got)
}
