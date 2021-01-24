package controllers

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestAnalyze(t *testing.T) {
	cases := []struct {
		desc    string
		resu    string
		jobPost string
		result  *Result
	}{
		{
			desc:    "TestAnalyzeBothSkillsMissing",
			resu:    "A 20USD resume is more complex than a $1.00 but not as much as a 20% fee",
			jobPost: "automation clarity",
			result: &Result{
				LinkedIn:     false,
				HardSkills:   []string{"automation"},
				SoftSkills:   []string{"clarity"},
				ResumeLength: 17,
				Measurable:   3,
			},
		},
		{
			desc:    "TestAnalyzeNoMatchInJobDescr",
			resu:    "resume is more complex",
			jobPost: "a poorly written job description",
			result: &Result{
				LinkedIn:     false,
				HardSkills:   nil,
				SoftSkills:   nil,
				ResumeLength: 4,
				Measurable:   0,
			},
		},
	}
	for _, tc := range cases {
		n := Review{}
		got, _ := n.Analyze(tc.resu, tc.jobPost)
		if !cmp.Equal(got, tc.result) {
			t.Errorf("%s: got %v want %v", tc.desc, got, tc.result)
		}
	}
}

func TestFindSkills(t *testing.T) {
	cases := []struct {
		desc string
		resu string
		set  []string
		want []string
	}{
		{"TestFindsHardSkillInResume", "automation is fun", Hard, []string{"automation"}},
		{"TestFindSoftSkillInResume", "clarity is what we need", Soft, []string{"clarity"}},
		{"TestFindSkillsEmptyString", "", Soft, nil},
		{"TestFindSkillsWonkyCase", "ClArItY is what we need", Soft, []string{"clarity"}},
	}
	for _, tc := range cases {
		got := FindSkills(tc.resu, tc.set)
		if !cmp.Equal(got, tc.want) {
			t.Errorf("%s: got %v want %v", tc.desc, got, tc.want)
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

func TestDiff(t *testing.T) {
	cases := []struct {
		desc   string
		skills []string
		post   []string
		want   []string
	}{
		{"TestDiffHasOne", []string{"eating"}, []string{"C++"}, []string{"C++", "eating"}},
		{"TestDiffNone", []string{"automation"}, []string{"automation"}, nil},
		{"TestDiffEmptyString", []string{""}, []string{"Resume"}, []string{"", "Resume"}},
	}
	for _, tc := range cases {
		got := Diff(tc.skills, tc.post)
		if !cmp.Equal(got, tc.want) {
			t.Errorf("%s: got %v want %v", tc.desc, got, tc.want)
		}
	}
}

func TestRemoveDups(t *testing.T) {
	cases := []struct {
		desc  string
		words []string
		want  []string
	}{
		{"TestDupsRemoves", []string{"soap", "soap", "soap"}, []string{"soap"}},
	}
	for _, tc := range cases {
		got := RemoveDups(tc.words)
		if !cmp.Equal(got, tc.want) {
			t.Errorf("%s: got %v want %v", tc.desc, got, tc.want)
		}
	}
}

func TestMeasurableSkillCount(t *testing.T) {
	cases := []struct {
		desc string
		resu string
		want int
	}{
		{"TestFindNums", "Increased Fight Club recruits by 20% $40 $70.12 ", 3},
	}
	for _, tc := range cases {
		got := MeasurableSkillCount(tc.resu)
		if !cmp.Equal(got, tc.want) {
			t.Errorf("%s: got %v want %v", tc.desc, got, tc.want)
		}
	}
}
