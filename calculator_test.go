package calculator

import "testing"

func TestShouldSum2And3Result5(t *testing.T) {
	//Arrange
	expected := 5.0

	//Act
	result := Sum(2.0, 3.0)

	//Assert
	if result != expected {
		t.Error("Expected:", expected, "Result:", result)
	}
}

func TestShouldSubtract3And6ResultNegative3(t *testing.T) {
	//Arrange
	expected := -3.0
	//Act
	result := Subtract(3, 6)

	//Assert
	if result != expected {
		t.Error("Expected:", expected, "Result:", result)
	}
}

func TestShouldMultiply4And5Result20(t *testing.T) {
	//Arrange
	expected := 20.0

	//Act
	result := Multiply(4.0, 5.0)

	//Assert
	if expected != result {
		t.Error("Expected:", expected, "Result:", result)
	}
}

func TestShouldDivide4And2Result2(t *testing.T) {
	//Arrange
	expected := 2.0

	//Act
	result := Divide(4.0, 2.0)

	//Expected
	if expected != result {
		t.Error("Expected:", expected, "Result:", result)
	}
}
