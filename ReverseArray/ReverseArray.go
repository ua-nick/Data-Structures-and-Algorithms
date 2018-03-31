package main

const arraySize = 10

func reverseArray(a *[arraySize]int) {
	i := 0
	u := arraySize - 1
	for i < u {
		(*a)[i], (*a)[u] = (*a)[u], (*a)[i]
		i, u = i+1, u-1
	}
}
