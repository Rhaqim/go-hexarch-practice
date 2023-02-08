package ports

type ArithmeticPort interface {
	Add(a, b int) (int32, error)
	Sub(a, b int) (int32, error)
	Mul(a, b int) (int32, error)
	Div(a, b int) (int32, error)
}
