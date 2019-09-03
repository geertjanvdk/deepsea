package submarine

import "testing"

func TestSubmarine_Dive(t *testing.T) {
	var mySub Submarine

	exp := 1000
	mySub.Submerge(exp)
	if mySub.depth != exp {
		t.Fatal("submarine not at position", exp)
	}
}
