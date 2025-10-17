package channel_directions

import (
	"reflect"
	"testing"
)

func TestDbPool(t *testing.T) {
	t.Run("Queries processed successfully!", func(t *testing.T) {
		got := DbPool()
		want := []int{}
		for i := 1; i <= 50; i++ {
			want = append(want, i*2)
		}

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("Expected %v, got %v", want, got)
		}
	})
}
