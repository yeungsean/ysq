package ysq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistinctBy(t *testing.T) {
	petList := testGetPetList()
	netPetList := append(petList, testPet{
		ID:   5,
		Name: "Huanhuan",
		Age:  5,
	})
	actual := FromSlice(netPetList).DistinctBy(func(tp testPet) int64 {
		return int64(tp.ID)
	}).ToSlice()
	assert.Equal(t, petList, actual)
}
