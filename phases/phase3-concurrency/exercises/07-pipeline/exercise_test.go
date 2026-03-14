package exercise_07

import (
	"reflect"
	"testing"
)

func collect(ch <-chan int) []int {
	var result []int
	for v := range ch {
		result = append(result, v)
	}
	return result
}

func TestGenerate(t *testing.T) {
	got := collect(Generate(1, 5))
	want := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Generate(1,5) = %v, want %v", got, want)
	}
}

func TestFilter(t *testing.T) {
	in := Generate(1, 10)
	evens := Filter(in, func(v int) bool { return v%2 == 0 })
	got := collect(evens)
	want := []int{2, 4, 6, 8, 10}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Filter(evens) = %v, want %v", got, want)
	}
}

func TestMap(t *testing.T) {
	in := Generate(1, 5)
	doubled := Map(in, func(v int) int { return v * 2 })
	got := collect(doubled)
	want := []int{2, 4, 6, 8, 10}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Map(double) = %v, want %v", got, want)
	}
}

func TestReduce(t *testing.T) {
	sum := Reduce(Generate(1, 5), 0, func(acc, v int) int { return acc + v })
	if sum != 15 {
		t.Errorf("Reduce(sum 1..5) = %d, want 15", sum)
	}
}

func TestPipeline(t *testing.T) {
	// Generate 1..20 → filter evens → map *3 → sum
	result := Reduce(
		Map(
			Filter(Generate(1, 20), func(v int) bool { return v%2 == 0 }),
			func(v int) int { return v * 3 },
		),
		0,
		func(acc, v int) int { return acc + v },
	)
	// evens: 2,4,6,8,10,12,14,16,18,20 sum=110; *3=330
	if result != 330 {
		t.Errorf("pipeline result = %d, want 330", result)
	}
}
