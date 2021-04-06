package main

import (
	"fmt"
)

func IsSorted(aux []int, low, high int) bool {
	if low == high {
		fmt.Println("IsSorted low == high true")
		return true
	}
	for index := range aux[low:high] {
		if aux[index+low] > aux[index+low+1] {
			fmt.Println("IsSorted false", aux[index+low], aux[index+low+1])
			return false
		}
	}
	fmt.Println("IsSorted  true")
	return true
}

func MergeSort(aux []int, low, mid, high int) []int {
	var tmp []int
	fmt.Println("MergeSort: ", low, mid, high)
	lowSorted := IsSorted(aux, low, mid)
	highSorted := IsSorted(aux, mid+1, high)
	if !lowSorted {
		newMid := mid - (mid-low)/2
		newHigh := mid
		newLow := low
		tmp = MergeSort(aux, newLow, newMid, newHigh)
	} else {
		tmp = append(tmp, aux[low:mid+1]...)
	}
	if !highSorted {
		newMid := high - (high-mid)/2
		newHigh := high
		newLow := mid + 1
		tmp = append(tmp, MergeSort(aux, newLow, newMid, newHigh)...)
	} else {
		tmp = append(tmp, aux[mid+1:high+1]...)
	}
	var sorted []int
	i := 0
	j := mid - low + 1

	if tmp[j-1] < tmp[j] {
		fmt.Println("already sorted: ", tmp)
		return tmp
	}

	for {
		if tmp[i] < tmp[j] {
			sorted = append(sorted, tmp[i])
			i++
		} else {
			sorted = append(sorted, tmp[j])
			j++
		}
		if i == mid-low+1 {
			sorted = append(sorted, tmp[j:high-low+1]...)
			break
		}
		if j == high-low+1 {
			sorted = append(sorted, tmp[i:mid-low+1]...)
			break
		}
	}
	fmt.Println("final sorted: ", sorted)
	return sorted
}

func main() {
	a := []int{1, 2, 3, 2, 4, 5, 6, 5, 7, 8, 9, 8}
	IsSorted(a, 0, len(a)-1)
	merged := MergeSort(a, 0, len(a)/2, len(a)-1)
	fmt.Println("sorted: ", merged)
}
