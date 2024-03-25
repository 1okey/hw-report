package markdown

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeading(t *testing.T) {
	h := Heading{text: "AN UNEXPECTED PARTY", size: 3}

	assert.Equal(t, "### AN UNEXPECTED PARTY\n", h.ToString())
}

func TestMarkdown(t *testing.T) {
	md := Markdown{}

	h := Heading{text: "AN UNEXPECTED PARTY", size: 3}
	text := Text{text: "In a hole in the ground there lived a hobbit. Not a nasty, dirty, wet hole, filled with the ends of worms and an oozy smell, nor yet a dry, bare, sandy hole with nothing in it to sit down on or to eat: it was a hobbit-hole, and that means comfort.", style: Normal}
	md.Add(h, text)

	assert.Len(t, strings.Split(md.Print(), "\n"), 2)

	one := Text{text: "In a hole in the ground there lived a hobbit. ", style: Normal}
	other := Text{text: "Not a nasty, dirty, wet hole,", style: Normal}
	other_bold := Text{text: "filled with the ends of worms", style: Bold}
	other_second := Text{text: "and an oozy smell, nor yet a dry, bare, sandy hole with nothing in it to sit down on or to eat: it was a hobbit-hole,", style: Normal}
	other_second_italic := Text{text: "and that means comfort.", style: Italic}

	p := Paragraph{text_elements: []Text{one, other, other_bold, other_second, other_second_italic}}
	assert.Equal(t, "In a hole in the ground there lived a hobbit. Not a nasty, dirty, wet hole, __filled with the ends of worms__ and an oozy smell, nor yet a dry, bare, sandy hole with nothing in it to sit down on or to eat: it was a hobbit-hole, *and that means comfort.* \n", p.ToString())

	md.Add(p)
	assert.Len(t, strings.Split(md.Print(), "\n"), 3)
}
