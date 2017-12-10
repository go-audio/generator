package euclidean

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRhythm(t *testing.T) {
	type args struct {
		accents    int
		totalSteps int
	}
	tests := []struct {
		args args
		want []bool
	}{
		// test values extracted from the white paper on euclidean rhythms
		// http://cgm.cs.mcgill.ca/~godfried/publications/banff.pdf
		// more accents than steps
		{args{13, 5}, []bool{true, true, true, true, true}},
		// invalid accent number
		{args{-1, 5}, []bool{false, false, false, false, false}},
		// invalid step number
		{args{5, -1}, []bool{}},
		// tricky one
		{args{5, 13}, []bool{true, false, false, true, false, true, false, false, true, false, true, false, false}},
		// basic
		{args{1, 1}, []bool{true}},
		{args{1, 2}, []bool{true, false}},
		{args{1, 3}, []bool{true, false, false}},
		{args{1, 4}, []bool{true, false, false, false}},
		// West African, Latin American
		{args{2, 3}, []bool{true, false, true}},
		// Classical jazz and Persian
		{args{2, 5}, []bool{true, false, true, false, false}},
		// Trinidad and Persian
		{args{3, 4}, []bool{true, false, true, true}},
		// Rumanian and Persian necklaces
		{args{3, 5}, []bool{true, false, true, false, true}},
		// Bulgarian Folk
		{args{3, 7}, []bool{true, false, true, false, true, false, false}},
		// West Africa
		{args{3, 8}, []bool{true, false, false, true, false, false, true, false}},
		// Bulgaria
		{args{4, 7}, []bool{true, false, true, false, true, false, true}},
		// Turkish
		{args{4, 9}, []bool{true, false, true, false, true, false, true, false, false}},
		// Frank Zappa
		{args{4, 11}, []bool{true, false, false, true, false, false, true, false, false, true, false}},
		// Arab
		{args{5, 6}, []bool{true, false, true, true, true, true}},
		// Arab
		{args{5, 7}, []bool{true, false, true, true, false, true, true}},
		// West African
		{args{5, 8}, []bool{true, false, true, true, false, true, true, false}},
		// Arab rhythm, South African and Rumanian necklaces
		{args{5, 9}, []bool{true, false, true, false, true, false, true, false, true}},
		// Classical
		{args{5, 11}, []bool{true, false, true, false, true, false, true, false, true, false, false}},
		// South African
		{args{5, 12}, []bool{true, false, false, true, false, true, false, false, true, false, true, false}},
		// Brazilian necklace
		{args{5, 16}, []bool{true, false, false, true, false, false, true, false, false, true, false, false, true, false, false, false}},
		// Tuareg rhythm of Libya
		{args{7, 8}, []bool{true, false, true, true, true, true, true, true}},
		// West African
		{args{7, 12}, []bool{true, false, true, true, false, true, false, true, true, false, true, false}},
		// Brazilian necklace
		{args{7, 16}, []bool{true, false, false, true, false, true, false, true, false, false, true, false, true, false, true, false}},
		// West and Central African, and Brazilian necklaces
		{args{9, 16}, []bool{true, false, true, true, false, true, false, true, false, true, true, false, true, false, true, false}},
		// Central African
		{args{11, 24}, []bool{true, false, false, true, false, true, false, true, false, true, false, true, false, false, true, false, true, false, true, false, true, false, true, false}},
		// Central African necklace
		{args{13, 24}, []bool{true, false, true, true, false, true, false, true, false, true, false, true, false, true, true, false, true, false, true, false, true, false, true, false}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d-%d", tt.args.accents, tt.args.totalSteps), func(t *testing.T) {
			if got := Rhythm(tt.args.accents, tt.args.totalSteps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Failed:\n%v, want\n%v", got, tt.want)
			}
		})
	}
}
