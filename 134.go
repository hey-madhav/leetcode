package main

func canCompleteCircuit(gas []int, cost []int) int {
	var totalGas, totalCost, start, gasInTank int

	for i := range gas {
		totalGas += gas[i]
		totalCost += cost[i]
		gasInTank = totalGas - cost[i]

		if gasInTank < 0 {
			start = i + 1
			gasInTank = 0
		}
	}

	if totalGas < totalCost {
		return -1
	}

	return start
}
