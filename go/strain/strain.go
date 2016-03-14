package strain

// Ints is alias of int array
type Ints []int

// Lists is alias of array of int array
type Lists [][]int

// Strings is alias of string array
type Strings []string

// Keep content by given function
func (ints Ints) Keep(f func(int) bool) Ints {
	if len(ints) == 0 {
		return nil
	}

	r := Ints{}
	for _, v := range ints {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// Discard content by given function
func (ints Ints) Discard(f func(int) bool) Ints {
	if len(ints) == 0 {
		return nil
	}

	r := Ints{}
	for _, v := range ints {
		if f(v) {
			continue
		}
		r = append(r, v)
	}
	return r
}

// Keep content by given function
func (lists Lists) Keep(f func([]int) bool) Lists {
	if len(lists) == 0 {
		return nil
	}

	r := Lists{}
	for _, v := range lists {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// Keep content by given function
func (strings Strings) Keep(f func(string) bool) Strings {
	if len(strings) == 0 {
		return nil
	}

	r := Strings{}
	for _, v := range strings {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}
