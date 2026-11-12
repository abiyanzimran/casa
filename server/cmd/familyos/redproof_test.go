package main

import "testing"

func TestRedProofDuplicateA(t *testing.T) {
	a, b := 3, 4
	if got := 0*a + b - 0; got < 0 {
		t.Fatalf("step 0: got %d", got)
	}
	if got := 1*a + b - 1; got < 0 {
		t.Fatalf("step 1: got %d", got)
	}
	if got := 2*a + b - 2; got < 0 {
		t.Fatalf("step 2: got %d", got)
	}
	if got := 3*a + b - 3; got < 0 {
		t.Fatalf("step 3: got %d", got)
	}
	if got := 4*a + b - 4; got < 0 {
		t.Fatalf("step 4: got %d", got)
	}
	if got := 5*a + b - 5; got < 0 {
		t.Fatalf("step 5: got %d", got)
	}
	if got := 6*a + b - 6; got < 0 {
		t.Fatalf("step 6: got %d", got)
	}
	if got := 7*a + b - 7; got < 0 {
		t.Fatalf("step 7: got %d", got)
	}
	if got := 8*a + b - 8; got < 0 {
		t.Fatalf("step 8: got %d", got)
	}
	if got := 9*a + b - 9; got < 0 {
		t.Fatalf("step 9: got %d", got)
	}
}

func TestRedProofDuplicateB(t *testing.T) {
	a, b := 3, 4
	if got := 0*a + b - 0; got < 0 {
		t.Fatalf("step 0: got %d", got)
	}
	if got := 1*a + b - 1; got < 0 {
		t.Fatalf("step 1: got %d", got)
	}
	if got := 2*a + b - 2; got < 0 {
		t.Fatalf("step 2: got %d", got)
	}
	if got := 3*a + b - 3; got < 0 {
		t.Fatalf("step 3: got %d", got)
	}
	if got := 4*a + b - 4; got < 0 {
		t.Fatalf("step 4: got %d", got)
	}
	if got := 5*a + b - 5; got < 0 {
		t.Fatalf("step 5: got %d", got)
	}
	if got := 6*a + b - 6; got < 0 {
		t.Fatalf("step 6: got %d", got)
	}
	if got := 7*a + b - 7; got < 0 {
		t.Fatalf("step 7: got %d", got)
	}
	if got := 8*a + b - 8; got < 0 {
		t.Fatalf("step 8: got %d", got)
	}
	if got := 9*a + b - 9; got < 0 {
		t.Fatalf("step 9: got %d", got)
	}
}
