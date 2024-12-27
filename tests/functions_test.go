package tests

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	// Функция, выполняемая для всех тестов одного пакета
	fmt.Println("setup")
	res := m.Run()
	fmt.Printf("res = %v\n tear-down\n", res)
}
func TestMultiple(t *testing.T) {
	// Группы выполняются последовательно, а тесты в группах параллельно
	t.Run("simple", func(t *testing.T) {
		x, y, expectedResult := 3, 5, 15
		realResult := Multiple(x, y)
		if expectedResult != realResult {
			t.Errorf("Expected res %v != %v Real res", expectedResult, realResult)
		}
		t.Run("1", func(t *testing.T) {
			x, y, expectedResult := 1, 1, 1
			realResult := Multiple(x, y)
			if expectedResult != realResult {
				t.Errorf("Expected res %v != %v Real res", expectedResult, realResult)
			}
		})
	})
	t.Run("medium", func(t *testing.T) {
		x, y, expectedResult := 22, 5, 110
		realResult := Multiple(x, y)
		if expectedResult != realResult {
			t.Errorf("Expected res %v != %v Real res", expectedResult, realResult)
		}
	})
	t.Run("hard", func(t *testing.T) {
		x, y, expectedResult := 2223, 5121, 11383983
		realResult := Multiple(x, y)
		if expectedResult != realResult {
			t.Errorf("Expected res %v != %v Real res", expectedResult, realResult)
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("groupA", func(t *testing.T) {
		t.Log("GroupA")
		t.Run("simple", func(t *testing.T) {
			t.Parallel()
			t.Log("simple")
			x, y, expectedResult := 3, 5, 15
			realResult := Add(x, y)
			if expectedResult != realResult {
				t.Errorf("Expected res %v != %v Real res", expectedResult, realResult)
			}
			t.Run("1", func(t *testing.T) {
				x, y, expectedResult := 1, 1, 1
				realResult := Add(x, y)
				if expectedResult != realResult {
					t.Errorf("Expected res %v != %v Real res", expectedResult, realResult)
				}
			})
		})
		t.Run("medium", func(t *testing.T) {
			t.Parallel()
			t.Log("medium")
			x, y, expectedResult := 22, 5, 110
			realResult := Add(x, y)
			if expectedResult != realResult {
				t.Errorf("Expected res %v != %v Real res", expectedResult, realResult)
			}
		})
		t.Run("hard", func(t *testing.T) {
			t.Parallel()
			t.Log("hard")
			x, y, expectedResult := 2223, 5121, 11383983
			realResult := Add(x, y)
			if expectedResult != realResult {
				t.Errorf("Expected res %v != %v Real res", expectedResult, realResult)
			}
		})
	})
	t.Run("groupB", func(t *testing.T) {
		t.Log("GroupB")
		t.Run("simple", func(t *testing.T) {
			t.Parallel()
			t.Log("simple")
			x, y, expectedResult := 3, 5, 15
			realResult := Add(x, y)
			if expectedResult != realResult {
				t.Errorf("Expected res %v != %v Real res", expectedResult, realResult)
			}
			t.Run("1", func(t *testing.T) {
				x, y, expectedResult := 1, 1, 1
				realResult := Add(x, y)
				if expectedResult != realResult {
					t.Errorf("Expected res %v != %v Real res", expectedResult, realResult)
				}
			})
		})
		t.Run("medium", func(t *testing.T) {
			t.Parallel()
			t.Log("medium")
			x, y, expectedResult := 22, 5, 110
			realResult := Add(x, y)
			if expectedResult != realResult {
				t.Errorf("Expected res %v != %v Real res", expectedResult, realResult)
			}
		})
		t.Run("hard", func(t *testing.T) {
			t.Parallel()
			t.Log("hard")
			x, y, expectedResult := 2223, 5121, 11383983
			realResult := Add(x, y)
			if expectedResult != realResult {
				t.Errorf("Expected res %v != %v Real res", expectedResult, realResult)
			}
		})
	})
}

// Тесты собираются отдельно от основного кода, поэтому для проверки того собираются ли тесты нужно использовать go test
