package tests

import (
	"testing"

	"github.com/geertjanvdk/deepsea/submarine"
)

func TestSomething(t *testing.T) {
	(&submarine.Submarine{}).Ping()
	t.Fatal("I simply fail")
}
