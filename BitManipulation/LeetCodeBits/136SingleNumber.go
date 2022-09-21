package LeetCodeBits

func singleNumber(nums []int) int {

	m := make(map[int]int8, len(nums)/2)
	res := 0

	for _, n := range nums {
		m[n]++         //по ключю n 1ді  қосамыз значенияға
		if m[n] == 1 { //егер ол 1 болса сол санды қосамыз
			res += n
		} else {
			res -= n //әйтпесе ол санды қоспаймыз ответке
		}
	}

	return res
}

/*
func singleNumber(nums []int) int {

	res := 0

	for _, n := range nums {
		res ^= n
	}

	return res
}
*/

/*
func singleNumber(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }else {
        for i := 0; i < len(nums); i++{
            counter := 0
            for j := 0; j < len(nums); j++{
                if nums[j] == nums[i] {
                    counter++
                }
            }
            if counter != 2 {
                return nums[i]
            }
        }

        return nums[0]
    }
}*/
