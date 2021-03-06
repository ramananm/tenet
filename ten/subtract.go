package ten

func Subtract(a, b Tensor) (c Tensor) {
	c = NewLike(a)

	for i := range c.Data {
		c.Data[i] = a.Data[i] - b.Data[i]
	}

	return
}

func SubtractScalar(a Tensor, b float64) (c Tensor) {
	c = NewLike(a)

	for i := range c.Data {
		c.Data[i] = a.Data[i] - b
	}

	return
}

func DualSubtract(a, b Tensor) (c Tensor) {
	c = Subtract(a, b)
	return
}

func DualSubtractScalar(a Tensor, bReal, bDual float64) (c Tensor) {
	c = NewLike(a)

	c.AssignReal(SubtractScalar(a.Real(), bReal))
	c.AssignDual(SubtractScalar(a.Dual(), bDual))

	return
}

func SubtractGradient(dc Tensor) (da, db Tensor) {
	da, db = NewLike(dc), NewLike(dc)

	for i := range dc.Data {
		da.Data[i] = dc.Data[i]
		db.Data[i] = -dc.Data[i]
	}

	return
}
