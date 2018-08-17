package erru_test

import (
	"testing"

	"github.com/clavoie/erru"
)

func TestNewDiDefs(t *testing.T) {
	defs := erru.NewDiDefs()

	if defs == nil {
		t.Fatal("Expecting non-nil defs")
	}

	defs2 := erru.NewDiDefs()
	if defs[0] == defs2[0] {
		t.Fatal("Not expecting defs to match")
	}
}
