package fizzbuzz

import "sync"

type Hit struct {
	fizzMultiple int
	buzzMultiple int
	limit        int
	fizzStr      string
	buzzStr      string
}

type StatisticsRepository interface {
	Store(hit Hit) error
	GetMostUsedWithCount() (Hit, int, error)
}

type memoryStatisticsRepository struct {
	hits          map[Hit]int
	mu            sync.RWMutex
	mostUsed      Hit
	countMostUsed int
}

func NewMemoryStatisticsRepository() StatisticsRepository {
	return &memoryStatisticsRepository{hits: make(map[Hit]int)}
}

func (r *memoryStatisticsRepository) Store(hit Hit) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.hits[hit]
	if !ok {
		r.hits[hit] = 1
	} else {
		r.hits[hit]++
	}

	if r.hits[hit] > r.countMostUsed {
		r.mostUsed = hit
		r.countMostUsed = r.hits[hit]
	}

	return nil
}

func (r *memoryStatisticsRepository) GetMostUsedWithCount() (Hit, int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.mostUsed, r.countMostUsed, nil
}
