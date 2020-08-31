package modle

type student struct {
	Name string
	Age  int
}

func NewStudent(n string, a int) student {
	return student{n, a}
}
