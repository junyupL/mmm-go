package mmm

import "golang.org/x/exp/constraints"

func Max[O constraints.Ordered](first O, elements ...O) O {
	for _, element := range elements {
		if element > first {
			first = element
		}
	}
	return first
}

func Min[O constraints.Ordered](first O, elements ...O) O {
	for _, element := range elements {
		if element < first {
			first = element
		}
	}
	return first
}
