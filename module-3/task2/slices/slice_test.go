package slices

import (
	"reflect"
	"testing"
)

func TestSortStringsPure(t *testing.T) {
	tests := []struct {
		in       []string
		expected []string
	}{
		{in: []string{}, expected: []string{}},
		{in: []string{"F", "D", "H", "G", "B"}, expected: []string{"B", "D", "F", "G", "H"}},
		{in: []string{"3", "2", "44", "1"}, expected: []string{"1", "2", "3", "44"}},
	}

	for i, tt := range tests {
		before := reflect.DeepEqual(tt.expected, tt.in)
		if got := SortStringsPure(tt.in); !reflect.DeepEqual(tt.expected, got) {
			t.Errorf("Unexpected result in Test #%d, expected %+v, got %+v", i, tt.expected, got)
		}
		after := reflect.DeepEqual(tt.expected, tt.in)
		if !before && after {
			t.Errorf("Unexpected result in Test #%d, pure sorting change order in input data.", i)
		}
	}
}

func TestSortStringsInPlace(t *testing.T) {
	tests := []struct {
		in       []string
		expected []string
	}{
		{in: []string{}, expected: []string{}},
		{in: []string{"F", "D", "H", "G", "B"}, expected: []string{"B", "D", "F", "G", "H"}},
		{in: []string{"3", "2", "44", "1"}, expected: []string{"1", "2", "3", "44"}},
	}

	for i, tt := range tests {
		if SortStringsInPlace(tt.in); !reflect.DeepEqual(tt.expected, tt.in) {
			t.Errorf("Unexpected result in Test #%d, expected %+v, got %+v", i, tt.expected, tt.in)
		}
	}
}

func TestSortUsersPure(t *testing.T) {
	tests := []struct {
		in       []User
		expected []User
	}{
		{in: []User{}, expected: []User{}},
		{
			in: []User{
				{FirstName: "Rob", LastName: "Pike"},
				{FirstName: "Rob", LastName: "Martin"},
				{FirstName: "Russ", LastName: "Cox"},
				{FirstName: "Martin", LastName: "Kleppmann"},
			},
			expected: []User{
				{FirstName: "Martin", LastName: "Kleppmann"},
				{FirstName: "Rob", LastName: "Martin"},
				{FirstName: "Rob", LastName: "Pike"},
				{FirstName: "Russ", LastName: "Cox"},
			},
		},
	}

	for i, tt := range tests {
		before := reflect.DeepEqual(tt.expected, tt.in)
		if got := SortUsersPure(tt.in); !reflect.DeepEqual(tt.expected, got) {
			t.Errorf("Unexpected result in Test #%d, expected %+v, got %+v", i, tt.expected, got)
		}
		after := reflect.DeepEqual(tt.expected, tt.in)
		if !before && after {
			t.Errorf("Unexpected result in Test #%d, pure sorting change order in input data.", i)
		}
	}
}

func TestSortUsersPureDesc(t *testing.T) {
	tests := []struct {
		in       []User
		expected []User
	}{
		{in: []User{}, expected: []User{}},
		{
			in: []User{
				{FirstName: "Rob", LastName: "Pike"},
				{FirstName: "Rob", LastName: "Martin"},
				{FirstName: "Russ", LastName: "Cox"},
				{FirstName: "Martin", LastName: "Kleppmann"},
			},
			expected: []User{
				{FirstName: "Russ", LastName: "Cox"},
				{FirstName: "Rob", LastName: "Pike"},
				{FirstName: "Rob", LastName: "Martin"},
				{FirstName: "Martin", LastName: "Kleppmann"},
			}},
	}

	for i, tt := range tests {
		before := reflect.DeepEqual(tt.expected, tt.in)
		if got := SortUsersPureDesc(tt.in); !reflect.DeepEqual(tt.expected, got) {
			t.Errorf("Unexpected result in Test #%d, expected %+v, got %+v", i, tt.expected, got)
		}
		after := reflect.DeepEqual(tt.expected, tt.in)
		if !before && after {
			t.Errorf("Unexpected result in Test #%d, pure sorting change order in input data.", i)
		}
	}

}
