package main

func main() {
	m1 := make(map[string]any)
	m1["add"] = add
	m1["sub"] = sub

	m1["mul"] = multiply
	m1["div"] = divison
	m1["greet"] = greet
	for i, v := range m1 {
		switch vt := v.(type) {
		case func(int, int) int:
			switch i {
			case "add":
				println(vt(10, 20))
			case "sub":
				println(vt(10, 20))
			case "mul":
				println(vt(10, 20))
			case "div":
				println(vt(20, 20))

			}
		case func() string:
			println(vt())
		}
	}

}
func add(no1 int, no2 int) int {
	return no1 + no2

}

func sub(no1 int, no2 int) int {
	return no1 - no2

}

func multiply(no1 int, no2 int) int {
	return no1 * no2

}

func divison(no1 int, no2 int) int {
	return no1 / no2

}
func greet() string {
	return "hello world"
}
