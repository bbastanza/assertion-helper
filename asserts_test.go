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

// func TestIsSubstring(t *testing.T) {
// 	result := IsSubstring("abchabdahisx", "hi")
//
// 	if result != true {
// 		t.Error("result should be true, but got ", result)
// 	}
//
// }
