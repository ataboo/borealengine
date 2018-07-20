package mathies

import (
	"testing"
	"math"
)

func TestVector2_Add(t *testing.T) {
	v1 := Vector2{1, 1}
	v2 := Vector2{2, 2}

	product := v1.Add(v2)
	if product != (Vector2{3, 3}) {
		t.Errorf("Wrong produce %+v", product)
	}

	if v1 != (Vector2{1, 1}) {
		t.Errorf("should be immutable")
	}
}

func TestVector2_Sub(t *testing.T) {
	v1 := Vector2{1, 1}
	v2 := Vector2{2, 2}

	product := v1.Sub(v2)
	if product != (Vector2{-1, -1}) {
		t.Errorf("Wrong product %+v", product)
	}

	if v1 != (Vector2{1, 1}) {
		t.Errorf("should be immutable")
	}
}

func TestVector2_Mul(t *testing.T) {
	v1 := Vector2{2, 2}

	product := v1.Mul(2)
	if product != (Vector2{4, 4}) {
		t.Errorf("Wrong product %+v", product)
	}

	if v1 != (Vector2{2, 2}) {
		t.Errorf("should be immutable")
	}
}

func TestVector2_Dot(t *testing.T) {
	v1 := Vector2{2, 2}
	v2 := Vector2{3, 3}

	product := v1.Dot(v2)

	if product != 12 {
		t.Errorf("wrong product %f", product)
	}

	if v1 != (Vector2{2, 2}) {
		t.Errorf("should be immutable")
	}
}

func TestVector2_Mag(t *testing.T) {
	v1 := Vector2{3, 4}

	product := v1.Mag()

	if product != 5 {
		t.Errorf("wrong magnitude %f", product)
	}

	if v1 != (Vector2{3, 4}) {
		t.Errorf("should be immutable")
	}
}

func TestVector2_Normalized(t *testing.T) {
	v1 := Vector2{1, 1}

	product := v1.Normalized()
	mag := product.Mag()

	if math.Abs(float64(mag - 1)) > 0.0001 || product.X != product.Y {
		t.Errorf("invalid product %+v", product)
	}

	if v1 != (Vector2{1, 1}) {
		t.Errorf("should be immutable")
	}
}

func TestVector2_Project(t *testing.T) {
	v1 := Vector2{3, 3}
	v2 := Vector2{1, 0}
	v3 := Vector2{-1, 0}

	positiveProduct := v1.Project(v2)

	if positiveProduct != 3 {
		t.Errorf("wrong product %f", positiveProduct)
	}

	negativeProduct := v1.Project(v3)
	if negativeProduct != -3 {
		t.Errorf("wrong product %f", negativeProduct)
	}

	if v1 != (Vector2{3,3}) {
		t.Errorf("should be immutable")
	}
}

func TestNormalizeWontNan(t *testing.T) {
	v1 := Vector2{0, 0}
	v2 := Vector2{1, 1}

	norm := v1.Normalized()

	if norm != (Vector2{0, 0}) {
		t.Errorf("should be 0, 0 %+v", norm)
	}

	product := v2.Project(v1)

	if product != 0 {
		t.Errorf("should be 0 %f", product)
	}
}