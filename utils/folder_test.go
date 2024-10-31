package utils

import "testing"

func TestMkdir(t *testing.T) {
	got := GenScriptPath(10, "test.sh")
		
	t.Log(got)
}
