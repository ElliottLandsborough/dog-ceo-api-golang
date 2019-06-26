package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceContainsString(t *testing.T) {
	s := []string{"string1", "string2"}
	result := sliceContainsString(s, "string1")

	assert.Equal(t, true, result)
}

func TestGetRandomItemFromSliceString(t *testing.T) {
	s := []string{"string1"}
	result := GetRandomItemFromSliceString(s)

	assert.Equal(t, "string1", result)
}

func TestListAllBreeds(t *testing.T) {
	breeds := []string{"affenpinscher", "spaniel", "spaniel-cocker"}
	got := ListAllBreeds(breeds)
	expected := map[string][]string{
		"affenpinscher": []string{},
		"spaniel":       []string{"cocker"},
	}

	assert.Equal(t, expected, got)
}

func TestListMasterBreeds(t *testing.T) {
	breeds := []string{"affenpinscher", "spaniel", "spaniel-cocker"}
	got := ListMasterBreeds(breeds)
	expected := []string{"affenpinscher", "spaniel"}

	assert.Equal(t, expected, got)
}

func TestListSubBreeds(t *testing.T) {
	breeds := []string{"affenpinscher", "spaniel", "spaniel-cocker"}
	got := ListSubBreeds("spaniel", breeds)
	expected := []string{"cocker"}

	assert.Equal(t, expected, got)
}

func TestStringSlicesAreEqual(t *testing.T) {
	s1 := []string{"string1", "string2", "string3", "string4", "string5"}
	s2 := []string{"string1", "string2", "string3", "string4", "string5"}
	s3 := []string{"string1"}
	s4 := []string{"1", "2", "3", "4", "5"}
	one := stringSlicesAreEqual(s1, s2)
	two := stringSlicesAreEqual(s1, s3)
	three := stringSlicesAreEqual(s1, s4)

	assert.Equal(t, true, one)
	assert.Equal(t, false, two)
	assert.Equal(t, false, three)
}

func TestShuffleSlice(t *testing.T) {
	s := []string{"string1", "string2", "string3", "string4", "string5"}
	result := shuffleSlice(s)
	assert.Equal(t, s, result)
}

func TestGetMultipleRandomItemsFromSliceString(t *testing.T) {
	s := []string{"string1", "string2", "string3", "string4", "string5"}

	result := len(getMultipleRandomItemsFromSliceString(s, 2))
	assert.Equal(t, 2, result)

	result2 := len(getMultipleRandomItemsFromSliceString(s, 9999))
	assert.Equal(t, 5, result2)
}

func TestListBreedImageRandom(t *testing.T) {
	images := []string{"image1.jpg"}

	assert.Equal(t, images, ListBreedImageRandom(images))
}

func TestListAnyBreedMultiImageRandom(t *testing.T) {
	images := []string{"image1.jpg", "image1.jpg", "image1.jpg"}
	result := ListAnyBreedMultiImageRandom(images, "2")
	expected := []string{"image1.jpg", "image1.jpg"}
	assert.Equal(t, expected, result)

	// biggest 64bit integer is 9223372036854775807 (add 1 to get 9223372036854775808)
	result2 := ListAnyBreedMultiImageRandom(images, "9223372036854775808")
	assert.Equal(t, 1, len(result2))
}

func TestGenerateBreedYamlKey(t *testing.T) {
	str := GenerateBreedYamlKey("testbreedname")
	good := "breed-info/testbreedname.yaml"

	assert.Equal(t, good, str)
}

func TestParseYamlToJSON(t *testing.T) {
	yaml := `item:
  - subkey1: "string1"
  - subkey2: "string2"`
	result := ParseYamlToJSON(yaml)
	expected := `{"item":[{"subkey1":"string1"},{"subkey2":"string2"}]}`

	assert.Equal(t, expected, result)

	invalidYAML := "!&^%#-"
	result2 := ParseYamlToJSON(invalidYAML)
	expected2 := `{"error":"yaml: did not find expected whitespace or line break"}`

	assert.Equal(t, expected2, result2)
}

func TestPrependStringToAllSliceStrings(t *testing.T) {
	images := []string{"image1.jpg", "image2.jpg"}
	prefix := "http://website.com"
	expected := []string{prefix + "image1.jpg", prefix + "image2.jpg"}
	got := PrependStringToAllSliceStrings(images, prefix)

	assert.Equal(t, expected, got)
}

func TestGetBreedFromPathParams(t *testing.T) {
	params1 := map[string]string{"breed1": "spaniel", "breed2": "cocker"}
	params2 := map[string]string{"breed1": "spaniel"}
	result1 := GetBreedFromPathParams(params1)
	result2 := GetBreedFromPathParams(params2)

	assert.Equal(t, result1, "spaniel-cocker")
	assert.Equal(t, result2, "spaniel")
}
