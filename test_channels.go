package math

import ("testing";"math")

func TestAverage(t *testing.T){
  var v float64
  v = math.Average([]float64{1,2})
  if v != 1.5 {
    t.Error("Expected 1.5, got ", v)
  }
}
