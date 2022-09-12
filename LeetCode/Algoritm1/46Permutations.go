package main

import "fmt"

func main() {
	var nums []int
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	//nums := [...]int{1,2,3}
	permute(nums)
	fmt.Println(nums)
	fmt.Println(permute(nums))

}

type Nums struct {
	Ans  [][]int //для ответа
	Ns   []int
	Used []bool //для проверки использовали ли м уже рание
	Tmp  []int
	Len  int
}

// nums = [1,2,3]
func permute(nums []int) [][]int {

	//init инициализация
	n := new(Nums)
	n.Ns = nums //жұмыс жасаймыз осымен
	n.Len = len(nums)
	n.Tmp = make([]int, 0, n.Len) //временный массив с длиной 0 и вместимостю 3 тобы ансверға салып отыруға
	fmt.Println("TMP")
	fmt.Println(n.Tmp)
	n.Used = make([]bool, n.Len) //булевой массви для проверки то были ли у нас эти числа
	fmt.Println("USED")
	fmt.Println(n.Used)

	n.backtrack()

	return n.Ans
}

func (n *Nums) backtrack() {

	//Goal  ( Constraints2 )
	if len(n.Tmp) == n.Len {
		n.Ans = append(n.Ans, append([]int{}, n.Tmp...)) //когда решени будет столько сколько нужно добавим их в массива ответа
		return
	}

	for i := 0; i < n.Len; i++ {

		//Constraints1
		if !n.Used[i] { //если мы еще неиспользовали это число

			//Choice
			n.Used[i] = true
			n.Tmp = append(n.Tmp, n.Ns[i]) //Go добавим в временный массив нейспользованное число

			n.backtrack()

			n.Used[i] = false            //укажем что в этом индексе элемент испоользован только после того как добавим 3 неповтор элемента только после этого фолс
			n.Tmp = n.Tmp[:len(n.Tmp)-1] //Back и удаляем последни элемент массива
			//будем удалять и делать фолс до первого элемента
		}
	}

}
