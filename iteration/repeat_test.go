package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("should repeat five times", func (t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
	
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("should repeat five times when the repeat value is invalid", func (t *testing.T) {
		repeated := Repeat("a", -1)
		expected := "aaaaa"
	
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}