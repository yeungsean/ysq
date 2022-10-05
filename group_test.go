package ysq

import (
	"reflect"
	"testing"
)

type testPet struct {
	ID   int
	Name string
	Age  float32
}

type testPetList []testPet

func testGetPetList() []testPet {
	petCat := testPet{
		ID:   1,
		Name: "Cat",
		Age:  8.9,
	}
	petDog := testPet{
		ID:   2,
		Name: "Dog",
		Age:  5.3,
	}
	petFox := testPet{
		ID:   3,
		Name: "Fox",
		Age:  4,
	}
	petNiki := testPet{
		ID:   4,
		Name: "Niki",
		Age:  8.2,
	}
	petHuanhuan := testPet{
		ID:   5,
		Name: "Huanhuan",
		Age:  5,
	}
	return testPetList{petCat, petDog, petFox, petNiki, petHuanhuan}
}

func testComparePet(prev, current testPet) int {
	if prev.Age > current.Age {
		return -1
	} else if prev.Age < current.Age {
		return 1
	}
	return 0
}

func TestGroupByWithC(t *testing.T) {
	petList := testGetPetList()
	pSlice := FromSlice(petList)
	q := GroupByWithC(pSlice,
		func(p testPet) int { return int(p.Age) },
		func(age int, pets []testPet) map[string]interface{} {
			return map[string]interface{}{
				"key":   age,
				"count": len(pets),
				"min":   FromSlice(pets).Min(testComparePet).Name,
				"max":   FromSlice(pets).Max(testComparePet).Name,
			}
		},
	)
	want8 := KeyValuePair[int, map[string]interface{}]{
		Key: 8,
		Value: map[string]interface{}{
			"key":   8,
			"count": 2,
			"min":   "Niki",
			"max":   "Cat",
		},
	}
	want5 := KeyValuePair[int, map[string]interface{}]{
		Key: 5,
		Value: map[string]interface{}{
			"key":   5,
			"count": 2,
			"min":   "Huanhuan",
			"max":   "Dog",
		},
	}
	eq := true
	for k, v := range q.ToSlice() {
		switch k {
		case 8:
			if !reflect.DeepEqual(v, want8) {
				eq = false
			}
		case 5:
			if !reflect.DeepEqual(v, want5) {
				eq = false
			}
		}
	}

	if !eq {
		t.Errorf("From(%v).GroupBy()=%v", petList, q.ToSlice())
	}
}

func TestPetGroupBy(t *testing.T) {
	petList := testGetPetList()
	pSlice := FromSlice(petList)
	q := GroupBy(pSlice, func(p testPet) int { return int(p.Age) })

	want8 := []testPet{petList[0], petList[3]}
	want5 := []testPet{petList[1], petList[4]}
	want4 := []testPet{petList[2]}

	next := q.Next()
	eq := true
	for item, ok := next(); ok; item, ok = next() {
		switch group := item; group.Key {
		case 8:
			if !reflect.DeepEqual(group.List, want8) {
				eq = false
			}
		case 4:
			if !reflect.DeepEqual(group.List, want4) {
				eq = false
			}
		case 5:
			if !reflect.DeepEqual(group.List, want5) {
				eq = false
			}
		default:
			eq = false
		}
	}

	if !eq {
		t.Errorf("From(%v).GroupBy()=%v", petList, q.ToSlice())
	}
}

func TestEvenOddGroupBy(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	wantEven := []int{2, 4, 6, 8}
	wantOdd := []int{1, 3, 5, 7, 9}

	intSlice := FromSlice(input)
	q := GroupBy(intSlice, func(i int) int { return i % 2 })

	next := q.Next()
	eq := true
	for item, ok := next(); ok; item, ok = next() {
		switch group := item; group.Key {
		case 0:
			if !reflect.DeepEqual(group.List, wantEven) {
				eq = false
			}
		case 1:
			if !reflect.DeepEqual(group.List, wantOdd) {
				eq = false
			}
		default:
			eq = false
		}
	}

	if !eq {
		t.Errorf("From(%v).GroupBy()=%v", input, q.ToSlice())
	}
}
