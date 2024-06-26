package assert_test

type person struct {
	Name    string
	Age     int
	Hobbies []hobby
}

type hobby struct {
	Name     string
	Category string
}
