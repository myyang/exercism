/*
Package leap provides function to check leap

Here's infomation about leap year:

The tricky thing here is that a leap year in the Gregorian calendar occurs:

```plain
on every year that is evenly divisible by 4
  except every year that is evenly divisible by 100
    unless the year is also evenly divisible by 400
```

*/

package leap

const testVersion = 2

// IsLeapYear returns a bool to indicate year is leap or not
func IsLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
