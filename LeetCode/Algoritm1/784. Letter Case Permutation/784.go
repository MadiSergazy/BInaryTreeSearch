package main

import "fmt"

func main() {
	fmt.Println(letterCasePermutation("Al"))
}

type Str struct {
	Ans    []string
	S      []byte
	Lenght int
}

func letterCasePermutation(s string) []string {

	str := new(Str)
	str.S = []byte(s)
	str.Lenght = len(str.S)

	str.BackTrack([]byte{}, 0)

	return str.Ans
}

func (s *Str) BackTrack(temp []byte, index int) {

	if len(temp) == s.Lenght {
		s.Ans = append(s.Ans, string(temp))
		//fmt.Println(s.Ans)
	}

	s.BackTrack(append(temp, s.S[index]), index+1)

	if s.S[index] >= 'a' && s.S[index] <= 'z' {
		//temp = append(temp, temp)
		//	fmt.Println(s.S)
		s.BackTrack(append(temp, s.S[index]+'A'-'a'), index+1)
		//temp[index] = byte(temp[index] + 'A' - 'a')
	} else if s.S[index] >= 'a' && s.S[index] <= 'z' {
		//temp[index] = byte(temp[index] + 'a' - 'A')
		//	fmt.Println(s.S)
		s.BackTrack(append(temp, s.S[index]+'a'-'A'), index+1)
	}

}

/*
type Str struct {
	Ans []string //
	S   []byte
	Len uint8
}

func letterCasePermutation(s string) []string {

	//init
	st := new(Str)
	st.S = []byte(s)
	st.Len = uint8(len(st.S))

	st.backtrack([]byte{}, 0)

	return st.Ans
}

func (s *Str) backtrack(tmp []byte, index uint8) {

	if index == s.Len {
		s.Ans = append(s.Ans, string(tmp))
		return
	}

	//0-9
	s.backtrack(append(tmp, s.S[index]), index+1)

	if 'a' <= s.S[index] && s.S[index] <= 'z' {
		s.backtrack(append(tmp, s.S[index]-32), index+1)

	} else if 'A' <= s.S[index] && s.S[index] <= 'Z' {
		s.backtrack(append(tmp, s.S[index]+32), index+1)
	}
}*/

/*

func letterCasePermutation(S string) []string {
	permutations := []string{S}
	for i, v := range S {
		if v >= 'a' && v <= 'z' { //65-90
			length := len(permutations)
			for ip := 0; ip < length; ip++ {
				b := []byte(permutations[ip])
				b[i] = byte(v + 'A' - 'a') //-32
				permutations = append(permutations, string(b))
			}
		} else if v >= 'A' && v <= 'Z' { //97-122
			length := len(permutations)
			for ip := 0; ip < length; ip++ {
				b := []byte(permutations[ip])
				b[i] = byte(v + 'a' - 'A') //+32
				permutations = append(permutations, string(b))
			}
		}
	}
	return permutations
}
*/
