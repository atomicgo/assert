package assert_test

type person struct {
	Name    string
	Hobbies []hobby
	Age     int
}

type hobby struct {
	Name     string
	Category string
}
