package link

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFind(t *testing.T) {
	tests := []struct {
		filename string
		want     []Link
		wantErr  bool
	}{
		{
			filename: "ex1.html",
			want: []Link{
				{
					Href: "/other-page",
					Text: "A link to another page",
				},
			},
		},
		{
			filename: "ex2.html",
			want: []Link{
				{
					Href: "https://www.twitter.com/joncalhoun",
					Text: "Check me out on twitter",
				},
				{
					Href: "https://github.com/gophercises",
					Text: "Gophercises is on Github !",
				},
			},
		},
		{
			filename: "ex3.html",
			want: []Link{
				{
					Href: "#",
					Text: "Login",
				},
				{
					Href: "/lost",
					Text: "Lost? Need help?",
				},
				{
					Href: "https://twitter.com/marcusolsson",
					Text: "@marcusolsson",
				},
			},
		},
		{
			filename: "ex4.html",
			want: []Link{
				{
					Href: "/dog-cat",
					Text: "dog cat",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.filename, func(t *testing.T) {
			filename := filepath.Join("..", "..", tt.filename)
			f, err := os.Open(filename)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()
			got, err := Find(f)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !equal(got, tt.want) {
				t.Errorf("Find() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func equal(got, want []Link) bool {
	if len(got) != len(want) {
		return false
	}
	wm := toMap(want)
	for _, v := range got {
		if text, ok := wm[v.Href]; !ok || text != v.Text {
			return false
		}
	}
	return true
}

func toMap(links []Link) map[string]string {
	m := make(map[string]string, len(links))
	for _, v := range links {
		m[v.Href] = v.Text
	}
	return m
}
