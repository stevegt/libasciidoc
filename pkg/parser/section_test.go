package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("sections", func() {

	Context("in raw documents", func() {

		Context("valid sections", func() {

			It("header only", func() {
				source := "= a header"
				doctitle := []interface{}{
					types.RawLine("a header"),
				}
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   10,
						},
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: doctitle,
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("header with many spaces around content", func() {
				source := "= a header   "
				doctitle := []interface{}{
					types.RawLine("a header   "),
				}
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   13,
						},
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: doctitle,
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("header and paragraph", func() {
				source := `= a header

and a paragraph`

				doctitle := []interface{}{
					types.RawLine("a header"),
				}
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   12,
						},
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: doctitle,
								Elements: []interface{}{
									&types.BlankLine{}, // will be filtered out
								},
							},
						},
					},
					{
						Position: types.Position{
							Start: 12,
							End:   27,
						},
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									types.RawLine("and a paragraph"),
								},
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("two sections with level 0", func() {
				source := `= a first header

= a second header`
				doctitle := []interface{}{
					types.RawLine("a first header"),
				}
				otherDoctitle := []interface{}{
					types.RawLine("a second header"),
				}

				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   18,
						},
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: doctitle,
								Elements: []interface{}{
									&types.BlankLine{},
								},
							},
						},
					},
					{
						Position: types.Position{
							Start: 18,
							End:   35,
						},
						Elements: []interface{}{
							&types.Section{
								Level: 0,
								Title: otherDoctitle,
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("section level 1 alone", func() {
				source := `== section 1`
				section1Title := []interface{}{
					types.RawLine("section 1"),
				}
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   12,
						},
						Elements: []interface{}{
							&types.Section{
								Level: 1,
								Title: section1Title,
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("section level 1 with custom idseparator", func() {
				source := `:idseparator: -
				
== section 1`
				section1Title := []interface{}{
					types.RawLine("section 1"),
				}
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   21,
						},
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name:  types.AttrIDSeparator,
										Value: "-",
									},
									&types.BlankLine{},
								},
							},
						},
					},
					{
						Position: types.Position{
							Start: 21,
							End:   33,
						},
						Elements: []interface{}{
							&types.Section{
								Level: 1,
								Title: section1Title,
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("section level 1 with quoted text", func() {
				source := `==  *2 spaces and bold content*`
				sectionTitle := []interface{}{
					types.RawLine("*2 spaces and bold content*"),
				}
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   31,
						},
						Elements: []interface{}{
							&types.Section{
								Level: 1,
								Title: sectionTitle,
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("section level 0 with nested section level 1", func() {
				source := `= a header

== section 1`
				doctitle := []interface{}{
					types.RawLine("a header"),
				}
				section1Title := []interface{}{
					types.RawLine("section 1"),
				}
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   12,
						},
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: doctitle,
								Elements: []interface{}{
									&types.BlankLine{},
								},
							},
						},
					},
					{
						Position: types.Position{
							Start: 12,
							End:   24,
						},
						Elements: []interface{}{
							&types.Section{
								Level: 1,
								Title: section1Title,
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("section level 0 with nested section level 2", func() {
				source := `= a header

=== section 2`
				doctitle := []interface{}{
					types.RawLine("a header"),
				}
				section2Title := []interface{}{
					types.RawLine("section 2"),
				}
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   12,
						},
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: doctitle,
								Elements: []interface{}{
									&types.BlankLine{},
								},
							},
						},
					},
					{
						Position: types.Position{
							Start: 12,
							End:   25,
						},
						Elements: []interface{}{
							&types.Section{
								Level: 2,
								Title: section2Title,
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("section level 1 with immediate paragraph", func() {
				source := `== a title
and a paragraph`
				section1Title := []interface{}{
					types.RawLine("a title"),
				}
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   11,
						},
						Elements: []interface{}{
							&types.Section{
								Level: 1,
								Title: section1Title,
							},
						},
					},
					{
						Position: types.Position{
							Start: 11,
							End:   26,
						},
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									types.RawLine("and a paragraph"),
								},
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("section level 1 with a paragraph separated by empty line", func() {
				source := `== a title
			
and a paragraph`
				section1Title := []interface{}{
					types.RawLine("a title"),
				}
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   11,
						},
						Elements: []interface{}{
							&types.Section{
								Level: 1,
								Title: section1Title,
							},
						},
					},
					{
						Position: types.Position{
							Start: 11,
							End:   15,
						},
						Elements: []interface{}{
							&types.BlankLine{},
						},
					},
					{
						Position: types.Position{
							Start: 15,
							End:   30,
						},
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									types.RawLine("and a paragraph"),
								},
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("section level 1 with a paragraph separated by non-empty line", func() {
				source := "== a title\n    \nand a paragraph"
				section1Title := []interface{}{
					types.RawLine("a title"),
				}
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   11,
						},
						Elements: []interface{}{
							&types.Section{
								Level: 1,
								Title: section1Title,
							},
						},
					},
					{
						Position: types.Position{
							Start: 11,
							End:   16,
						},
						Elements: []interface{}{
							&types.BlankLine{},
						},
					},
					{
						Position: types.Position{
							Start: 16,
							End:   31,
						},
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									types.RawLine("and a paragraph"),
								},
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("section levels 0, 1, 2, 1", func() {
				source := `= a header

== Section A
a paragraph

=== Section A.a
a paragraph

== Section B
a paragraph`
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   12,
						},
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: []interface{}{
									types.RawLine("a header"),
								},
								Elements: []interface{}{
									&types.BlankLine{},
								},
							},
						},
					},
					{
						Position: types.Position{
							Start: 12,
							End:   25,
						},
						Elements: []interface{}{
							&types.Section{
								Level: 1,
								Title: []interface{}{
									types.RawLine("Section A"),
								},
							},
						},
					},
					{
						Position: types.Position{
							Start: 25,
							End:   37,
						},
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									types.RawLine("a paragraph"),
								},
							},
						},
					},
					{
						Position: types.Position{
							Start: 37,
							End:   38,
						},
						Elements: []interface{}{
							&types.BlankLine{},
						},
					},
					{
						Position: types.Position{
							Start: 38,
							End:   54,
						},
						Elements: []interface{}{
							&types.Section{
								Level: 2,
								Title: []interface{}{
									types.RawLine("Section A.a"),
								},
							},
						},
					},
					{
						Position: types.Position{
							Start: 54,
							End:   66,
						},
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									types.RawLine("a paragraph"),
								},
							},
						},
					},
					{
						Position: types.Position{
							Start: 66,
							End:   67,
						},
						Elements: []interface{}{
							&types.BlankLine{},
						},
					},
					{
						Position: types.Position{
							Start: 67,
							End:   80,
						},
						Elements: []interface{}{
							&types.Section{
								Level: 1,
								Title: []interface{}{
									types.RawLine("Section B"),
								},
							},
						},
					},
					{
						Position: types.Position{
							Start: 80,
							End:   91,
						},
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									types.RawLine("a paragraph"),
								},
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("single section with custom ID", func() {
				source := `[[custom_header]]
== a header`
				sectionTitle := []interface{}{
					types.RawLine("a header"),
				}
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   29,
						},
						Elements: []interface{}{
							&types.Section{
								Level: 1,
								Attributes: types.Attributes{
									types.AttrID: "custom_header",
								},
								Title: sectionTitle,
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("multiple sections with custom IDs", func() {
				source := `[[custom_header]]
= a header

== Section F [[ignored]] [[foo]]

[[bar]]
== Section B
a paragraph`
				doctitle := []interface{}{
					types.RawLine("a header"),
				}
				fooTitle := []interface{}{
					types.RawLine("Section F [[ignored]] [[foo]]"),
				}
				barTitle := []interface{}{
					types.RawLine("Section B"),
				}
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   30,
						},
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: doctitle,
								Attributes: types.Attributes{
									types.AttrID: "custom_header",
								},
								Elements: []interface{}{
									&types.BlankLine{},
								},
							},
						},
					},
					{
						Position: types.Position{
							Start: 30,
							End:   63,
						},
						Elements: []interface{}{
							&types.Section{
								Level: 1,
								Title: fooTitle,
							},
						},
					},
					{
						Position: types.Position{
							Start: 63,
							End:   64,
						},
						Elements: []interface{}{
							&types.BlankLine{},
						},
					},
					{
						Position: types.Position{
							Start: 64,
							End:   85,
						},
						Elements: []interface{}{
							&types.Section{
								Level: 1,
								Attributes: types.Attributes{
									types.AttrID: "bar",
								},
								Title: barTitle,
							},
						},
					},
					{
						Position: types.Position{
							Start: 85,
							End:   96,
						},
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									types.RawLine("a paragraph"),
								},
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("sections with same title", func() {
				source := `== section 1

== section 1`
				section1aTitle := []interface{}{
					types.RawLine("section 1"),
				}
				section1bTitle := []interface{}{
					types.RawLine("section 1"),
				}
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   13,
						},
						Elements: []interface{}{
							&types.Section{
								Level: 1,
								Title: section1aTitle,
							},
						},
					},
					{
						Position: types.Position{
							Start: 13,
							End:   14,
						},
						Elements: []interface{}{
							&types.BlankLine{},
						},
					},
					{
						Position: types.Position{
							Start: 14,
							End:   26,
						},
						Elements: []interface{}{
							&types.Section{
								Level: 1,
								Title: section1bTitle,
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("header with preamble then section level 1", func() {
				source := `= a title
		
a short preamble

== section 1`
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   13,
						},
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: []interface{}{
									types.RawLine("a title"),
								},
								Elements: []interface{}{
									&types.BlankLine{},
								},
							},
						},
					},
					{
						Position: types.Position{
							Start: 13,
							End:   30,
						},
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									types.RawLine("a short preamble"),
								},
							},
						},
					},
					{
						Position: types.Position{
							Start: 30,
							End:   31,
						},
						Elements: []interface{}{
							&types.BlankLine{},
						},
					},
					{
						Position: types.Position{
							Start: 31,
							End:   43,
						},
						Elements: []interface{}{
							&types.Section{
								Level: 1,
								Title: []interface{}{
									types.RawLine("section 1"),
								},
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("header with doc attributes and preamble then section level 1", func() {
				source := `= a title
:toc:
		
a short preamble

== section 1`
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   19,
						},
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: []interface{}{
									types.RawLine("a title"),
								},
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: "toc",
									},
									&types.BlankLine{},
								},
							},
						},
					},
					{
						Position: types.Position{
							Start: 19,
							End:   36,
						},
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									types.RawLine("a short preamble"),
								},
							},
						},
					},
					{
						Position: types.Position{
							Start: 36,
							End:   37,
						},
						Elements: []interface{}{
							&types.BlankLine{},
						},
					},
					{
						Position: types.Position{
							Start: 37,
							End:   49,
						},
						Elements: []interface{}{
							&types.Section{
								Level: 1,
								Title: []interface{}{
									types.RawLine("section 1"),
								},
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("header with 2 paragraphs and CRLFs", func() {
				source := "= a title\r\n\r\na first paragraph\r\n\r\na second paragraph"
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   13,
						},
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: []interface{}{
									types.RawLine("a title"),
								},
								Elements: []interface{}{
									&types.BlankLine{},
								},
							},
						},
					},
					{
						Position: types.Position{
							Start: 13,
							End:   32,
						},
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									types.RawLine("a first paragraph"),
								},
							},
						},
					},
					{
						Position: types.Position{
							Start: 32,
							End:   34,
						},
						Elements: []interface{}{
							&types.BlankLine{},
						},
					},
					{
						Position: types.Position{
							Start: 34,
							End:   52,
						},
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									types.RawLine("a second paragraph"),
								},
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

		})

		Context("invalid sections", func() {

			It("header invalid - missing space", func() {
				source := "=a header"
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   9,
						},
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									types.RawLine("=a header"),
								},
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("header invalid - header space", func() {
				source := " = a header with a prefix space"
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   31,
						},
						Elements: []interface{}{
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrStyle:            types.Literal,
									types.AttrLiteralBlockType: types.LiteralBlockWithSpacesOnFirstLine,
								},
								Elements: []interface{}{
									types.RawLine(" = a header with a prefix space"),
								},
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("header with invalid section1", func() {
				source := `= a header

   == section with prefix space`
				title := []interface{}{
					types.RawLine("a header"),
				}
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   12,
						},
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: title,
								Elements: []interface{}{
									&types.BlankLine{},
								},
							},
						},
					},
					{
						Position: types.Position{
							Start: 12,
							End:   43,
						},
						Elements: []interface{}{
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrStyle:            types.Literal,
									types.AttrLiteralBlockType: types.LiteralBlockWithSpacesOnFirstLine,
								},
								Elements: []interface{}{
									types.RawLine("   == section with prefix space"),
								},
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})
		})
	})

	Context("in final documents", func() {

		Context("valid sections", func() {

			It("header only", func() {
				source := "= a header"
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{
									Content: "a header",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("header with many spaces around content", func() {
				source := "= a header   "
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{
									Content: "a header   ",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("header and paragraph", func() {
				source := `= a header

and a paragraph`

				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "and a paragraph"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section with link in title", func() {
				source := `== link to https://foo.bar
`
				section1aTitle := []interface{}{
					&types.StringElement{Content: "link to "},
					&types.InlineLink{
						Location: &types.Location{
							Scheme: "https://",
							Path:   "foo.bar",
						},
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_link_to_httpsfoo_bar",
							},
							Level: 1,
							Title: section1aTitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"_link_to_httpsfoo_bar": section1aTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_link_to_httpsfoo_bar",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section 0, 1 and paragraph with bold quote", func() {
				source := `= a header
				
== section 1

a paragraph with *bold content*`

				section1Title := []interface{}{
					&types.StringElement{Content: "section 1"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
						},
						&types.Section{
							Level: 1,
							Title: section1Title,
							Attributes: types.Attributes{
								types.AttrID: "_section_1",
							},
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "a paragraph with "},
										&types.QuotedText{
											Kind: types.SingleQuoteBold,
											Elements: []interface{}{
												&types.StringElement{Content: "bold content"},
											},
										},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_section_1": section1Title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_section_1",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("two sections with level 0", func() {
				source := `= a first header

= a second header`
				otherDoctitle := []interface{}{
					&types.StringElement{Content: "a second header"},
				}

				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a first header"},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_a_second_header",
							},
							Level: 0,
							Title: otherDoctitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"_a_second_header": otherDoctitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_a_second_header",
								Level: 0,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section level 1 alone", func() {
				source := `== section 1`
				section1Title := []interface{}{
					&types.StringElement{Content: "section 1"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_1",
							},
							Level: 1,
							Title: section1Title,
						},
					},
					ElementReferences: types.ElementReferences{
						"_section_1": section1Title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_section_1",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section level 1 with quoted text", func() {
				source := `==  *2 spaces and bold content*`
				sectionTitle := []interface{}{
					&types.QuotedText{
						Kind: types.SingleQuoteBold,
						Elements: []interface{}{
							&types.StringElement{Content: "2 spaces and bold content"},
						},
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_2_spaces_and_bold_content",
							},
							Level: 1,
							Title: sectionTitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"_2_spaces_and_bold_content": sectionTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_2_spaces_and_bold_content",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section level 0 with nested section level 1", func() {
				source := `= a header

== section 1`
				section1Title := []interface{}{
					&types.StringElement{Content: "section 1"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_1",
							},
							Level: 1,
							Title: section1Title,
						},
					},
					ElementReferences: types.ElementReferences{
						"_section_1": section1Title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_section_1",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section level 0 with nested section level 2", func() {
				source := `= a header

=== section 2`
				section2Title := []interface{}{
					&types.StringElement{Content: "section 2"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_2",
							},
							Level: 2,
							Title: section2Title,
						},
					},
					ElementReferences: types.ElementReferences{
						"_section_2": section2Title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_section_2",
								Level: 2,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section level 1 with immediate paragraph", func() {
				source := `== a title
and a paragraph`
				section1Title := []interface{}{
					&types.StringElement{Content: "a title"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_a_title",
							},
							Level: 1,
							Title: section1Title,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "and a paragraph"},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_a_title": section1Title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_a_title",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section level 1 with a paragraph separated by empty line", func() {
				source := `== a title
			
and a paragraph`
				section1Title := []interface{}{
					&types.StringElement{Content: "a title"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_a_title",
							},
							Level: 1,
							Title: section1Title,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "and a paragraph"},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_a_title": section1Title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_a_title",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section level 1 with a paragraph separated by non-empty line", func() {
				source := "== a title\n    \nand a paragraph"
				section1Title := []interface{}{
					&types.StringElement{Content: "a title"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_a_title",
							},
							Level: 1,
							Title: section1Title,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "and a paragraph"},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_a_title": section1Title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_a_title",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section levels 0, 1, 2, 1", func() {
				source := `= a header

== Section A
a paragraph

=== Section A.a
a paragraph

== Section B
a paragraph`
				sectionATitle := []interface{}{
					&types.StringElement{Content: "Section A"},
				}
				sectionAaTitle := []interface{}{
					&types.StringElement{Content: "Section A.a"},
				}
				sectionBTitle := []interface{}{
					&types.StringElement{Content: "Section B"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_a",
							},
							Level: 1,
							Title: sectionATitle,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "a paragraph"},
									},
								},
								&types.Section{
									Attributes: types.Attributes{
										types.AttrID: "_section_a_a",
									},
									Level: 2,
									Title: sectionAaTitle,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "a paragraph"},
											},
										},
									},
								},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_b",
							},
							Level: 1,
							Title: sectionBTitle,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "a paragraph"},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_section_a":   sectionATitle,
						"_section_a_a": sectionAaTitle,
						"_section_b":   sectionBTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_section_a",
								Level: 1,
								Children: []*types.ToCSection{
									{
										ID:    "_section_a_a",
										Level: 2,
									},
								},
							},
							{
								ID:    "_section_b",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section levels 0, 1, 2, 2", func() {
				source := `= a header

== Section A
a paragraph

=== Section A.a
a paragraph

=== Section A.b
a paragraph`
				sectionATitle := []interface{}{
					&types.StringElement{Content: "Section A"},
				}
				sectionAaTitle := []interface{}{
					&types.StringElement{Content: "Section A.a"},
				}
				sectionBTitle := []interface{}{
					&types.StringElement{Content: "Section A.b"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_a",
							},
							Level: 1,
							Title: sectionATitle,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "a paragraph"},
									},
								},
								&types.Section{
									Attributes: types.Attributes{
										types.AttrID: "_section_a_a",
									},
									Level: 2,
									Title: sectionAaTitle,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "a paragraph"},
											},
										},
									},
								},
								&types.Section{
									Attributes: types.Attributes{
										types.AttrID: "_section_a_b",
									},
									Level: 2,
									Title: sectionBTitle,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "a paragraph"},
											},
										},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_section_a":   sectionATitle,
						"_section_a_a": sectionAaTitle,
						"_section_a_b": sectionBTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_section_a",
								Level: 1,
								Children: []*types.ToCSection{
									{
										ID:    "_section_a_a",
										Level: 2,
									},
									{
										ID:    "_section_a_b",
										Level: 2,
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section levels 0, 2, 3, 3", func() {
				source := `= a header

=== Section A
a paragraph

==== Section A.a
a paragraph

==== Section A.b
a paragraph`
				sectionATitle := []interface{}{
					&types.StringElement{Content: "Section A"},
				}
				sectionAaTitle := []interface{}{
					&types.StringElement{Content: "Section A.a"},
				}
				sectionAbTitle := []interface{}{
					&types.StringElement{Content: "Section A.b"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_a",
							},
							Level: 2,
							Title: sectionATitle,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "a paragraph"},
									},
								},
								&types.Section{
									Attributes: types.Attributes{
										types.AttrID: "_section_a_a",
									},
									Level: 3,
									Title: sectionAaTitle,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "a paragraph"},
											},
										},
									},
								},
								&types.Section{
									Attributes: types.Attributes{
										types.AttrID: "_section_a_b",
									},
									Level: 3, // level is adjusted
									Title: sectionAbTitle,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "a paragraph"},
											},
										},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_section_a":   sectionATitle,
						"_section_a_a": sectionAaTitle,
						"_section_a_b": sectionAbTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_section_a",
								Level: 2,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section levels 0, 2, 2, 2", func() {
				source := `= a header

=== Section A
a paragraph

=== Section B
a paragraph

=== Section C
a paragraph`
				sectionATitle := []interface{}{
					&types.StringElement{Content: "Section A"},
				}
				sectionBTitle := []interface{}{
					&types.StringElement{Content: "Section B"},
				}
				sectionCTitle := []interface{}{
					&types.StringElement{Content: "Section C"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_a",
							},
							Level: 2,
							Title: sectionATitle,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "a paragraph"},
									},
								},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_b",
							},
							Level: 2,
							Title: sectionBTitle,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "a paragraph"},
									},
								},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_c",
							},
							Level: 2,
							Title: sectionCTitle,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "a paragraph"},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_section_a": sectionATitle,
						"_section_b": sectionBTitle,
						"_section_c": sectionCTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_section_a",
								Level: 2,
							},
							{
								ID:    "_section_b",
								Level: 2,
							},
							{
								ID:    "_section_c",
								Level: 2,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single with custom block ID", func() {
				source := `[[custom_header]]
== a header`
				sectionTitle := []interface{}{
					&types.StringElement{Content: "a header"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID:       "custom_header",
								types.AttrCustomID: true,
							},
							Level: 1,
							Title: sectionTitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"custom_header": sectionTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "custom_header",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single with custom inline ID", func() {
				source := `== a header [[custom_header]]`
				sectionTitle := []interface{}{
					&types.StringElement{Content: "a header"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID:       "custom_header",
								types.AttrCustomID: true,
							},
							Level: 1,
							Title: sectionTitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"custom_header": sectionTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "custom_header",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single with attached inline anchor", func() {
				source := `== a header[[bookmark]]`
				sectionTitle := []interface{}{
					&types.StringElement{Content: "a header"},
					&types.InlineLink{
						Attributes: types.Attributes{
							types.AttrID: "bookmark",
						},
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_a_header",
							},
							Level: 1,
							Title: sectionTitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"_a_header": sectionTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_a_header",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single with attached inline anchor and inline ID", func() {
				source := `== a header[[bookmark]] [[custom_header]]`
				sectionTitle := []interface{}{
					&types.StringElement{Content: "a header"},
					&types.InlineLink{
						Attributes: types.Attributes{
							types.AttrID: "bookmark",
						},
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID:       "custom_header",
								types.AttrCustomID: true,
							},
							Level: 1,
							Title: sectionTitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"custom_header": sectionTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "custom_header",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single with detached inline anchor and inline ID", func() {
				source := `== a header [[bookmark]] [[custom_header]]`
				sectionTitle := []interface{}{
					&types.StringElement{Content: "a header "},
					&types.InlineLink{
						Attributes: types.Attributes{
							types.AttrID: "bookmark",
						},
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID:       "custom_header",
								types.AttrCustomID: true,
							},
							Level: 1,
							Title: sectionTitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"custom_header": sectionTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "custom_header",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("multiple sections with multiple inline custom IDs", func() {
				source := `[[custom_header]]
= a header

== Section F [[ignored]] [[foo]]

[[bar]]
== Section B
a paragraph`
				fooTitle := []interface{}{
					&types.StringElement{Content: "Section F "},
					&types.InlineLink{
						Attributes: types.Attributes{
							types.AttrID: "ignored",
						},
					},
				}
				barTitle := []interface{}{
					&types.StringElement{Content: "Section B"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
							Attributes: types.Attributes{
								types.AttrID:       "custom_header",
								types.AttrCustomID: true,
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID:       "foo",
								types.AttrCustomID: true,
							},
							Level: 1,
							Title: fooTitle,
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID:       "bar",
								types.AttrCustomID: true,
							},
							Level: 1,
							Title: barTitle,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "a paragraph"},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"foo": fooTitle,
						"bar": barTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "foo",
								Level: 1,
							},
							{
								ID:    "bar",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("sections with same title", func() {
				source := `== section 1

== section 1`
				section1aTitle := []interface{}{
					&types.StringElement{Content: "section 1"},
				}
				section1bTitle := []interface{}{
					&types.StringElement{Content: "section 1"},
				}

				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_1",
							},
							Level: 1,
							Title: section1aTitle,
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_1_2",
							},
							Level: 1,
							Title: section1bTitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"_section_1":   section1aTitle,
						"_section_1_2": section1bTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_section_1",
								Level: 1,
							},
							{
								ID:    "_section_1_2",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section level 0 with nested section level 1 and custom ID prefix", func() {
				source := `= a header
:idprefix: custom_

== section 1`
				section1Title := []interface{}{
					&types.StringElement{Content: "section 1"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name:  types.AttrIDPrefix,
									Value: "custom_",
								},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "custom_section_1",
							},
							Level: 1,
							Title: section1Title,
						},
					},
					ElementReferences: types.ElementReferences{
						"custom_section_1": section1Title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "custom_section_1",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section level 0 with nested sections level 1 and custom ID prefixes - with idprefix as doc attribute", func() {
				source := `= a header
:idprefix: custom1a_

== section 1a

:idprefix: custom1b_

== section 1b`
				section1aTitle := []interface{}{
					&types.StringElement{Content: "section 1a"},
				}
				section1bTitle := []interface{}{
					&types.StringElement{Content: "section 1b"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name:  "idprefix",
									Value: "custom1a_",
								},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "custom1a_section_1a",
							},
							Level: 1,
							Title: section1aTitle,
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name:  "idprefix",
									Value: "custom1b_",
								},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "custom1b_section_1b",
							},
							Level: 1,
							Title: section1bTitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"custom1a_section_1a": section1aTitle,
						"custom1b_section_1b": section1bTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "custom1a_section_1a",
								Level: 1,
							},
							{
								ID:    "custom1b_section_1b",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section level 0 with nested sections level 1 and custom ID prefixes - without idprefix as doc attribute", func() {
				source := `= a header

:idprefix: custom1a_

== section 1a

:idprefix: custom1b_

== section 1b`
				section1aTitle := []interface{}{
					&types.StringElement{Content: "section 1a"},
				}
				section1bTitle := []interface{}{
					&types.StringElement{Content: "section 1b"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name:  "idprefix",
									Value: "custom1a_",
								},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "custom1a_section_1a",
							},
							Level: 1,
							Title: section1aTitle,
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name:  "idprefix",
									Value: "custom1b_",
								},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "custom1b_section_1b",
							},
							Level: 1,
							Title: section1bTitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"custom1a_section_1a": section1aTitle,
						"custom1b_section_1b": section1bTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "custom1a_section_1a",
								Level: 1,
							},
							{
								ID:    "custom1b_section_1b",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("header with preamble then section level 1", func() {
				source := `= A Title

a short preamble

== Section 1`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "A Title"},
							},
						},
						&types.Preamble{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "a short preamble"},
									},
								},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_1",
							},
							Level: 1,
							Title: []interface{}{
								&types.StringElement{Content: "Section 1"},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_section_1": []interface{}{
							&types.StringElement{Content: "Section 1"},
						},
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_section_1",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

		})

		Context("invalid sections", func() {

			It("header invalid - too many spaces", func() {
				source := "======= a header"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "======= a header"},
							},
						},
					}}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("header invalid - missing space", func() {
				source := "=a header"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "=a header"},
							},
						},
					}}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("header invalid - header space", func() {
				source := " = a header with a prefix space"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle:            types.Literal,
								types.AttrLiteralBlockType: types.LiteralBlockWithSpacesOnFirstLine,
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: " = a header with a prefix space",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("header with invalid section1", func() {
				source := `= a header

 == section with prefix space`
				title := []interface{}{
					&types.StringElement{Content: "a header"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: title,
						},
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle:            types.Literal,
								types.AttrLiteralBlockType: types.LiteralBlockWithSpacesOnFirstLine,
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: " == section with prefix space",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

		})

		Context("unsupported section syntax", func() {

			It("should not fail with underlined title", func() {
				source := `Document Title
==============
Doc Writer <thedoc@asciidoctor.org>`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "Document Title\n==============\nDoc Writer ",
								},
								&types.SpecialCharacter{
									Name: "<",
								},
								&types.StringElement{
									Content: "thedoc@asciidoctor.org",
								},
								&types.SpecialCharacter{
									Name: ">",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})
	})
})
