package fizzbuzz

import "strconv"

func FizzBuzz(fizzMultiple int, buzzMultiple int, limit int, fizzStr string, buzzStr string) []string {
	list := make([]string, limit)
	for i := 1; i <= limit; i++ {
		if i%fizzMultiple == 0 && i%buzzMultiple == 0 {
			list[i-1] = fizzStr + buzzStr
			continue
		}
		if i%fizzMultiple == 0 {
			list[i-1] = fizzStr
			continue
		}
		if i%buzzMultiple == 0 {
			list[i-1] = buzzStr
			continue
		}

		list[i-1] = strconv.Itoa(i)
	}

	return list
}
