package controllers

import (
	"strings"
	"testing"
)

func TestFindSkills(t *testing.T) {
	cases := []struct {
		desc string
		txt  string
		set  []string
		want []string
	}{
		{"TestFindsHardSkill", "automation is fun", Hard, []string{"automation"}},
		{"TestFindSoftSkill", "clarity is what we need", Soft, []string{"clarity"}},
		{"TestFindSkillsEmptyString", "", Soft, []string{""}},
		{"TestFindSkillsWonkyCase", "ClArItY is what we need", Soft, []string{"clarity"}},
	}
	for _, tc := range cases {
		tc.txt = strings.ToLower(tc.txt)
		got := FindSkills(tc.txt, tc.set)
		if strings.ToLower(got[0]) != tc.want[0] {
			t.Errorf("%s: got %v want %v", tc.desc, tc.txt, tc.want)
		}
	}
}

func TestHasLinkedIn(t *testing.T) {
	cases := []struct {
		desc string
		txt  string
		want bool
	}{
		{"TestHasLinkedIn", "a resume should have a linkedin", true},
		{"TestNoLinkedIn", "some one's resume", false},
	}
	for _, tc := range cases {
		got := HasLinkedIn(tc.txt)
		if got != tc.want {
			t.Errorf("%s: got %v want %v", tc.desc, tc.txt, tc.want)
		}
	}
}
