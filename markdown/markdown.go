package markdown

import (
	"strings"
)

type Serializable interface {
	ToString() string
}

type Heading struct {
	text Text
	size uint8
}

func (h Heading) ToString() string {
	return strings.Repeat("#", int(h.size)) + " " + h.text.ToString() + "\n"
}

type TextStyle string

const (
	Normal TextStyle = "normal"
	Code   TextStyle = "code"
	Bold   TextStyle = "bold"
	Italic TextStyle = "italic"
)

type Text struct {
	text  string
	style TextStyle
}

func (p Text) ToString() string {
	switch p.style {
	default:
	case Normal:
		return p.text

	case Code:
		return " `" + p.text + "` "

	case Bold:
		return " __" + p.text + "__ "

	case Italic:
		return " *" + p.text + "* "
	}

	return ""
}

type Paragraph struct {
	text_elements []Text
}

func (p Paragraph) ToString() string {
	elements := make([]string, 0)

	for _, el := range p.text_elements {
		elements = append(elements, el.ToString())
	}

	return strings.Join(elements, "") + "\n"
}

type Markdown struct {
	elements []Serializable
}

func (d *Markdown) Add(elements ...Serializable) *Markdown {
	d.elements = append(d.elements, elements...)
	return d
}

func (d *Markdown) Flush() *Markdown {
	clear(d.elements)
	return d
}

func (d *Markdown) Print() string {
	output := make([]string, 0)

	for _, e := range d.elements {
		output = append(output, e.ToString())
	}

	return strings.Join(output, "")
}
