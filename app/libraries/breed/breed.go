package lib

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/ghodss/yaml"
)

// sliceContainsString checks if a string exists in a slice of strings
func sliceContainsString(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// GetRandomItemFromSliceString gets a random item from a slice of strings
func GetRandomItemFromSliceString(slice []string) string {
	// initialize global pseudo random generator
	rand.Seed(time.Now().Unix())

	// pick random string from slice
	return slice[rand.Intn(len(slice))]
}

// ListAllBreeds gets all breeds (master and sub)
func ListAllBreeds(breeds []string) map[string][]string {
	// create map of string arrays
	twoDimensionalArray := map[string][]string{}

	// loop through breeds
	for _, breed := range breeds {
		// explode by -
		exploded := strings.Split(breed, "-")

		// master breed will always be at 0
		master := exploded[0]

		_, ok := twoDimensionalArray[master]

		// master breed isn't in 2d array yet, add it
		if !ok {
			twoDimensionalArray[master] = []string{}
		}

		// sub breed exists?
		if len(exploded) > 1 {
			// sub will always be 1
			sub := exploded[1]

			// append item to slice
			twoDimensionalArray[master] = append(twoDimensionalArray[master], sub)
		}
	}

	return twoDimensionalArray
}

// ListMasterBreeds gets all master breeds
func ListMasterBreeds(breeds []string) []string {
	// slice of strings
	s := []string{}

	// loop through breeds
	for _, breed := range breeds {
		// explode by -
		exploded := strings.Split(breed, "-")

		if !sliceContainsString(s, exploded[0]) {
			// append to breeds
			s = append(s, exploded[0])
		}
	}

	return s
}

// ListSubBreeds gets all sub breeds by master breed name
func ListSubBreeds(breedFromURL string, breeds []string) []string {

	// slice of strings
	s := []string{}

	// loop through breeds
	for _, breed := range breeds {
		// explode by -
		exploded := strings.Split(breed, "-")

		// primary breed will always be there
		primary := exploded[0]

		// does the url segment match this item?
		if breedFromURL == primary {
			// sub breed exists?
			if len(exploded) > 1 {
				// sub will always be 1
				sub := exploded[1]

				// append item to slice
				s = append(s, sub)
			}
		}
	}

	return s
}

// stringSlicesAreEqual tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func stringSlicesAreEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func shuffleSlice(slice []string) []string {
	// initialize global pseudo random generator
	rand.Seed(time.Now().Unix())

	// shuffle the items
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})

	return slice
}

func getMultipleRandomItemsFromSliceString(slice []string, amount int) []string {
	// shuffle the items
	slice = shuffleSlice(slice)

	// dont bother if we want all the items
	if amount > len(slice) {
		return slice
	}

	// return {amount} items from slice
	return slice[0:amount]
}

// ListBreedImageRandom gets a random image from all the master breed images
func ListBreedImageRandom(images []string) []string {
	// pick random image from slice
	image := GetRandomItemFromSliceString(images)

	return []string{image}
}

// ListAnyBreedMultiImageRandom gets all images from a random breed, returns {count} images
func ListAnyBreedMultiImageRandom(slice []string, count string) []string {
	// string to int
	i, err := strconv.Atoi(count)
	if err != nil {
		// handle error
		i = 1
	}

	return getMultipleRandomItemsFromSliceString(slice, i)
}

// GenerateBreedYamlKey generates the breeds yaml key
func GenerateBreedYamlKey(breed string) string {
	return "breed-info/" + breed + ".yaml"
}

// ParseYamlToJSON takes a yaml string, returns a JSON string
func ParseYamlToJSON(yamlString string) string {

	data, err := yaml.YAMLToJSON([]byte(yamlString))
	if err != nil {
		errorMap := map[string]string{
			"error": err.Error(),
		}
		errorJSON, _ := json.Marshal(errorMap)
		return string(errorJSON)
	}

	return string(data)
}
