package interfaces

import "fmt"

type Structure struct {
	n1, n2 int
}

type Structure2 struct {
	n1, n2 int
}

type SumInterface interface {
	Sum() int
}

func (s Structure) Sum() int {
	return s.n1 + s.n2
}

func (s *Structure2) Sum() int {
	return s.n1 + s.n2
}

func InterfaceStruct() {

	s1 := Structure{1, 3}
	s2 := Structure2{4, 1}
	var i SumInterface = s1
	fmt.Println(i.Sum())
	i = &s2 // i = s2 не сработает, т.к. s2 реализует интерфейс для указателя
	var ifc SumInterface
	var s Structure
	ifc = s
	if ifc == nil {
		fmt.Println("nil")
	} else {
		fmt.Printf("%v, %T", ifc, ifc) // Выведется именно это, т.к значение nil, но тип не nil
	}

}

func InterfaceAsType() {
	var a interface{} // может быть люым типом, реализующим 0 методов
	a = "John Jakor"
	fmt.Printf("%v, %T\n", a, a)
	a = 43
	fmt.Printf("%v, %T\n", a, a)

	// type assertions

	s, ok := a.(string)
	fmt.Println(s, ok)
	// s, ok = a.(int) Ошибка, т.к. нельзя делать приведение типов для интерфейсa уже есть тип,
	// однако можно менять значение и тип данных для переменной a

	switch a.(type) {
	case int:
		fmt.Println("a is int")
	case string:
		fmt.Println("a is string")
	case bool:
		fmt.Println("a is bool")
	default:
		fmt.Printf("unknown type %T\n", a)
	}

	json := make(map[string]interface{})
	json["name"] = "Alex"
	json["age"] = 19
	fmt.Println(json["name"], json["age"])
}
