package ports

type APIPort interface {
	GetAddition(a, b int) (int32, error)
	GetSubtraction(a, b int) (int32, error)
	GetMultiplication(a, b int) (int32, error)
	GetDivision(a, b int) (int32, error)
}
