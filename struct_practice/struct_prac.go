package struct_practice

import "fmt"

type GoStruct struct {
	Name string `json:"name"` //ex: In json file the name will be Name. (struct label)
	age  int
}

type InhStruct struct {
	GoStruct
	Height int
}

type PtrInhStruct struct {
	*GoStruct
	Weight int
}

func (a GoStruct) PrintName() {
	fmt.Println("Name in object:", a.Name)
}

func (s *GoStruct) SetAge(input int) {
	s.age = input
}

func (a GoStruct) PrintAge() {
	fmt.Println("age in object:", a.age)
}

func Construct() {
	var g1 GoStruct = GoStruct{
		Name: "Ash",
	}
	g1.age = 12
	var g2 *GoStruct = &GoStruct{
		Name: "Ben",
	}
	g2.age = 15 // equal to: (*g2).age = 15
	fmt.Println("g1:", g1)
	fmt.Println("g2:", g2)

	var i1 InhStruct = InhStruct{
		GoStruct: GoStruct{
			Name: "aaa",
			age:  g1.age,
		},
		Height: 190,
	}
	var i2 *PtrInhStruct = &PtrInhStruct{
		GoStruct: &GoStruct{}, // or g2 but not g1
		Weight:   80,
	}
	i2.Name = "Cat" // equal to i2.GoStruct.Name
	fmt.Println("i1:", i1)
	fmt.Println("i2:", i2)
}

// TextMsg : For Test_25
type TextMsg struct {
	Type string
	Text string
}

func (tm *TextMsg) SetText() {
	tm.Text = "Text_set"
}

func (tm *TextMsg) SetType() {
	tm.Type = "Text_type_set"
}

type ImgMsg struct {
	Type string
	Img  string
}

func (tm *ImgMsg) SetImg() {
	tm.Img = "Image_set"
}

func (tm *ImgMsg) SetType() {
	tm.Type = "Image_type_set"
}

type MsgInter interface {
	SetType()
}

func SetMsgType(m MsgInter) {
	m.SetType()
	switch mptr := m.(type) {
	case *TextMsg:
		mptr.SetText()
	case *ImgMsg:
		mptr.SetImg()
	}
	fmt.Println("m=", m)
}
