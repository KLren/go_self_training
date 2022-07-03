package main

import (
	"fmt"
	"math/rand"
	functest "self_training/func_practice"
	ptrtest "self_training/ptr_practice" //自定義函數，開頭是別名(alias)
	sttest "self_training/struct_practice"
	"self_training/util" // Test_17 此行需加註解
	"strconv"
	"testing" //使用測試框架
	"time"
	// structtest "self_training/struct_practice"
)

var globalVariable int = 10

//For Test_17
//var A = util.F("main.A")

//For Test_17
//func init() {
//	util.F("main 1")
//}
//func init() {
//	util.F("main 2")
//}

// Variable and Constant practice
func Test_05(t *testing.T) {
	var v1 int
	var v2 int = 3
	v3 := 4
	var (
		v4     = 5
		v5 int = 6
	)
	fmt.Printf("v1=%v, v2=%v, v3=%v, v4=%v, v5=%v\n", v1, v2, v3, v4, v5)
	const (
		c1 = 1    // 第0行
		c2 = iota // 第1行，iota為當前行數
		c3        // 第2行。不定義，值會和前行相同(iota)
		c4 = 9
		c5 = "abc"
		c6 = iota //第5行
	)
	fmt.Printf("c1=%v, c2=%v, c3=%v, c4=%v, c5=%v, c6=%v\n", c1, c2, c3, c4, c5, c6)
	globalVariable = 11
	fmt.Println("Global variable:", globalVariable)
}

// Data type practice
func Test_06(t *testing.T) {
	var (
		i0 int     = 2 // int = int32
		i1 uint16  = 3
		i2 int8    = -5
		i3         = 0b11 // Binary
		i4         = 0o27 // Octal
		i5         = 0x1F // hexadecimal
		f1         = 3.14 // float = float64
		f2 float32 = 5.123
	)
	fmt.Println("Integer:")
	fmt.Printf("i0=%v, type=%T,\n", i0, i0)
	fmt.Printf("i1=%v, type=%T,\n", i1, i1)
	fmt.Printf("i2=%v, type=%T,\n", i2, i2)
	fmt.Printf("Binary i3=%v, type=%T,\n", i3, i3)
	fmt.Printf("Octal i4=%v, type=%T,\n", i4, i4)
	fmt.Printf("hexadecimal i5=%v, type=%T,\n", i5, i5)

	fmt.Println("Float:")
	fmt.Printf("f1=%v, type=%T,\n", f1, f1)
	fmt.Printf("f2=%v, type=%T,\n", f2, f2)
	var (
		s1 byte = 98
		s2      = '8'
		s3      = '帥'
		s4      = 97
		s5 rune = 99
	)
	fmt.Println("character:")
	fmt.Printf("s1=%v, 對應字=%c, type=%T,\n", s1, s1, s1)
	fmt.Printf("s2=%v, 對應字=%c, type=%T,\n", s2, s2, s2)
	fmt.Printf("s3=%v, 對應字=%c, type=%T,\n", s3, s3, s3)
	fmt.Printf("s4=%v, 對應字=%c, type=%T,\n", s4, s4, s4)
	fmt.Printf("s5=%v, 對應字=%c, type=%T,\n", s5, s5, s5)
	s5++
	fmt.Printf("s5=%v, 對應字=%c, type=%T,\n", s5, s5, s5)

	var (
		b bool   //boolean 預設值為 false
		s string = "abc"
	)
	fmt.Printf("b=%v, type=%T,\n", b, b)
	fmt.Printf("s=%v, type=%T, 長度為=%v\n", s, s, len(s))
	q := `var (
		b bool   //boolean 預設值為 false
		s string = "abc"
	)`
	fmt.Printf("q=%v\ntype=%T,\n", q, q)
}

// Pointer practice
func Test_07(t *testing.T) {
	// function parameter pass by pointer
	// 傳遞 pointer 會影響 function 外的 variable
	a := 1
	fmt.Println("address to a: ", &a, "\nvalue of a: ", a)
	ptrtest.PtrPractice(&a)

	// 建立新 pointer (create empty pointer)
	ptr := new(int)
	ptr2 := new(string)
	fmt.Println("address of ptr point to: ", ptr, "\nvalue of ptr point to: ", *ptr)
	fmt.Println("address of ptr2 point to: ", ptr2, "\nvalue of ptr2 point to: ", *ptr2)
}

// Print format practice
func Test_08(t *testing.T) {
	a := 97
	// %U: Unicode
	// %c: Character to unicode
	// %q: Character with quotes
	fmt.Printf("%%U: %U\n%%c: %c\n%%q: %q\n", a, a, a)

	b := 3.14
	// %b = {int_1}*2^{int_2}, ex: 7070651414971679p-51 => 707065... is int_1, -51 is int_2
	// %x: Scientific notation in hexadecimal
	fmt.Printf("%%f: %f\n%%.2f: %.2f\n%%10f: %10f\n", b, b, b)
	fmt.Printf("%%b: %b\n", b)
	fmt.Printf("%%e: %e\n%%x: %x\n", b, b)
	fmt.Printf("%%t: %t\n", b == 314)

	s := "go training"
	fmt.Printf("%%s: %s\n%%q: %q\n%%x: %x\n", s, s, s)
	fmt.Printf("%%p: %p\n", &s) //format in 0x, used for pointer
}

// Operation practice
func Test_09(t *testing.T) {
	// 左移、右移運算
	var (
		b  uint8 = 0b00111100
		b1 uint8 = 0b11001111
	)
	fmt.Printf("b>>2: %b\n", b>>2)
	fmt.Printf("b<<2: %b\n", b<<2)

	fmt.Printf("b&b1: %b\n", b&b1) //交集 intersection
	fmt.Printf("b|b1: %b\n", b|b1) //聯集 union
	fmt.Printf("b^b1: %b\n", b^b1) //差集 difference
}

// If...else practice
func Test_11(t *testing.T) {
	var a int = 15
	// fmt.Scanln(&a)
	if a < 10 {
		fmt.Println("abc")
	} else if a < 20 {
		fmt.Println("def")
	} else {
		fmt.Println("ghi")
	}
}

// Switch practice
func Test_12(t *testing.T) {
	a := 1
	// fmt.Scanln(&a)
	switch {
	case a < 3:
		fmt.Println("abc")
	case a < 6:
		fmt.Println("def")
	default:
		fmt.Println("ghi")
	}

	b := 3
	// fmt.Scanln(&b)
	switch b {
	case 2:
		fmt.Println("aaa")
	case 4:
		fmt.Println("bbb")
	default:
		fmt.Println("ccc")
	}
}

//"for" loop practice
func Test_13(t *testing.T) {
	i := 1
	for {
		fmt.Print(i, " ")
		if i == 5 {
			fmt.Println()
			break
		}
		i++
	}

	i = 1
	for i < 7 {
		fmt.Print(i, " ")
		i += 2
	}
	fmt.Println()

	for i = 0; i < 8; i += 2 {
		if i == 4 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
}

// Lable & goto practice
func Test_14(t *testing.T) {
	fmt.Println("Lable")
	k := 0
outside:
	for i := 0; i < 3; i++ {
		for j := 1; j <= 10; j++ {
			k = i*10 + j
			fmt.Print(k, " ")
			if k == 19 {
				break outside
			}
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("Go to")
	fmt.Println("1")
	if i := 1; i == 1 {
		goto four
	}
	fmt.Println("2")
	fmt.Println("3")
four:
	fmt.Println("4")
	fmt.Println("5")
}

// Function practice
func Test_15(t *testing.T) {
	s1, s2 := functest.FuncPractice(3, 4)
	fmt.Println(s1, s2)
	//fmt.Printf("func: %v, type: %T \n", functest.FuncPractice, functest.FuncPractice)

	//匿名函式
	f2 := func(m, n int) (diff int) {
		diff = m - n
		return
	}(5, 3)

	fmt.Println("diff func: ", f2)
}

// "defer" function
func Test_16(t *testing.T) {
	functest.Defer()
	functest.DeferRecover()
	fmt.Println("After defer recover")
}

//"init" function practice
func Test_17(t *testing.T) {
	fmt.Println("Set init at top of the file.")
}

// Array practice
func Test_19(t *testing.T) {
	// 可以忽略等號左邊的[5]int不寫，等號右邊的長度可以用[...]int讓程式自動判斷
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 2維array (2 dimension)
	arr2 := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	for i, v := range arr2 {
		for j, k := range v {
			fmt.Printf("arr[%v][%v]=%v ", i, j, k)
		}
		fmt.Println()
	}
}

// Slice Practice
func Test_20(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	var s []int = arr[1:len(arr)]
	fmt.Printf("arr: %v\ns: %v\n", arr, s)
	s[0] = 99

	s = append(s, 6, 7)
	fmt.Printf("after append\narr: %v\ns: %v\n", arr, s)

	s2 := make([]int, 3, 5) //3 is len, 5 is cap
	fmt.Printf("s2: %v\n", s2)

	sli := []string{}
	sli = append(sli, "abc", "def", "ghi")
	for i, v := range sli {
		fmt.Printf("index: %v, value: %v\n", i, v)
	}
	s3 := []int{1, 1}
	copy(s2, s3) //將後者複製到前者，僅複製元素，不複製長度
	fmt.Printf("s2: %v\n", s2)

	str := "hello"
	fmt.Printf("[]byte(str): %s\n", []byte(str))

	util.ForTest20("abc", "def", "ghi")
}

func Test_21(t *testing.T) {
	var m map[string]string
	fmt.Printf("m==nil?: %v\n", m == nil)

	m1 := make(map[string]string, 2)
	m1["a1"] = "b1"
	m1["a2"] = "b2"
	fmt.Printf("m1: %v\n", m1)

	m2 := map[string]string{
		"c1": "d1",
		"c2": "d2",
	}
	fmt.Printf("m2: %v\n", m2)

	v, ok := m2["c3"]
	if ok {
		fmt.Printf("v: %v\n", v)
	} else {
		fmt.Println("not existent")
	}

	delete(m1, "a1")
	fmt.Printf("m1 after delete: %v\n", m1)
	//刪除全部
	//m1 = nil
	//m2 = make(map[string]string)
	//fmt.Printf("m1 after delete: %v\nm2 after delete: %v\n", m1, m2)

	for k, v := range m2 {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}
}

func Test_22(t *testing.T) {
	type myUint16 uint16
	var u1000 uint16 = 1000
	var myu myUint16 = myUint16(u1000)
	fmt.Printf("myu: %v, type of myu: %T\n", myu, myu)

	//alias
	type aUint16 = uint16
	var au aUint16 = u1000
	fmt.Printf("au: %v, type of au: %T\n", au, au)
}

// Struct practice
func Test_23(t *testing.T) {
	sttest.Construct()
}

// Struct method
func Test_24(t *testing.T) {
	gm1 := sttest.GoStruct{}
	gm1.Name = "Ken"
	gm1.PrintName()

	gm1.SetAge(5)
	gm1.PrintAge()
}

// "interface" Practice
func Test_25(t *testing.T) {
	tm := sttest.TextMsg{}
	sttest.SetMsgType(&tm)
	im := sttest.ImgMsg{}
	sttest.SetMsgType(&im)

	n1 := 1
	n1Interface := interface{}(n1)
	n2, ok := n1Interface.(int)
	if ok {
		fmt.Println("n2: ", n2)
	} else {
		fmt.Println("assertion fail")
	}
}

// Goroutine practice
func Test_26(t *testing.T) {
	functest.GoRoutine()
}

// Channel practice
func Test_27(t *testing.T) {
	functest.Channel()
}

// Random number practice
func Test_29(t *testing.T) {
	for i := 0; i < 10; i++ {
		//seed := rand.Seed(time.Now().UnixNano())
		rand.Seed(time.Now().UnixNano())
		fmt.Println(rand.Intn(10) + 1)
		//seed++
	}
}

// 字串轉換
func Test_30(t *testing.T) {
	i1 := 123
	s1 := "hello"
	s2 := fmt.Sprintf("%daaa%s", i1, s1)
	fmt.Println("s2=", s2)

	var (
		i2 int
		s3 string
	)
	n, err := fmt.Sscanf(s2, "%daaa%s", &i2, &s3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("解析了%d個資料\n", n)
	fmt.Printf("i2= %v, s3= %s\n", i2, s3)

	s4 := strconv.FormatInt(123, 4)
	fmt.Printf("s4= %v\n", s4)
	u1, err := strconv.ParseUint(s4, 4, 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("u1= %v\n", u1)
}
