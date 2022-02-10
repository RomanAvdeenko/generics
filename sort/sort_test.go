package sort_test

import (
	"math"
	_lib "sort"
	"testing"
	. "utils/sort"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var float64s = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
var strings = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}

func TestIsSorted(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		t.Run("unsorted", func(t *testing.T) {
			s := ints[:]
			if IsSortedAsc(s) {
				t.Error("This is not a sorted slice")
			}
		})
		t.Run("sorted", func(t *testing.T) {
			data := ints
			s := data[:]
			_lib.Ints(s)
			if !IsSortedAsc(s) {
				t.Error("This is a ascending sorted slice")
			}
			Reverse(s)
			if !IsSortedDsc(s) {
				t.Error("This is a descending  sorted slice")
			}
		})
	})
	t.Run("float64", func(t *testing.T) {
		t.Run("unsorted", func(t *testing.T) {
			s := float64s[:]
			res := IsSortedAsc(s)
			if res {
				t.Error("This is not a sorted slice")
			}
		})
		t.Run("sorted", func(t *testing.T) {
			data := float64s
			s := data[:]
			_lib.Float64s(s)
			if !IsSortedAsc(s) {
				t.Error("This is a ascending sorted slice")
			}
			Reverse(s)
			if !IsSortedDsc(s) {
				t.Error("This is a descending  sorted slice")
			}
		})
	})
	t.Run("string", func(t *testing.T) {
		t.Run("unsorted", func(t *testing.T) {
			s := strings[:]
			if IsSortedAsc(s) {
				t.Error("This is not a sorted slice")
			}
		})
		t.Run("sorted", func(t *testing.T) {
			data := strings
			s := data[:]
			_lib.Strings(s)
			if !IsSortedAsc(s) {
				t.Error("This is a ascending sorted slice")
			}
			Reverse(s)
			if !IsSortedDsc(s) {
				t.Error("This is a descending  sorted slice")
			}
		})
	})
}

func TestSort(t *testing.T) {
	t.Run("Quick sort", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			//s := []int{3, 4, 1, 2, 5, 7, -1, 0}
			//s := []int{4, 2, 0, -1, -3}

			data := ints
			s := data[:]
			SortAsc(s)
			if !IsSortedAsc(s) {
				t.Error("It should have been ascending sorted")
			}
			data = ints
			s = data[:]
			SortDesc(s)
			if !IsSortedDsc(s) {
				t.Error("It should have been descending sorted")
			}
			t.Run("float64", func(t *testing.T) {
				//data := float64s
				//s := data[:]
				s := []float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, 0, 0, math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
				SortAsc(s)
				if !IsSortedAsc(s) {
					t.Error("It should have been ascending sorted")
				}
				//data = float64s
				//s = data[:]
				s = []float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, 0, 0, math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
				SortDesc(s)
				if !IsSortedDsc(s) {
					t.Error("It should have been descending sorted")
				}
			})
			t.Run("string", func(t *testing.T) {
				data := strings
				s := data[:]
				SortAsc(s)
				if !IsSortedAsc(s) {
					t.Error("It should have been ascending sorted")
				}
				data = strings
				s = data[:]
				SortDesc(s)
				if !IsSortedDsc(s) {
					t.Error("It should have been descending sorted")
				}
			})
		})
	})
	t.Run("Bouble sort", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			data := ints
			s := data[:]
			SortBoubleAsc(s)
			if !IsSortedAsc(s) {
				t.Error("It should have been ascending sorted")
			}
			data = ints
			s = data[:]
			SortBoubleDsc(s)
			if !IsSortedDsc(s) {
				t.Error("It should have been descending sorted")
			}
		})
		t.Run("float64", func(t *testing.T) {
			data := float64s
			s := data[:]
			SortBoubleAsc(s)
			if !IsSortedAsc(s) {
				t.Error("It should have been ascending sorted")
			}
			data = float64s
			s = data[:]
			SortBoubleDsc(s)
			if !IsSortedDsc(s) {
				t.Error("It should have been descending sorted")
			}
		})
		t.Run("string", func(t *testing.T) {
			data := strings
			s := data[:]
			SortBoubleAsc(s)
			if !IsSortedAsc(s) {
				t.Error("It should have been ascending sorted")
			}
			data = strings
			s = data[:]
			SortBoubleDsc(s)
			if !IsSortedDsc(s) {
				t.Error("It should have been descending sorted")
			}
		})
	})
}

func BenchmarkSortInt1K_CoreLIB(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		_lib.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkSortInt1K_Local_Quick(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		SortAsc(data)
		b.StopTimer()
	}
}

func BenchmarkSortInt1K_Local_Bouble(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		SortAsc(data)
		b.StopTimer()
	}
}
