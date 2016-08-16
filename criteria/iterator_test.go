// +build unit

package criteria

import (
	"reflect"
	"testing"
)

func TestIterator(t *testing.T) {
	visited := []Expression{}
	l := Field("a")
	r := Literal(5)
	expr := Equals(l, r)
	expected := []Expression{l, r, expr}
	recorder := func(expr Expression) bool {
		visited = append(visited, expr)
		return true
	}
	IteratePostOrder(expr, recorder)
	if !reflect.DeepEqual(expected, visited) {
		t.Errorf("Visited should be %v, but is %v", expected, visited)
	}

	visited = []Expression{}
	recorder = func(expr Expression) bool {
		visited = append(visited, expr)
		if expr == r {
			return false
		}
		return true
	}
	IteratePostOrder(expr, recorder)
	expected = []Expression{l, r}
	if !reflect.DeepEqual(expected, visited) {
		t.Errorf("Visited should be %v, but is %v", expected, visited)
	}

}
