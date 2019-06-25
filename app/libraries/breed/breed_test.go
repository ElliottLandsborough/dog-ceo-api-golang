package lib

import (
	"reflect"
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
	result := GetRandomItemFromSliceString(s)
	if result != "string1" {
		t.Errorf("Incorrect, got: %s, want: %s.", result, "string1")
	}
}

func TestListAllBreeds(t *testing.T) {
	breeds := []string{"affenpinscher", "spaniel", "spaniel-cocker"}
	got := ListAllBreeds(breeds)
	expected := map[string][]string{
		"affenpinscher": []string{},
		"spaniel":       []string{"cocker"},
	}
	if reflect.DeepEqual(got, expected) != true {
		t.Errorf("Incorrect, got: %s, want: %s.", got, expected)
	}
}

func TestListMasterBreeds(t *testing.T) {
	breeds := []string{"affenpinscher", "spaniel", "spaniel-cocker"}
	got := ListMasterBreeds(breeds)
	expected := []string{"affenpinscher", "spaniel"}
	if stringSlicesAreEqual(got, expected) == false {
		t.Errorf("Incorrect, got: %s, want: %s.", got, expected)
	}
}

func TestListSubBreeds(t *testing.T) {
	breeds := []string{"affenpinscher", "spaniel", "spaniel-cocker"}
	got := ListSubBreeds("spaniel", breeds)
	expected := []string{"cocker"}
	if stringSlicesAreEqual(got, expected) == false {
		t.Errorf("Incorrect, got: %s, want: %s.", got, expected)
	}
}

func TestStringSlicesAreEqual(t *testing.T) {
	s1 := []string{"string1", "string2", "string3", "string4", "string5"}
	s2 := []string{"string1", "string2", "string3", "string4", "string5"}
	s3 := []string{"string1"}
	s4 := []string{"1", "2", "3", "4", "5"}
	one := stringSlicesAreEqual(s1, s2)
	two := stringSlicesAreEqual(s1, s3)
	three := stringSlicesAreEqual(s1, s4)

	if one == false {
		t.Errorf("Incorrect, got: %s, want: %s.", "false", "true")
	}

	if two == true {
		t.Errorf("Incorrect, got: %s, want: %s.", "true", "false")
	}

	if three == true {
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

	result2 := len(getMultipleRandomItemsFromSliceString(s, 9999))
	if result2 != 5 {
		t.Errorf("Incorrect, got: %d, want: %d.", result, 5)
	}
}

func TestListBreedImageRandom(t *testing.T) {
	images := []string{"image1.jpg"}

	if stringSlicesAreEqual(ListBreedImageRandom(images), images) == false {
		t.Errorf("Incorrect, got: %s, want: %s.", "false", "true")
	}
}

func TestListAnyBreedMultiImageRandom(t *testing.T) {
	images := []string{"image1.jpg", "image1.jpg", "image1.jpg"}
	result := ListAnyBreedMultiImageRandom(images, "2")
	expected := []string{"image1.jpg", "image1.jpg"}

	if stringSlicesAreEqual(result, expected) == false {
		t.Errorf("Incorrect, got: %s, want: %s.", "false", "true")
	}

	// biggest 64bit integer is 9223372036854775807 (add 1 to get 9223372036854775808)
	result2 := ListAnyBreedMultiImageRandom(images, "9223372036854775808")

	if len(result2) != 1 {
		t.Errorf("Incorrect, got: %d, want: %d.", len(result2), 1)
	}
}

func TestGenerateBreedYamlKey(t *testing.T) {
	str := GenerateBreedYamlKey("testbreedname")
	good := "breed-info/testbreedname.yaml"
	if str != good {
		t.Errorf("Incorrect, got: %s, want: %s.", str, good)
	}
}

func TestParseYamlToJSON(t *testing.T) {
	yaml := `item:
  - subkey1: "string1"
  - subkey2: "string2"`
	result := ParseYamlToJSON(yaml)
	expected := `{"item":[{"subkey1":"string1"},{"subkey2":"string2"}]}`
	if result != expected {
		t.Errorf("Incorrect, got: %s, want: %s.", result, expected)
	}

	invalidYAML := "!&^%#-"
	result2 := ParseYamlToJSON(invalidYAML)
	expected2 := `{"error":"yaml: did not find expected whitespace or line break"}`
	if result2 != expected2 {
		t.Errorf("Incorrect, got: %s, want: %s.", result2, expected2)
	}
}
