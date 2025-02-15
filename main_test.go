package main

import (
	"math"
	"sort"
	"testing"
)

// Тест для кубического с одним действительным корнем
func TestCubic_OneRealRoot(t *testing.T) {
	a, b, c, d := 1.0, 0.0, 0.0, -1.0
	expected := []float64{1.0}

	roots, count := cubicRoots(a, b, c, d)

	if count != len(expected) {
		t.Errorf("Количество корней: %v; Ожидается %v корень", count, len(expected))
	}

	if !floatSlicesEqual(roots, expected) {
		t.Errorf("Корни: %v; Ожидаемые корни: %v", roots, expected)
	}
}

func TestCubic_TwoRealRoots(t *testing.T) {
	a, b, c, d := 0.0, 1.0, -5.0, 6.0
	expected := []float64{2.0, 3.0}

	roots, count := cubicRoots(a, b, c, d)

	if count != len(expected) {
		t.Errorf("Количество корней: %v; Ожидается %v корня", count, len(expected))
	}

	if !floatSlicesEqual(roots, expected) {
		t.Errorf("Корни: %v; Ожидаемые корни: %v", roots, expected)
	}
}

func TestCubic_ThreeRealRoots(t *testing.T) {
	a, b, c, d := 1.0, -6.0, 11.0, -6.0
	expected := []float64{1.0, 2.0, 3.0}

	roots, count := cubicRoots(a, b, c, d)

	if count != len(expected) {
		t.Errorf("Количество корней: %v; Ожидается %v корня", count, len(expected))
	}

	if !floatSlicesEqual(roots, expected) {
		t.Errorf("Корни: %v; Ожидаемые корни: %v", roots, expected)
	}
}

func TestCubic_NoRoots(t *testing.T) {
	a, b, c, d := 0.0, 1.0, 2.0, 5.0
	expected := 0

	_, count := cubicRoots(a, b, c, d)

	if count != expected {
		t.Errorf("Ожидалось, что действительных корней нет, однако найдено %v корней(-я)", count)
	}
}

func TestCubic_InfiniteRoots(t *testing.T) {
	a, b, c, d := 0.0, 0.0, 0.0, 0.0
	expected := -1

	_, count := cubicRoots(a, b, c, d)

	if count != expected {
		t.Errorf("Ожидалось бесконечное количество корней, однако что-то пошло не так")
	}
}

func floatSlicesEqual(a, b []float64) bool {
	sort.Float64s(a)
	sort.Float64s(b)

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if math.Abs(a[i]-b[i]) > 1e-12 {
			return false
		}
	}

	return true
}
