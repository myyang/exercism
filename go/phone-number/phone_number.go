// Package phonenumber provides functions to validate phone number
package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Number validate phone number
func Number(phoneNum string) (string, error) {
	re, err := regexp.Compile(`[[:digit:]]*`)
	if err != nil {
		return "", errors.New("error regexp")
	}
	nums := strings.Join(re.FindAllString(phoneNum, -1), "")

	if len(nums) < 10 || len(nums) > 11 || (len(nums) == 11 && nums[:1] != "1") {
		return "", errors.New("Error digits")
	} else if len(nums) == 11 {
		return nums[1:], nil
	}
	return nums, nil
}

// AreaCode extract code from number
func AreaCode(phoneNum string) (string, error) {
	r, err := Number(phoneNum)
	if err != nil {
		return "", err
	}
	return r[:3], nil
}

// Format the phone number with: (XXX) XXX-XXXX
func Format(phoneNum string) (string, error) {
	r, err := Number(phoneNum)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%v) %v-%v", r[:3], r[3:6], r[6:]), nil
}
