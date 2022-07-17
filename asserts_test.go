package main

import "testing"

func TestIsSubstring_SameThing_ShouldBeTrue(t *testing.T) {
	result := IsSubstring("hi", "hi")

	if result != true {
		t.Error("result should be true, but got ", result)
	}
}

func TestIsSubstring_SubstringAfterStringStart_ShouldBeTrue(t *testing.T) {
	result := IsSubstring("abasadfsehi", "hi")

	if result != true {
		t.Error("result should be true, but got ", result)
	}
}

func TestIsSubstring_FirstLetterAppears_ShouldBeTrue(t *testing.T) {
	result := IsSubstring("abasadhsehi", "hi")

	if result != true {
		t.Error("result should be true, but got ", result)
	}
}

func TestIsSubstring_DoesNotAppear_ShouldBeFalse(t *testing.T) {
	result := IsSubstring("abasadhsehd", "hi")

	if result != false {
		t.Error("result should be false, but got ", result)
	}
}

func TestRemoveEmpty_GivenEmptyStrings_Removes(t *testing.T) {
	testArray := []string{"ABC", " ", "DEF", ""}

	result := RemoveEmpty(testArray)

	if len(result) != 2 {
		t.Error("expected to remove the two blank items, result: ", result)
	}

	if result[0] != "ABC" {
		t.Error("Expected ABC, but got ", result)
	}

	if result[1] != "DEF" {
		t.Error("Expected DEF, but got ", result)
	}
}

func TestIndexOf_FirstElement_ReturnsZero(t *testing.T) {
	testArray := []string{"ABC", " ", "DEF", ""}

	result := IndexOf(testArray, "ABC")

	if result != 0 {
		t.Error("Expected 0 index, but got ", result)
	}
}

func TestIndexOf_ThirdElement_ReturnsTwo(t *testing.T) {
	testArray := []string{"ABC", " ", "DEF", ""}

	result := IndexOf(testArray, "DEF")

	if result != 2 {
		t.Error("Expected 2 index, but got ", result)
	}
}

func TestIndexOf_NotPresent_ReturnsNegOne(t *testing.T) {
	testArray := []string{"ABC", " ", "DHE"}

	result := IndexOf(testArray, "DEF")

	if result != -1 {
		t.Error("Expected -1 index, but got ", result)
	}
}

func TestContainsString_NotPresent_ReturnsFalse(t *testing.T) {
	testArray := "ABC DEF GHI"

	result := ContainsString(testArray, "EFG")

	if result != false {
		t.Error("Expected false, but got ", result)
	}
}

func TestContainsString_Present_ReturnsTrue(t *testing.T) {
	testArray := "ABC DEF GHI"

	result := ContainsString(testArray, "DEF")

	if result != true {
		t.Error("Expected true, but got ", result)
	}
}

func TestContainsChar_NotPresent_ReturnsFalse(t *testing.T) {
	testArray := "ABC DEF GHI"

	result := ContainsChar(testArray, 'Z')

	if result != false {
		t.Error("Expected false, but got ", result)
	}
}

func TestContainsChar_Present_ReturnsTrue(t *testing.T) {
	testLine := "ABC DEF GHI"

	result := ContainsChar(testLine, 'E')

	if result != true {
		t.Error("Expected true, but got ", result)
	}
}

func TestArrContainsChar_NotPresent_ReturnsFalse(t *testing.T) {
	testArray := []string{"ABC", " ", "DEF", ""}

	result := ArrContainsChar(testArray, 'Z')

	if result != false {
		t.Error("Expected false, but got ", result)
	}
}

func TestArrContainsChar_Present_ReturnsTrue(t *testing.T) {
	testArray := []string{"ABC", " ", "DEF", ""}

	result := ArrContainsChar(testArray, 'E')

	if result != true {
		t.Error("Expected true, but got ", result)
	}
}
