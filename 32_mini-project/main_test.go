package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	// Simulasikan panggilan fungsi main
	main()
	
	// Lakukan pengujian
	if got := port; got != "8080" {
		t.Errorf("expected port %s, got %s", "8080", got)
	}
}
