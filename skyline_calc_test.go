package skylineCalc

import (
	"testing"
)

func Test_SkylineCalc_Dominate(t *testing.T) {
	sd1 := SkylineData{"1", []float32{1, 2, 3}}
	sd2 := SkylineData{"2", []float32{2, 4, 5}}

}

func Test_SkylineCalc_NonDominate(t *testing.T) {
	sd1 := SkylineData{"1", []float32{1, 2, 3}}
	sd2 := SkylineData{"2", []float32{2, 1, 4}}

}

func Test_SkylineCalc_Be_Dominated(t *testing.T) {
	sd1 := SkylineData{"1", []float32{2, 4, 5}}
	sd2 := SkylineData{"2", []float32{1, 2, 3}}

}

func Test_SkylineCalc_Error(t *testing.T) {
	sd1 := SkylineData{"1", []float32{1, 2, 3}}
	sd2 := SkylineData{"2", []float32{1, 5}}

}
