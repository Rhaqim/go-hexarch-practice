package arithmetic

type Adpater struct{}

func NewArithmetic() *Adpater {
	return &Adpater{}
}

func (arith Adpater) Add(a, b int) (int32, error) {
	return int32(a + b), nil
}

func (arith Adpater) Sub(a, b int) (int32, error) {
	return int32(a - b), nil
}

func (arith Adpater) Mul(a, b int) (int32, error) {
	return int32(a * b), nil
}

func (arith Adpater) Div(a, b int) (int32, error) {
	return int32(a / b), nil
}
