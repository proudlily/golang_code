package main

import "fmt"

/*
从凡是最小公因子的倍数都能找到解得到启发，算法的实质是在枚举数集 { (A*x+B*y) mod A}的
过程中构造所题目所需要的解。
*/
type state struct {
	s1 int
	s2 int
}

func (s state) String() string {
	return fmt.Sprintf("{%v,%v}", s.s1, s.s2)
}

func generator(A_Cap,B_Cap int)[]state{
	if A_Cap<B_Cap{
		A_Cap,B_Cap=B_Cap,A_Cap
	}
	s1 := A_Cap
	var r []state
	for {
		for s1 > B_Cap {
			r = append(r, state{s1, 0})
			s1 = s1 - B_Cap
			r = append(r, state{s1, B_Cap})

		}
		r = append(r, state{s1, 0})
		r = append(r, state{0, s1})
		r = append(r, state{A_Cap, s1})
		if s1 == B_Cap {
			break
		}
		s1 = A_Cap - (B_Cap - s1)
		r = append(r, state{s1, B_Cap})
	}
	return r
}
func search(r []state, target int) []state {
	for i, s := range r {
		if s.s1 == target || s.s2 == target {
			return r[:i+1]
		}
	}
	return nil
}

func main() {
	var A_Cap, B_Cap int
	fmt.Print("Input Capacity of two bottles >> ")
	fmt.Scan(&A_Cap, &B_Cap)
	for {
		fmt.Print("Input Target >> ")
		var t int
		fmt.Scan(&t)
		r:=generator(A_Cap,B_Cap)
		r = search(r, t)
		if r == nil {
			fmt.Println("No Solution!")
			continue
		}
		fmt.Println("solution:")
		for i, s := range r {
			fmt.Printf("%v:\t%v\n",i, s)
		}
	}

}
