package main

import "testing"

func TestLowercase(t *testing.T) {
  	svc := stringService{}
    result, _ := svc.Lowercase("aBc");
    if(result != "abc") {
      t.Error("Expected abc got ", result);
    }
}
