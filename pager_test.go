package gopager

import (
	"testing"
)

type TestStruct struct {
	ID   uint
	Name string
	Age  uint
}

type TestStructs []TestStruct

func (t TestStructs) Len() int {
	return len(t)
}

func TestPaginater(t *testing.T) {
	obj := TestStructs{
		{
			ID:   1,
			Name: "Mr.A",
			Age:  20,
		},
		{
			ID:   2,
			Name: "Mr.B",
			Age:  30,
		},
		{
			ID:   3,
			Name: "Mr.C",
			Age:  25,
		},
		{
			ID:   4,
			Name: "Mr.D",
			Age:  33,
		},
		{
			ID:   5,
			Name: "Mr.E",
			Age:  34,
		},
	}
	expectMaxPageNum := 3

	// TEST: NewPaginater()
	p := NewPaginater(TestStructs{}, 2)
	if p.MaxPage() != 0 {
		t.Fatalf("expect: 0, but actual: %v\n", p.MaxPage())
	}

	// TEST: MaxPage()
	p = NewPaginater(obj, 2)
	if p.MaxPage() != expectMaxPageNum {
		t.Fatalf("expect: %v, but actual: %v\n", expectMaxPageNum, p.MaxPage())
	}

	// TEST: Current()
	actual := TestStructs{}
	p.Current(&actual)
	if len(actual) != 0 {
		t.Fatalf("expect: 0, but actual: %v\n", actual)
	}

	var currentTests = []struct {
		in  int
		out int
	}{
		{
			in:  1,
			out: 2,
		},
		{
			in:  2,
			out: 2,
		},
		{
			in:  3,
			out: 1,
		},
	}
	for _, ct := range currentTests {
		actual := TestStructs{}
		p.Page(ct.in).Current(&actual)
		t.Logf("idx: %v actual: %+v\n", ct.in, actual)
		if len(actual) != ct.out {
			t.Fatalf("expect: %v, but actual: %v\n", ct.out, len(actual))
		}
	}
	// TEST: Page() & CurrentPage()
	p.Page(-1)
	if p.CurrentPage() != 0 {
		t.Fatalf("expect: 0, but actual: %v\n", p.CurrentPage())
	}
	p.Page(0)

	// TEST: Next()
	var nextTests = []int{2, 2, 1}
	idx := 0
	for p.HasNext() {
		actual := TestStructs{}
		p.Next(&actual)
		t.Logf("curPage: %v actual: %+v\n", p.CurrentPage(), actual)
		if len(actual) != nextTests[idx] {
			t.Fatalf("expect: %v, but actual: %v\n", nextTests[idx], len(actual))
		}
		idx++
	}

	// TEST: Previous()
	var previousTests = []int{2, 2}
	idx = 0
	for p.HasPrevious() {
		actual := TestStructs{}
		p.Previous(&actual)
		t.Logf("curPage: %v actual: %+v\n", p.CurrentPage(), actual)
		if len(actual) != previousTests[idx] {
			t.Fatalf("expect: %v, but actual: %v\n", previousTests[idx], len(actual))
		}
		idx++
	}
}
