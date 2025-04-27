package main

type WithdrawStrategy interface {
	Withdraw(amount int, available map[int]int) (map[int]int, bool)
}

type BigToSmall struct{}
type SmallToBig struct{}

func (g *BigToSmall) Withdraw(amount int, avavilable map[int]int) (map[int]int, bool) {
	result := make(map[int]int)
	remaining := amount

	// try biggest to smallest
	for _, denom := range []int{100, 20, 5, 1} {
		count := avavilable[denom]
		if count == 0 || denom > remaining {
			continue
		}

		needed := remaining / denom
		use := min(count, needed)

		result[denom] = use
		remaining -= denom * use
	}

	if remaining == 0 {
		return result, true
	}
	return nil, false
}

func (g *SmallToBig) Withdraw(amount int, avavilable map[int]int) (map[int]int, bool) {
	result := make(map[int]int)
	remaining := amount

	// try biggest to smallest
	for _, denom := range []int{1, 5, 20, 100} {
		count := avavilable[denom]
		if count == 0 || denom > remaining {
			continue
		}

		needed := remaining / denom
		use := min(count, needed)

		result[denom] = use
		remaining -= denom * use
	}

	if remaining == 0 {
		return result, true
	}
	return nil, false
}
