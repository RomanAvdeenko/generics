package sort

import (
	"constraints"
)

type compareFunc[T constraints.Ordered] func(T, T) bool

func less[T constraints.Ordered](a, b T) bool {
	if a < b {
		return true
	}
	return false
}

func more[T constraints.Ordered](a, b T) bool {
	if a > b {
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
		if cf(s[i], s[i+1]) {
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
	qSort(s, compareFunc[T](less[T]), compareFunc[T](more[T]))
}

func SortDesc[T constraints.Ordered](s []T) {
	qSort(s, compareFunc[T](more[T]), compareFunc[T](less[T]))
}

func qSort[T constraints.Ordered](s []T, less compareFunc[T], more compareFunc[T]) {
	if len(s) < 2 {
		return
	}
	p := arrange(s, less, more)
	qSort(s[:p], less, more)
	qSort(s[p:], less, more)
}

func arrange[T constraints.Ordered](s []T, less compareFunc[T], more compareFunc[T]) int {
	l := 0
	r := len(s) - 1
	var p T
	// Problem  when p is NaN
	p = s[len(s)/2]
	//p := (s[l] + s[len(s)/2] + s[r]) /3 !!! DIV is not realized
	//pi := len(s) / 2
	//	c := uint32(0)

	//	log.Printf("S: l=%v[%v], r=%v[%v] %v %v\n", l, s[l], r, s[r], p, s)
	//log.Printf("S: l=%v[%v], r=%v[%v] p=[%v]%v %v\n", l, s[l], r, s[r], pi, p, s)
	for {
		//		log.Printf("%v: l=%v[%v], r=%v[%v] %v\n", c, l, s[l], r, s[r], s)
		for ; less(s[l], p); l++ {
			//			log.Println("shift to the left")
		}
		for ; more(s[r], p); r-- {
			//			log.Println("shift to the right")
		}
		//log.Printf("after shifting: l=%v[%v], r=%v[%v]\n", l, s[l], r, s[r])
		if l >= r {
			//			log.Println("break:")
			break
		}
		if s[l] == s[r] {
			l++
		}
		swap(s, l, r)
		//		log.Printf("swap: l=%v[%v], r=%v[%v]\n", l, s[l], r, s[r])
		//		c++
	}
	//	log.Printf("end:%v: l=%v[%v], r=%v[%v] %v\n", c, l, s[l], r, s[r], s)
	return r
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
			if cf(s[i], s[j]) {
				swap(s, i, j)
			}
		}
	}
}
