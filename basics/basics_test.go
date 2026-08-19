package basics

import "testing"

func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2, 3) = %d, want 5", got)
	}
}

func TestIsEven(t *testing.T) {
	if !IsEven(8) {
		t.Fatal("IsEven(8) = false, want true")
	}
	if IsEven(7) {
		t.Fatal("IsEven(7) = true, want false")
	}
}

func TestWordFrequency(t *testing.T) {
	got := WordFrequency([]string{"go", "go", "test"})
	if got["go"] != 2 {
		t.Fatalf("WordFrequency()['go'] = %d, want 2", got["go"])
	}
	if got["test"] != 1 {
		t.Fatalf("WordFrequency()['test'] = %d, want 1", got["test"])
	}
}

func TestDivide(t *testing.T) {
	got, err := Divide(10, 2)
	if err != nil {
		t.Fatalf("Divide(10,2) unexpected err: %v", err)
	}
	if got != 5 {
		t.Fatalf("Divide(10,2) = %v, want 5", got)
	}

	if _, err := Divide(10, 0); err == nil {
		t.Fatal("Divide(10,0) expected error, got nil")
	}
}
