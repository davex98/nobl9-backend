package infrastracture_test

import (
	"github.com/davex98/nobl9-backend/random-generator/infrastracture"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestSliceConversion(t *testing.T) {
	tests := []struct {
		input []string
		want  []int
		err   string
	}{
		{input: []string{"1", "2", "3"}, want: []int{1, 2, 3}, err: ""},
		{input: []string{"-1", "2"}, want: []int{-1, 2}, err: ""},
		{input: []string{""}, want: nil, err: ""},
		{input: []string{"a"}, want: nil, err: "string instead of number"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run("Slice conversion test", func(t *testing.T) {
			t.Parallel()
			numbers, err := infrastracture.StringSliceToNumbers(tc.input)
			if tc.err != "" {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			if !reflect.DeepEqual(numbers, tc.want) {
				t.Errorf("expected %v, got %v", tc.want, numbers)
			}
		})
	}
}
