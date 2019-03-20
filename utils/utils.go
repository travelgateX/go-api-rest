package utils

// AppendIfMissing appends string to a slice if missing
func AppendIfMissing(a []string, i string) []string {
	for _, e := range a {
		if e == i {
			return a
		}
	}
	return append(a, i)
}

// AppendUniqueSlices join two slices omitting the duplicity
func AppendUniqueSlices(a, b []string) []string {
	for _, e := range a {
		if !SliceContainsString(e, b) {
			b = append(b, e)
		}
	}
	return b
}

// SliceContainsString check if the slice contains the string
func SliceContainsString(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
