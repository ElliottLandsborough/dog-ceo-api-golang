package lib

import (
	"testing"
)

func TestSliceContainsString(t *testing.T) {
	s := []string{"string1", "string2"}
	result := sliceContainsString(s, "string1")
	if result == false {
		t.Errorf("Incorrect, got: %s, want: %s.", "false", "true")
	}
}

func TestGetRandomItemFromSliceString(t *testing.T) {
	s := []string{"string1"}
	result := getRandomItemFromSliceString(s)
	if result != "string1" {
		t.Errorf("Incorrect, got: %s, want: %s.", result, "string1")
	}
}

func TestStringSlicesAreEqual(t *testing.T) {
	s1 := []string{"string1", "string2", "string3", "string4", "string5"}
	s2 := []string{"string1", "string2", "string3", "string4", "string5"}
	s3 := []string{"string1"}
	one := stringSlicesAreEqual(s1, s2)
	two := stringSlicesAreEqual(s1, s3)

	if one == false {
		t.Errorf("Incorrect, got: %s, want: %s.", "false", "true")
	}

	if two == true {
		t.Errorf("Incorrect, got: %s, want: %s.", "true", "false")
	}
}

func TestShuffleSlice(t *testing.T) {
	s := []string{"string1", "string2", "string3", "string4", "string5"}
	result := shuffleSlice(s)
	if stringSlicesAreEqual(s, result) == false {
		t.Errorf("Incorrect, got: %s, want: %s.", "true", "false")
	}
}

func TestGetMultipleRandomItemsFromSliceString(t *testing.T) {
	s := []string{"string1", "string2", "string3", "string4", "string5"}
	result := len(getMultipleRandomItemsFromSliceString(s, 2))
	if result != 2 {
		t.Errorf("Incorrect, got: %d, want: %d.", result, 2)
	}
}

func TestGenerateBreedYamlKey(t *testing.T) {
	str := generateBreedYamlKey("testbreedname")
	good := "breed-info/testbreedname.yaml"
	if str != good {
		t.Errorf("Incorrect, got: %s, want: %s.", str, good)
	}
}
