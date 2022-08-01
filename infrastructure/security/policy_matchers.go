package security

import "strings"

// Custom matcher where the first argument takes a collection of keys that have the signature
// key1|key2|key3 and the second argument takes a key and if the first argument contains
// the key in the second argument the function returns true. If the value passed is "any"
// the function will always return true.
func Contains(collection string, val string) bool {
	if val == "any" {
		return true
	}

	collectionKeys := strings.Split(collection, "|")

	for _, key := range collectionKeys {
		keyTrimmed := strings.Trim(key, " ")

		if val == keyTrimmed {
			return true
		}
	}

	return false
}

func ContainsFunc(args ...any) (any, error) {
	keys := args[0].(string)
	val := args[1].(string)

	return (bool)(Contains(keys, val)), nil
}
