package controllers

import (
	"fmt"
	"testing"
)

func TestFindSkills(t *testing.T) {
	desc := "software development"
	have := FindSkills(desc, Hard)
	want := []string{"software development"}
	fmt.Println("have", have)
	if have != want[0] {
		fmt.Printf("have %v want %v", have, want)
	}
}
