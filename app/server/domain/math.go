package domain

type MathDomain interface {
	TwoSum(numbers []int, target int) []int
}

type mathDomain struct{}

func newMathDomain() MathDomain {
	return &mathDomain{}
}

func (m *mathDomain) TwoSum(numbers []int, target int) []int {
	numberMap := make(map[int]int)
	for i, number := range numbers {
		numberMap[number] = i
	}

	twoSum := []int{}
	for i, number := range numbers {
		complement := target - number
		if j, present := numberMap[complement]; present {
			if i != j {
				twoSum = append(twoSum, complement)
				twoSum = append(twoSum, number)
				break
			}
		}
	}

	return twoSum
}
