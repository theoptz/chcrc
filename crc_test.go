package chcrc

import "testing"

func TestGetChCrc64(t *testing.T) {
	testCases := []struct {
		val string
		res uint64
	}{
		{
			val: "example",
			res: 3321340006651021620,
		},
		{
			val: "Hello, playground",
			res: 4363968896809504055,
		},
		{
			val: "16DE44E64A38E6D1991A10AC8E2F3BE9-79C1B6E52285CA563EEFB931DE4D2F88A2A40A95",
			res: 9358187628843672334,
		},
	}

	for _, tc := range testCases {
		res := GetCrc64(tc.val)

		if tc.res != res {
			t.Errorf("Res should be %d but got %d", tc.res, res)
		}
	}

}
