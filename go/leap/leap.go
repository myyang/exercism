/*
Package leap provides a function to check leap year

What's leap year? Please see: https://en.wikipedia.org/wiki/Leap_year .

And the tricky thing here is that a leap year in the Gregorian calendar occurs:
	on every year that is evenly divisible by 4
	except every year that is evenly divisible by 100
	unless the year is also evenly divisible by 400

*/
package leap

const testVersion = 2

// IsLeapYear returns a bool to indicate year is leap or not.
// Ex:
//     fmt.Print(IsLeapYear(1900)) # false
//     fmt.Print(IsLeapYear(2000)) # true
//     fmt.Print(IsLeapYear(2004)) # true
//
// Currently,testing positive year value only.
func IsLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
