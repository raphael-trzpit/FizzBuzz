package fizzbuzz

import (
	"fmt"
)

func ExampleNewMemoryStatisticsRepository() {
	repo := NewMemoryStatisticsRepository()
	if err := repo.Store(Hit{1, 2, 3, "str1", "str2"}); err != nil {
		fmt.Println(err.Error())
	}
	if err := repo.Store(Hit{2, 2, 2, "str2", "str3"}); err != nil {
		fmt.Println(err.Error())
	}
	if err := repo.Store(Hit{1, 2, 3, "str1", "str2"}); err != nil {
		fmt.Println(err.Error())
	}

	hit, count, _ := repo.GetMostUsedWithCount()
	fmt.Printf(
		"fizzMultiple: %d, buzzMultiple: %d, limit: %d, fizzStr: %s, buzzStr: %s, count: %d",
		hit.fizzMultiple,
		hit.buzzMultiple,
		hit.limit,
		hit.fizzStr,
		hit.buzzStr,
		count,
	)
	// Output: fizzMultiple: 1, buzzMultiple: 2, limit: 3, fizzStr: str1, buzzStr: str2, count: 2
}
