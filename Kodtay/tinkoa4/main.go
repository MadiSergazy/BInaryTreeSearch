package main

import (
	"fmt"
	"math"
)

type Pnt struct {
	x float64
	y float64
}

type Trigon [3]int
type Soltri []Trigon

type Solution struct {
	N     int
	count int
	S     float64
	sol   Soltri
}

var Rep map[pair]Solution

func pnt(i, N int) Pnt {
	var p Pnt
	p.x = math.Cos(2 * math.Pi * float64(i) / float64(N))
	p.y = math.Sin(2 * math.Pi * float64(i) / float64(N))
	return p
}

func Area(i, j, k, N int) float64 {
	return math.Abs(math.Sin(2*math.Pi*float64(j-i)/float64(N))+math.Sin(2*math.Pi*float64(k-j)/float64(N))+math.Sin(2*math.Pi*float64(i-k)/float64(N))) / 2
}

type pair struct {
	first  int
	second int
}

func S(d pair, st Soltri, N int) float64 {
	if d.second < 3 {
		return 0
	}
	if sl, ok := Rep[d]; ok {
		for _, t := range sl.sol {
			t[0] += d.first
			t[1] += d.first
			t[2] += d.first
			st = append(st, t)
		}
		return sl.S
	}

	maxS := 0.0
	var trs Soltri

	i := d.first
	k := d.first + d.second - 1
	for j := int(math.Max(float64(i+k)/2, float64(i+1))); j < int(math.Min(float64(i+k)/2+3, float64(k))); j++ {
		var cur Soltri
		t := [3]int{i, j, k}
		cur = append(cur, t)
		pS := Area(i, j, k, N)
		a := [4]pair{
			pair{d.first, i - d.first},
			pair{i + 1, j - i - 1},
			pair{j + 1, k - j - 1},
			pair{k + 1, d.second - k - 1}}
		pS += S(a[0], cur, N) + S(a[1], cur, N) + S(a[2], cur, N) + S(a[3], cur, N)
		if pS > maxS {
			maxS = pS
			trs = cur
		}
	}
	sol := Solution{N: N, count: d.second, S: maxS}
	for _, m := range trs {
		st = append(st, m)
		for i, tg := range m {
			m[i] = tg - d.first
		}
		sol.sol = append(sol.sol, m)
	}
	Rep[d] = sol

	return maxS
}

func Sol(st Soltri, N int) float64 {
	maxS := 0.0
	var trs Soltri
	i := 0
	for j := int(math.Max(float64(N/3-2), 1)); j < int(math.Max(float64(N/3+2), float64(N))); j++ {
		for k := int(math.Max(float64(2*N/3-2), float64(j+1))); k < int(math.Max(float64(2*N/3+2), float64(N))); k++ {
			var cur Soltri
			t := [3]int{i, j, k}
			cur = append(cur, t)
			pS := Area(i, j, k, N)
			a := [3]pair{
				pair{i + 1, j - i - 1},
				pair{j + 1, k - j - 1},
				pair{k + 1, N - k - 1}}
			pS += S(a[0], cur, N) + S(a[1], cur, N) + S(a[2], cur, N)
			if pS > maxS {
				maxS = pS
				trs = cur
			}
		}
	}
	for _, m := range trs {
		st = append(st, m)
	}
	return maxS
}

/*
func S(N int, sol Soltri) float64 {
	sol = sol[:0]
	return Sol(sol, N)
}*/

func Sn(N int) float64 {
	return float64(N) * math.Sin(2*math.Pi/float64(N)) / 2
}

func main() {
	Rep = make(map[pair]Solution)
	//	fmt.Printf("%.6f\n", math.Fixed)
	for _, N := range []int{250, 500} {
		sol := Soltri{}

		s := S(Sol(sol, N), sol, N)

		fmt.Printf("%d %10.6f %10.6f\n", N, s, s/Sn(N))
		for _, t := range sol {
			fmt.Println(t)
		}
	}
}

/*
package main

import (
	"fmt"
	"math"
	"sort"
)

type pnt struct {
	x, y float64
}

type trigon [3]int
type soltri []trigon

type solution struct {
	N, count int
	S        float64
	sol      soltri
}

var Rep map[pair]solution

// pair - N и количество вершин (отсчет c 0)
type pair struct {
	N, count int
}

func pnt(i, N int) pnt {
	p := pnt{}
	p.x = math.Cos(2*math.Pi*float64(i)/float64(N))
	p.y = math.Sin(2*math.Pi*float64(i)/float64(N))
	return p
}

// Площадь треугольника
func area(i, j, k, N int) float64 {
	return math.Abs(math.Sin(2*math.Pi*float64(j-i)/float64(N)) + math.Sin(2*math.Pi*float64(k-j)/float64(N)) + math.Sin(2*math.Pi*float64(i-k)/float64(N))) / 2
}

// pair - начальная вершина и их количество
func S(d pair, st soltri, N int) float64 {
	if d.count < 3 {
		return 0
	}
	if sl, ok := Rep[d]; ok {
		for _, t := range sl.sol {
			tg := t
			tg[0] += d.N
			tg[1] += d.N
			tg[2] += d.N
			st = append(st, tg)
		}
		return sl.S
	}

	maxS := 0.0
	trs := soltri{}

	i := d.N
	k := d.N + d.count - 1
	for j := int(math.Max(float64((i+k)/2), float64(i+1))); j < int(math.Min(float64((i+k)/2+3), float64(k))); j++ {
		cur := soltri{}
		t := trigon{i, j, k}
		cur = append(cur, t)
		pS := area(i, j, k, N)
		a := []pair{
			{d.N, i - d.N},
			{i + 1, j - i - 1},
			{j + 1, k - j - 1},
			{k + 1, d.count - k - 1},
		}

		pS += S(a[0], cur, N) + S(a[1], cur, N) + S(a[2], cur, N) + S(a[3], cur, N)

		if pS > maxS {
			maxS = pS
			trs = cur
		}
	}

	sol := solution{N, d.count, maxS, soltri{}}
	for _, m := range trs {
		st = append(st, m)
		for i := range m {
			m[i] -= d.N
		}
		sol.sol = append(sol.sol, m)
	}
	Rep[d] = sol

	return maxS
}

func sol(st soltri, N int) float64 {
	maxS := 0.0
	trs := soltri{}
	i := 0
	for j := int(math.Max(float64(N/3-2), float64(1))); j < int(math.Max(float64(N/3+2), float64(N))); j++ {
		for k := int(math.Max(float64(2*N/3-2), float64(j+1))); k < int(math.Max(float64(2*N/3+2), float64(N))); k++ {
			cur := soltri{}
			t := trigon{i, j, k}
			cur = append(cur, t)
			pS := area(i, j, k, N)
			a := []pair{
				{i + 1, j - i - 1},
				{j + 1, k - j - 1},
				{k + 1, N - k - 1},
			}
			pS += S(a[0], cur, N) + S(a[1], cur, N) + S(a[2], cur, N)
			if pS > maxS {
				maxS = pS
				trs = cur
			}
		}
	}
	for _, m := range trs {
		st = append(st, m)
	}
	return maxS
}

func S(N int, sol soltri) float64 {
	sol = soltri{}
	return sol(sol, N)
}

func Sn(N int) float64 {
	return float64(N) * math.Sin(2*math.Pi/float64(N)) / 2
}

func main() {
	Rep = make(map[pair]solution)
	fmt.Printf("%6s: %10s %10s\n", "N", "S", "S/Sn")
	for _, N := range []int{250, 500} {
		sol := soltri{}
		s := S(N, sol)

		fmt.Printf("%6d: %10.6f %10.6f\n", N, s, s/Sn(N))
		for _, t := range sol {
			sort.Ints(t[:])
			fmt.Println(t)
		}
	}
}
*/
