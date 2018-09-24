package main

import (
	"testing"
	loadlab "github.com/loadlab/go"
)

func TestLoadlab(t *testing.T) {
	l := loadlab.LoadLab{token: "xxxxxx"}
	t.Errorf(l.sites())
}


