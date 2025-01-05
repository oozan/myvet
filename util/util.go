package util

import (
	"myvet-v2-api/structs"
)

/*
x := 23
i := sort.Search(len(data), func(i int) bool { return data[i] >= x })
if i < len(data) && data[i] == x {
	// x is present at data[i]
} else {
	// x is not present in data,
	// but i is the index where it would be inserted.
}
*/
/*
 * helper function that abstracts the logic of linear search
 * if we need, we can write a similar one to binary search
 */
func FindIndex(n int, f func(int) bool) int {
	for i := 0; i < n; i++ {
		if f(i) {
			return i
		}
	}
	return -1
}

/*
 * helper function that abstracts the logic of binary search
 * implementation stolen directly from sort.Search with some minor tweaks
 */
func FindIndexSorted(n int, f func(int) int) int {
	i, j := 0, n
	for i < j {
		// halfpoint between i and j
		h := int(uint(i+j) >> 1) // // avoid overflow when computing h
		res := f(h)
		// f returns 0 if equal
		if res == 0 {
			return i
		}
		// search from the left side
		if res > 0 {
			j = h
		} else {
			// search from the right side
			i = h + 1
		}
	}
	return -1
}

// ContainsInt returns the index of target int or -1 if it doesn't exist in the slice.
func ContainsInt(vals []int, target int) int {
	for i, val := range vals {
		if val == target {
			return i
		}
	}
	return -1
}

/*
func ContainsAnimal(vals map[structs.Animal]struct{}, target structs.Animal) bool {
	for animal := range vals {
		if animal.AnimalID == target.AnimalID {
			return true
		}
	}
	return false
}*/

func ContainsCustomer(vals []structs.Customer, target structs.Customer) bool {
	for _, val := range vals {
		if val.CustomerID == target.CustomerID {
			return true
		}
	}
	return false
}

func containsDates(vals []structs.Payment, target structs.Payment) bool {
	for _, val := range vals {
		if val.PaymentDate == target.PaymentDate {
			return true
		}
	}
	return false
}

func containsStates(vals []structs.Appointment, target structs.Appointment) bool {
	for _, val := range vals {
		if val.State == target.State {
			return true
		}
	}
	return false
}
