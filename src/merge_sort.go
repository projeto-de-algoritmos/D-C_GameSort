package main

func mergeSort(items []Game, isReverse bool, flag string) []Game {
	var num = len(items)
	if num == 1 {
		return items
	}
	middle := int(num / 2)
	var (
		left  = make([]Game, middle)
		right = make([]Game, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left, isReverse, flag), mergeSort(right, isReverse, flag), isReverse, flag)
}

func merge(left, right []Game, isReverse bool, flag string) (result []Game) {
	result = make([]Game, len(left)+len(right))

	i := 0

	if isReverse == true && flag == "year" {
		for len(left) > 0 && len(right) > 0 {
			if left[0].GetYear() > right[0].GetYear() {
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
			i++
		}
	} else if isReverse == true && flag == "sales" {
		for len(left) > 0 && len(right) > 0 {
			if left[0].GetSales() > right[0].GetSales() {
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
			i++
		}
	} else if isReverse == true && flag == "name" {
		for len(left) > 0 && len(right) > 0 {
			if left[0].GetName() > right[0].GetName() {
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
			i++
		}
	} else if isReverse == false && flag == "sales" {
		for len(left) > 0 && len(right) > 0 {
			if left[0].GetSales() < right[0].GetSales() {
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
			i++
		}
	} else if isReverse == false && flag == "name" {
		for len(left) > 0 && len(right) > 0 {
			if left[0].GetName() < right[0].GetName() {
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
			i++
		}
	} else if isReverse == false && flag == "year" {
		for len(left) > 0 && len(right) > 0 {
			if left[0].GetYear() < right[0].GetYear() {
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
			i++
		}
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
