// Package brackets provides Bracket to check string contains balanced and valid bracket.
package brackets

const testVersion = 3

var brackMap = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
}

type stackError struct{}

// Error message
func (e stackError) Error() string {
	return "Stack error"
}

type stack []string

func (s *stack) count() int {
	return len(*s)
}

func (s *stack) empty() bool {
	return s.count() == 0
}

func (s *stack) top() (string, error) {
	if s.count() > 0 {
		return (*s)[s.count()-1], nil
	}
	return "", stackError{}
}

func (s *stack) push(new string) {
	*s = append(*s, new)
}

func (s *stack) pop() (string, error) {
	r, err := s.top()
	if err != nil {
		return "", stackError{}
	}
	*s = (*s)[:s.count()-1]
	return r, nil
}

// Bracket input string, return bool indicates wheather string is valid
func Bracket(str string) (bool, error) {
	st := stack{}
	for _, s := range str {
		s := string(s)
		if s == "(" || s == "{" || s == "[" {
			st.push(s)
			continue
		} else if s == ")" || s == "]" || s == "}" {
			r, err := st.pop()
			if err != nil || r != brackMap[s] {
				return false, nil
			}
		}
	}
	return st.empty(), nil
}
