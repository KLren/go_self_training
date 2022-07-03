package func_practice

import (
	"fmt"
	"sync"
	//"self_training/util"
)

//For Test_17
//var A = util.F("func_prac.A")

//func init() {
//	util.F("func_prac.init 1")
//}
//func init() {
//	util.F("func_prac.init 2")
//}

// FuncPractice : 可以在 function 定義 return type 的同時定義 variable name
// 即可省略函數內 return 後面的 variable name
func FuncPractice(n int, m int) (int, string) {
	sum := n + m
	if sum < 10 {
		return sum, "small"
	} else {
		return sum, "big"
	}
}

func deferUtil() func(int) int {
	i := 0
	return func(n int) int {
		i++
		fmt.Println("n= ", n)
		fmt.Printf("第 %v 次被調用\n", i)
		return i
	}
}

func Defer() int {
	f := deferUtil()
	//defer f(1)
	//return f(2)

	//defer f(f(3))
	//return f(2)

	defer f(1)
	defer f(2)
	return f(3)
}

func DeferRecover() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	n := 0
	fmt.Println(3 / n)
}

var (
	c    int
	lock sync.Mutex
)

func PrimeNumber(n int) {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return
		}
	}
	fmt.Printf("%v ", n)
	lock.Lock()
	c++
	lock.Unlock()
}

func GoRoutine() {
	for i := 2; i < 1000; i++ {
		go PrimeNumber(i)
	}
	var key string
	fmt.Scanln(&key)
	fmt.Printf("\n共%v個質數", c)
}

func PushNum(c chan int) {
	for i := 1; i <= 100; i++ {
		c <- i
	}
	close(c)
}

func PushPrimeNumber(n int, c chan int) {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return
		}
	}
	c <- n
}

func Channel() {
	var c1 chan int = make(chan int, 100)
	go PushNum(c1)

	//for i := 0; i < 100; i++ {
	//	v, ok := <-c1
	//	if ok {
	//		fmt.Printf("%v ", v)
	//	} else {
	//		break
	//	}
	//}
	//for v := range c1 {
	//	fmt.Printf("%v ", v)
	//}

	var c2 chan int = make(chan int, 100)
	for i := 2; i < 100; i++ {
		go PushPrimeNumber(i, c2)
	}

Print:
	for {
		select {
		case v := <-c2:
			fmt.Printf("%v ", v)
		default:
			fmt.Printf("已找到所有質數")
			break Print
		}
	}
}
