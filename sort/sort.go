package sort

import (
	"constraints"
)

type compareFunc[T constraints.Ordered] func([]T, int, int) bool

func less[T constraints.Ordered](s []T, i, j int) bool {
	if s[i] < s[j] {
		return true
	}
	return false
}

func more[T constraints.Ordered](s []T, i, j int) bool {
	if s[i] > s[j] {
		return true
	}
	return false
}

type swapFunc[T any] func([]T, int, int)

func swap[T any](s []T, i, j int) {
	s[i], s[j] = s[j], s[i]
}

func IsSortedAsc[T constraints.Ordered](s []T) bool {
	return isSorted(s, compareFunc[T](more[T]))
}

func IsSortedDsc[T constraints.Ordered](s []T) bool {
	return isSorted(s, compareFunc[T](less[T]))
}

func isSorted[T constraints.Ordered](s []T, cf compareFunc[T]) bool {
	for i := 0; i < len(s)-1; i++ {
		if cf(s, i, i+1) {
			return false
		}
	}
	return true
}

func swapRange[T any](s []T, a, b, n int) {
	for i := 0; i < n; i++ {
		s[a+i], s[b+i] = s[b+i], s[a+i]
	}
}

func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		swap(s, i, j)
	}
}

// Sorting block
// quick sort
func SortAsc[T constraints.Ordered](s []T) {
	qSort(s, compareFunc[T](more[T]))
}

func SortDesc[T constraints.Ordered](s []T) {
	qSort(s, compareFunc[T](less[T]))
}

func qSort[T constraints.Ordered](s []T, cf compareFunc[T]) {
	switch len(s) {
	case 0, 1:
		return
	case 2:
		// swap if needed
		if cf(s, 0, 1) {
			swap(s, 0, 1)
		}
		return
	default:
		p := arrange(s, cf)
		qSort(s[:p], cf)
		qSort(s[p:], cf)
	}
}

func arrange[T constraints.Ordered](s []T, cf compareFunc[T]) int {
	l := 0
	p := len(s) - 1 // Pivot is Right

	for l < p {
		if cf(s, l, p) {
			tmp := s[l]
			for i := l; i < p; i++ {
				s[i] = s[i+1]
			}
			s[p] = tmp
			p--
		} else {
			l++
		}
	}
	return p
}

// bouble sort
func SortBoubleAsc[T constraints.Ordered](s []T) {
	boubleSort(s, compareFunc[T](more[T]))
}

func SortBoubleDsc[T constraints.Ordered](s []T) {
	boubleSort(s, compareFunc[T](less[T]))
}

func boubleSort[T constraints.Ordered](s []T, cf compareFunc[T]) {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if cf(s, i, j) {
				swap(s, i, j)
			}
		}
	}
}
