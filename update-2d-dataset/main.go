package main

func do(current map[string]int, new map[string]int) map[string]int {
	for item, amount := range new {
		if _, ok := current[item]; !ok {
			current[item] = amount
		} else {
			current[item] = current[item] + amount
		}
	}
	return current

}
