package api

import "github.com/Rhaqim/hexarch/internal/ports"

type Adpater struct {
	db    ports.DBPort
	Arith ports.ArthmeticPort
}

func NewAdapter(db ports.DBPort, arith ports.ArthmeticPort) *Adpater {
	return &Adpater{
		db:    db,
		Arith: arith,
	}
}

func (apia Adpater) GetAddition(a, b int) (int32, error) {
	result, err := apia.Arith.Add(a, b)

	if err != nil {
		return 0, err
	}

	apia.db.AddtoHistory(result, "Addition")

	return result, nil
}

func (apia Adpater) GetSubtraction(a, b int) (int32, error) {
	result, err := apia.Arith.Sub(a, b)

	if err != nil {
		return 0, err
	}

	apia.db.AddtoHistory(result, "Subtraction")

	return result, nil
}

func (apia Adpater) GetMultiplication(a, b int) (int32, error) {
	result, err := apia.Arith.Mul(a, b)

	if err != nil {
		return 0, err
	}

	apia.db.AddtoHistory(result, "Multiplication")

	return result, nil
}

func (apia Adpater) GetDivision(a, b int) (int32, error) {
	result, err := apia.Arith.Div(a, b)

	if err != nil {
		return 0, err
	}

	apia.db.AddtoHistory(result, "Division")

	return result, nil
}
