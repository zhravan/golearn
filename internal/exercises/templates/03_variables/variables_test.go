package variables

import "testing"

func TestMakePerson(t *testing.T) {
	name, age := MakePerson()
	if name == "" || age == 0 {
		t.Fatalf("expected initialized name and age, got %q and %d", name, age)
	}
}
