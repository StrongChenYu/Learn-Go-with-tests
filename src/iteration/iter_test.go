package iteration

import "testing"

func TestIter(t *testing.T) {
	got := Repeat("a")
	expected := "aaaaa"

	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}