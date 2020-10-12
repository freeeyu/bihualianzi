package valid

import (
	"testing"
)

var v Validate = Validate{}

func TestDo(t *testing.T) {

	for _, unit := range []struct {
		p        int
		s        string
		d        []string
		expected bool
	}{
		{1, "商品ID", []string{"require", "numeric", "minsize:4"}, false},
		{12345, "商品ID", []string{"require", "numeric", "maxsize:4"}, false},
	} {
		v.Check(unit.p, unit.s, unit.d)
		if b, _ := v.Check(unit.p, unit.s, unit.d); b != unit.expected {
			t.Errorf("Do(%d,%s,%v) = %t; expected %t", unit.p, unit.s, unit.d, b, unit.expected)
		}
	}
}
