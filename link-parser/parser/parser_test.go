package parser

import (
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {

	type Test struct {
		name     string
		file     string
		expected []Link
	}

	tests := []Test{
		{
			name: "Parsing html file has only one 'a' element",
			file: "../testFiles/ex1.html",
			expected: []Link{
				{Href: "/other-page", Text: "A link to another page"},
				{Href: "/dog", Text: "Something in a span Text not in a span Bold text!"},
			},
		},
		{
			name: "Parsing html file has multiple 'a' elements",
			file: "../testFiles/ex2.html",
			expected: []Link{
				{Href: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"},
				{Href: "https://github.com/gophercises", Text: "Gophercises is on Github!"},
			},
		},
		{
			name: "Parsing html file has multiple 'a' elements with nested elements",
			file: "../testFiles/ex3.html",
			expected: []Link{
				{Href: "#", Text: "Login"},
				{Href: "/lost", Text: "Lost? Need help?"},
				{Href: "https://twitter.com/marcusolsson", Text: "@marcusolsson"},
			},
		},
		{
			name: "Parsing html file with comments",
			file: "../testFiles/ex4.html",
			expected: []Link{
				{Href: "/dog-cat", Text: "dog cat"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, err := os.Open(test.file)
			if err != nil {
				t.Errorf("unable to open file: %v", err)
			}

			if got := Parse(f); !reflect.DeepEqual(got, test.expected) {
				t.Errorf("Parse() = %v, want %v", got, test.expected)
			}
		})
	}
}
