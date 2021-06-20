package fileops

import (
	"github.com/jfixby/pin/lang"
	"testing"
	//"path/filepath"
	"path/filepath"
)

func TestPath(t *testing.T) {
	expected := []string{"1", "2", "3", "4", "5"}
	join := filepath.Join(expected...)
	result := PathToArray(join)
	if !lang.EqualArrays(result, expected) {
		t.Errorf("expected %v result %v", expected, result)
	}
}

func TestSplitPath(t *testing.T) {
	input := []string{"1", "2", "3", "4", "5"}
	prefix := filepath.Join(input[0:3]...)
	full := filepath.Join(input...)
	result := SplitPath(full, prefix)
	expected := filepath.Join(input[3:]...)
	if !(result == expected) {
		t.Errorf("expected %v, result %v", expected, result)
	}
}

