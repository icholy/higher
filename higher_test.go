package higher

import (
	"reflect"
	"strconv"
	"testing"
)

var FilterTable = []struct {
	In  interface{}
	Out interface{}
	Fn  interface{}
}{
	{
		In:  []int{1, 2, 3, 4, 5},
		Out: []int{2, 4},
		Fn:  func(x int) bool { return x%2 == 0 },
	},
}

func TestFilter(t *testing.T) {
	for _, test := range FilterTable {
		out := Filter(test.In, test.Fn)
		if !reflect.DeepEqual(test.Out, out) {
			t.Fatalf("%v should equal %v", out, test.Out)
		}
	}
}

func TestFilterWrapped(t *testing.T) {
	for _, test := range FilterTable {
		out := Wrap(test.In).Filter(test.Fn).Val()
		if !reflect.DeepEqual(test.Out, out) {
			t.Fatalf("%v should equal %v", out, test.Out)
		}
	}
}

var MapTable = []struct {
	In  interface{}
	Out interface{}
	Fn  interface{}
}{
	{
		In:  []int{1, 2, 3, 4},
		Out: []string{"1", "2", "3", "4"},
		Fn:  strconv.Itoa,
	},
}

func TestMap(t *testing.T) {
	for _, test := range MapTable {
		out := Map(test.In, test.Fn)
		if !reflect.DeepEqual(test.Out, out) {
			t.Fatalf("%v should equal %v", out, test.Out)
		}
	}
}

func TestMapWrapped(t *testing.T) {
	for _, test := range MapTable {
		out := Wrap(test.In).Map(test.Fn).Val()
		if !reflect.DeepEqual(test.Out, out) {
			t.Fatalf("%v should equal %v", out, test.Out)
		}
	}
}

var ReduceTable = []struct {
	In  interface{}
	Acc interface{}
	Out interface{}
	Fn  interface{}
}{
	{
		In:  []int{1, 2, 3, 4, 5},
		Acc: 0,
		Out: 15,
		Fn:  func(acc, x int) int { return acc + x },
	},
}

func TestReduce(t *testing.T) {
	for _, test := range ReduceTable {
		out := Reduce(test.In, test.Fn, test.Acc)
		if !reflect.DeepEqual(test.Out, out) {
			t.Fatalf("%v should equal %v", out, test.Out)
		}
	}
}

func TestReduceWrapped(t *testing.T) {
	for _, test := range ReduceTable {
		out := Wrap(test.In).Reduce(test.Fn, test.Acc).Val()
		if !reflect.DeepEqual(test.Out, out) {
			t.Fatalf("%v should equal %v", out, test.Out)
		}
	}
}

var AnyTable = []struct {
	In  interface{}
	Out bool
	Fn  interface{}
}{
	{
		In:  []string{"A", "B", "C", "D"},
		Out: true,
		Fn:  func(s string) bool { return s == "C" },
	},
}

func TestAny(t *testing.T) {
	for _, test := range AnyTable {
		out := Any(test.In, test.Fn)
		if !reflect.DeepEqual(test.Out, out) {
			t.Fatalf("%v should equal %v", out, test.Out)
		}
	}
}

func TestAnyWrapped(t *testing.T) {
	for _, test := range AnyTable {
		out := Wrap(test.In).Any(test.Fn)
		if !reflect.DeepEqual(test.Out, out) {
			t.Fatalf("%v should equal %v", out, test.Out)
		}
	}
}

var EveryTable = []struct {
	In  interface{}
	Out bool
	Fn  interface{}
}{
	{
		In:  []int{2, 4, 6, 8},
		Out: true,
		Fn:  func(x int) bool { return x%2 == 0 },
	},
}

func TestEvery(t *testing.T) {
	for _, test := range EveryTable {
		out := Every(test.In, test.Fn)
		if !reflect.DeepEqual(test.Out, out) {
			t.Fatalf("%v should equal %v", out, test.Out)
		}
	}
}

func TestEveryWrapped(t *testing.T) {
	for _, test := range EveryTable {
		out := Wrap(test.In).Every(test.Fn)
		if !reflect.DeepEqual(test.Out, out) {
			t.Fatalf("%v should equal %v", out, test.Out)
		}
	}
}

var ContainsTable = []struct {
	In  interface{}
	Val interface{}
	Out bool
}{
	{
		In:  []bool{false, false, true},
		Val: true,
		Out: true,
	},
}

func TestContains(t *testing.T) {
	for _, test := range ContainsTable {
		out := Contains(test.In, test.Val)
		if !reflect.DeepEqual(test.Out, out) {
			t.Fatalf("%v should equal %v", out, test.Out)
		}
	}
}

func TestContainsWrapped(t *testing.T) {
	for _, test := range ContainsTable {
		out := Wrap(test.In).Contains(test.Val)
		if !reflect.DeepEqual(test.Out, out) {
			t.Fatalf("%v should equal %v", out, test.Out)
		}
	}
}

var FindTable = []struct {
	In  interface{}
	Out interface{}
	Fn  interface{}
}{
	{
		In:  []int{1, 2, 3, 4, 5, 6},
		Out: 4,
		Fn:  func(x int) bool { return x > 3 },
	},
}

func TestFind(t *testing.T) {
	for _, test := range FindTable {
		out := Find(test.In, test.Fn)
		if !reflect.DeepEqual(test.Out, out) {
			t.Fatalf("%v should equal %v", out, test.Out)
		}
	}
}

func TestFindWrapped(t *testing.T) {
	for _, test := range FindTable {
		out := Wrap(test.In).Find(test.Fn)
		if !reflect.DeepEqual(test.Out, out) {
			t.Fatalf("%v should equal %v", out, test.Out)
		}
	}
}
