package container_with_most_water

import (
	. "github.com/onsi/gomega"
	"testing"
)

func Test_maxArea(t *testing.T) {
	testCases := []struct {
		name   string
		params []int
		res    int
	}{
		{
			name:   "case 1: max result is 49",
			params: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			res:    49,
		},
	}
	g := NewGomegaWithT(t)
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res := maxArea(tt.params)
			g.Expect(res).To(Equal(tt.res))
		})
	}
}
