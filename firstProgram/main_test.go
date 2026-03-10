package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 4
	res := count(b, false, false)
	if exp != res {
		t.Errorf("Expected %d, got %d", res, exp)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("zalupa\nezja\naaa")
	exp := 3
	res := count(b, true, false)
	if exp != res {
		t.Errorf("Expected %d, got %d", res, exp)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("banan")
	exp := 5
	res := count(b, false, true)
	if exp != res {
		t.Errorf("Expected %d, got %d", res, exp)
	}
}
