package fizzbuzz

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type StatisticsRepoMock struct {
	StoreCalls []Hit
	GetReturn  struct {
		hit   Hit
		count int
	}
}

func (r *StatisticsRepoMock) Store(hit Hit) error {
	r.StoreCalls = append(r.StoreCalls, hit)

	return nil
}

func (r *StatisticsRepoMock) GetMostUsedWithCount() (Hit, int, error) {
	return r.GetReturn.hit, r.GetReturn.count, nil
}

func TestFizzBuzzHandler(t *testing.T) {
	w := httptest.NewRecorder()
	repo := &StatisticsRepoMock{}

	req, _ := http.NewRequest("POST", "whatever", strings.NewReader(`{"int1": 2, "int2": 3, "limit": 10, "str1": "A", "str2": "B"}`))
	FizzBuzzHandler(repo)(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	expected := "1,A,B,A,5,AB,7,A,B,A"
	if expected != string(body) {
		t.Errorf("unexpected return (expected: %s, got: %s)", expected, string(body))
	}

	if len(repo.StoreCalls) != 1 {
		t.Fatalf("unexpected number of call to statisticsRepository.Store (expected: %d, got: %d)", 1, len(repo.StoreCalls))
	}

	expectedHit := Hit{2, 3, 10, "A", "B"}
	if repo.StoreCalls[0] != expectedHit {
		t.Errorf("unexpectedHit stored (expected: %v, got: %v)", expectedHit, repo.StoreCalls[0])
	}
}

func TestStatsHandler(t *testing.T) {
	w := httptest.NewRecorder()
	repo := &StatisticsRepoMock{
		GetReturn: struct {
			hit   Hit
			count int
		}{hit: Hit{1, 2, 3, "A", "B"}, count: 3},
	}

	req, _ := http.NewRequest("GET", "whatever", nil)
	StatsHandler(repo)(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	expected := `{"int1":1,"int2":2,"limit":3,"str1":"A","str2":"B","count":3}
`
	if expected != string(body) {
		t.Errorf("unexpected return (expected: %s, got: %s)", expected, string(body))
	}
}
