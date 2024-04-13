package utils

import (
	"math"
)

func Pagination[T interface{}](d []T, p int, s int) []T {
	cp := math.Ceil(float64(len(d)) / float64(s))
	if p <= int(cp) {
		var rs []T
		for i := (p * s) - s; i < p*s; i++ {
			if i < len(d) {
				rs = append(rs, d[i])
			}
		}
		return rs
	} else {
		return nil
	}
}
