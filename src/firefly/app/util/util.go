package util

type testString func(s string) bool

func IsIn(list []string) testString {
	return func(s string) bool {
		for _, item := range list {
			if item == s {
				return true
			}
		}
		return false
	}
}

func IsNotIn(list []string) testString {
	return func(s string) bool {
		for _, item := range list {
			if item == s {
				return false
			}
		}
		return true
	}
}

func Filter(source []string, f testString) []string {
	var result []string
	for _, item := range source {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}
