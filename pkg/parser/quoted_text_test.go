package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("quoted texts", func() {

	Context("in final documents", func() {

		Context("with single punctuation", func() {

			It("bold text with 1 word", func() {
				source := "*hello*"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{Content: "hello"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("bold text with 2 words", func() {
				source := "*bold    content*"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{Content: "bold    content"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("bold text with 3 words", func() {
				source := "*some bold content*"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{Content: "some bold content"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("italic text with 3 words in single quote", func() {
				source := "_some italic content_"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteItalic,
									Elements: []interface{}{
										&types.StringElement{Content: "some italic content"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("monospace text with 3 words", func() {
				source := "`some monospace content`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: "some monospace content"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("invalid subscript text with 3 words", func() {
				source := "~some subscript content~"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "~some subscript content~"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("invalid superscript text with 3 words", func() {
				source := "^some superscript content^"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "^some superscript content^"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("bold text within italic text", func() {
				source := "_some *bold* content_"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteItalic,
									Elements: []interface{}{
										&types.StringElement{Content: "some "},
										&types.QuotedText{
											Kind: types.SingleQuoteBold,
											Elements: []interface{}{
												&types.StringElement{Content: "bold"},
											},
										},
										&types.StringElement{Content: " content"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("monospace text within bold text within italic quote", func() {
				source := "*some _italic and `monospaced content`_*"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{Content: "some "},
										&types.QuotedText{
											Kind: types.SingleQuoteItalic,
											Elements: []interface{}{
												&types.StringElement{Content: "italic and "},
												&types.QuotedText{
													Kind: types.SingleQuoteMonospace,
													Elements: []interface{}{
														&types.StringElement{Content: "monospaced content"},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("subscript text attached", func() {
				source := "O~2~ is a molecule"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "O"},
								&types.QuotedText{
									Kind: types.SingleQuoteSubscript,
									Elements: []interface{}{
										&types.StringElement{Content: "2"},
									},
								},
								&types.StringElement{Content: " is a molecule"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("superscript text attached", func() {
				source := "M^me^ White"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "M"},
								&types.QuotedText{
									Kind: types.SingleQuoteSuperscript,
									Elements: []interface{}{
										&types.StringElement{Content: "me"},
									},
								},
								&types.StringElement{Content: " White"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("invalid subscript text with 3 words", func() {
				source := "~some subscript content~"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "~some subscript content~"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("bold text across paragraph", func() {
				source := "*some bold\n\ncontent*"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "*some bold"},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "content*"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("italic text across paragraph", func() {
				source := "_some italic\n\ncontent_"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "_some italic"},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "content_"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("monospace text across paragraph", func() {
				source := "`some monospace\n\ncontent`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "`some monospace"},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "content`"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("marked text across paragraph", func() {
				source := "#some marked\n\ncontent#"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "#some marked",
								},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "content#",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("suite of regular text mised with bold, italic and monospaced content", func() {
				source := "a _a_ b *b* c `c`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "a ",
								},
								&types.QuotedText{
									Kind: types.SingleQuoteItalic,
									Elements: []interface{}{
										&types.StringElement{Content: "a"},
									},
								},
								&types.StringElement{
									Content: " b ",
								},
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{Content: "b"},
									},
								},
								&types.StringElement{
									Content: " c ",
								},
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: "c"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with role and other attr", func() {
				source := "[.myrole,and=nothing_else]_italics_"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteItalic,
									Attributes: types.Attributes{
										types.AttrRoles: types.Roles{
											"myrole",
										},
										"and": "nothing_else",
									},
									Elements: []interface{}{
										&types.StringElement{Content: "italics"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with role attribute", func() {
				source := "[myrole]*bold*"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Attributes: types.Attributes{
										types.AttrRoles: types.Roles{
											"myrole",
										},
									},
									Elements: []interface{}{
										&types.StringElement{Content: "bold"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("bold with delimiter within", func() {
				source := "Write result to file *OUT*FILE*."
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "Write result to file ",
								},
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{
											Content: "OUT*FILE",
										},
									},
								},
								&types.StringElement{
									Content: ".",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("italic with delimiter within", func() {
				source := `Write result to file _OUT_FILE_.`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "Write result to file ",
								},
								&types.QuotedText{
									Kind: types.SingleQuoteItalic,
									Elements: []interface{}{
										&types.StringElement{
											Content: "OUT_FILE",
										},
									},
								},
								&types.StringElement{
									Content: ".",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("monospace with delimiter within", func() {
				source := "Write result to file `OUT`FILE`."
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "Write result to file ",
								},
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{
											Content: "OUT`FILE",
										},
									},
								},
								&types.StringElement{
									Content: ".",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("marked with delimiter within", func() {
				source := "Write result to file #OUT#FILE#."
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "Write result to file ",
								},
								&types.QuotedText{
									Kind: types.SingleQuoteMarked,
									Elements: []interface{}{
										&types.StringElement{
											Content: "OUT#FILE",
										},
									},
								},
								&types.StringElement{
									Content: ".",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("with double punctuation", func() {

			It("bold text of 1 word in double quote", func() {
				source := "**hello**"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.DoubleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{Content: "hello"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("italic text with 3 words in double quote", func() {
				source := "__some italic content__"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.DoubleQuoteItalic,
									Elements: []interface{}{
										&types.StringElement{Content: "some italic content"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("monospace text with 3 words in double quote", func() {
				source := "`` some monospace content ``"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.DoubleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: " some monospace content "},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("superscript text within italic text", func() {
				source := "__some ^superscript^ content__"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.DoubleQuoteItalic,
									Elements: []interface{}{
										&types.StringElement{Content: "some "},
										&types.QuotedText{
											Kind: types.SingleQuoteSuperscript,
											Elements: []interface{}{
												&types.StringElement{Content: "superscript"},
											},
										},
										&types.StringElement{Content: " content"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("superscript text within italic text within bold quote", func() {
				source := "**some _italic and ^superscriptcontent^_**"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.DoubleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{Content: "some "},
										&types.QuotedText{
											Kind: types.SingleQuoteItalic,
											Elements: []interface{}{
												&types.StringElement{Content: "italic and "},
												&types.QuotedText{
													Kind: types.SingleQuoteSuperscript,
													Elements: []interface{}{
														&types.StringElement{Content: "superscriptcontent"},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("bold text across paragraph", func() {
				source := "**some bold\n\ncontent**"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "**some bold",
								},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "content**"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("italic text across paragraph", func() {
				source := "__some italic\n\ncontent__"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "__some italic",
								},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "content__"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("monospace text across paragraph", func() {
				source := "``some monospace\n\ncontent``"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "``some monospace",
								},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "content``"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("double marked text across paragraph", func() {
				source := "##some marked\n\ncontent##"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "##some marked",
								},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "content##"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with quoted role attr", func() {
				source := `[."something <wicked>"]**bold**`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.DoubleQuoteBold,
									Attributes: types.Attributes{
										types.AttrRoles: types.Roles{
											"something <wicked>", // TODO: do we need to parse SpecialCharacters in inline attributes?
										},
									},
									Elements: []interface{}{
										&types.StringElement{Content: "bold"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

		})

		Context("quoted text inline", func() {

			It("inline content with bold text", func() {
				source := "a paragraph with *some bold content*"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a paragraph with "},
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{Content: "some bold content"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline content with invalid bold text - use case 1", func() {
				source := "a paragraph with *some bold content"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a paragraph with *some bold content"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline content with invalid bold text - use case 2", func() {
				source := "a paragraph with *some bold content *"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a paragraph with *some bold content *"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline content with invalid bold text - use case 3", func() {
				source := "a paragraph with * some bold content*"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a paragraph with * some bold content*"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("invalid italic text within bold text", func() {
				source := "some *bold and _italic content _ together*."
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "some "},
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{Content: "bold and _italic content _ together"},
									},
								},
								&types.StringElement{Content: "."},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("italic text within invalid bold text", func() {
				source := "some *bold and _italic content_ together *."
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "some *bold and "},
								&types.QuotedText{
									Kind: types.SingleQuoteItalic,
									Elements: []interface{}{
										&types.StringElement{Content: "italic content"},
									},
								},
								&types.StringElement{Content: " together *."},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline content with invalid subscript text - use case 1", func() {
				source := "a paragraph with ~some subscript content"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a paragraph with ~some subscript content"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline content with invalid subscript text - use case 2", func() {
				source := "a paragraph with ~some subscript content ~"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a paragraph with ~some subscript content ~"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline content with invalid subscript text - use case 3", func() {
				source := "a paragraph with ~ some subscript content~"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a paragraph with ~ some subscript content~"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline content with invalid superscript text - use case 1", func() {
				source := "a paragraph with ^some superscript content"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a paragraph with ^some superscript content"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline content with invalid superscript text - use case 2", func() {
				source := "a paragraph with ^some superscript content ^"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a paragraph with ^some superscript content ^"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline content with invalid superscript text - use case 3", func() {
				source := "a paragraph with ^ some superscript content^"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a paragraph with ^ some superscript content^"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("with nested quoted text", func() {

			It("italic text within bold text", func() {
				source := "some *bold and _italic content_ together*."
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "some "},
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{Content: "bold and "},
										&types.QuotedText{
											Kind: types.SingleQuoteItalic,
											Elements: []interface{}{
												&types.StringElement{Content: "italic content"},
											},
										},
										&types.StringElement{Content: " together"},
									},
								},
								&types.StringElement{Content: "."},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single-quote bold within single-quote bold text", func() {
				// kinda invalid content, and Asciidoc has the same way of parsing this content
				source := "*some *nested bold* content*"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{Content: "some *nested bold"},
									},
								},
								&types.StringElement{Content: " content*"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("double-quote bold within double-quote bold text", func() {
				// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "**some **nested bold** content**"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.DoubleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{Content: "some "},
									},
								},
								&types.StringElement{Content: "nested bold"},
								&types.QuotedText{
									Kind: types.DoubleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{Content: " content"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single-quote bold within double-quote bold text", func() {
				// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "**some *nested bold* content**"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.DoubleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{Content: "some "},
										&types.QuotedText{
											Kind: types.SingleQuoteBold,
											Elements: []interface{}{
												&types.StringElement{Content: "nested bold"},
											},
										},
										&types.StringElement{Content: " content"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("double-quote bold within single-quote bold text", func() {
				// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "*some **nested bold** content*"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{Content: "some "},
										&types.QuotedText{
											Kind: types.DoubleQuoteBold,
											Elements: []interface{}{
												&types.StringElement{Content: "nested bold"},
											},
										},
										&types.StringElement{Content: " content"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single-quote italic within single-quote italic text", func() {
				// kinda invalid content, and Asciidoc has the same way of parsing this content
				source := "_some _nested italic_ content_"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteItalic,
									Elements: []interface{}{
										&types.StringElement{Content: "some _nested italic"},
									},
								},
								&types.StringElement{Content: " content_"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("double-quote italic within double-quote italic text", func() {
				// here we don't allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "__some __nested italic__ content__"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.DoubleQuoteItalic,
									Elements: []interface{}{
										&types.StringElement{Content: "some "},
									},
								},
								&types.StringElement{Content: "nested italic"},
								&types.QuotedText{
									Kind: types.DoubleQuoteItalic,
									Elements: []interface{}{
										&types.StringElement{Content: " content"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single-quote italic within double-quote italic text", func() {
				// here we allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "__some _nested italic_ content__"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.DoubleQuoteItalic,
									Elements: []interface{}{
										&types.StringElement{Content: "some "},
										&types.QuotedText{
											Kind: types.SingleQuoteItalic,
											Elements: []interface{}{
												&types.StringElement{Content: "nested italic"},
											},
										},
										&types.StringElement{Content: " content"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("double-quote italic within single-quote italic text", func() {
				// here we allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "_some __nested italic__ content_"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteItalic,
									Elements: []interface{}{
										&types.StringElement{Content: "some "},
										&types.QuotedText{
											Kind: types.DoubleQuoteItalic,
											Elements: []interface{}{
												&types.StringElement{Content: "nested italic"},
											},
										},
										&types.StringElement{Content: " content"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single-quote monospace within single-quote monospace text", func() {
				// kinda invalid content, and Asciidoc has the same way of parsing this content
				source := "`some `nested monospace` content`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: "some `nested monospace"},
									},
								},
								&types.StringElement{Content: " content`"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("double-quote monospace within double-quote monospace text", func() {
				// here we don't allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "``some ``nested monospace`` content``"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.DoubleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: "some "},
									},
								},
								&types.StringElement{Content: "nested monospace"},
								&types.QuotedText{
									Kind: types.DoubleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: " content"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single-quote monospace within double-quote monospace text", func() {
				// here we allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "`some ``nested monospace`` content`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: "some "},
										&types.QuotedText{
											Kind: types.DoubleQuoteMonospace,
											Elements: []interface{}{
												&types.StringElement{Content: "nested monospace"},
											},
										},
										&types.StringElement{Content: " content"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("double-quote monospace within single-quote monospace text", func() {
				// here we allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "`some ``nested monospace`` content`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: "some "},
										&types.QuotedText{
											Kind: types.DoubleQuoteMonospace,
											Elements: []interface{}{
												&types.StringElement{Content: "nested monospace"},
											},
										},
										&types.StringElement{Content: " content"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("unbalanced bold in monospace - case 1", func() {
				source := "`*a`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: "*a"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("unbalanced bold in monospace - case 2", func() {
				source := "`a*b`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: "a*b"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("italic in monospace", func() {
				source := "`_a_`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.QuotedText{
											Kind: types.SingleQuoteItalic,
											Elements: []interface{}{
												&types.StringElement{Content: "a"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("unbalanced italic in monospace", func() {
				source := "`a_b`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: "a_b"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("unparsed bold in monospace", func() {
				source := "`a*b*`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: "a*b*"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("parsed subscript in monospace", func() {
				source := "`a~b~`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: "a"},
										&types.QuotedText{
											Kind: types.SingleQuoteSubscript,
											Elements: []interface{}{
												&types.StringElement{Content: "b"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("multiline in single quoted monospace - case 1", func() {
				source := "`a\nb`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: "a\nb"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("multiline in double quoted monospace - case 1", func() {
				source := "`a\nb`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: "a\nb"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("multiline in single quoted  monospace - case 2", func() {
				source := "`a\n*b*`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: "a\n"},
										&types.QuotedText{
											Kind: types.SingleQuoteBold,
											Elements: []interface{}{
												&types.StringElement{Content: "b"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("multiline in double quoted  monospace - case 2", func() {
				source := "`a\n*b*`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: "a\n"},
										&types.QuotedText{
											Kind: types.SingleQuoteBold,
											Elements: []interface{}{
												&types.StringElement{Content: "b"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("link in bold", func() {
				source := "*a link:/[b]*"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{Content: "a "},
										&types.InlineLink{
											Attributes: types.Attributes{
												types.AttrInlineLinkText: "b",
											},
											Location: &types.Location{
												Path: "/",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("image in bold", func() {
				source := "*a image:foo.png[]*"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{Content: "a "},
										&types.InlineImage{
											Location: &types.Location{
												Path: "foo.png",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("singleplus passthrough in bold", func() {
				source := "*a +image:foo.png[]+*"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{Content: "a "},
										&types.InlinePassthrough{
											Kind: types.SinglePlusPassthrough,
											Elements: []interface{}{
												&types.StringElement{Content: "image:foo.png[]"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("tripleplus passthrough in bold", func() {
				source := "*a +++image:foo.png[]+++*"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{Content: "a "},
										&types.InlinePassthrough{
											Kind: types.TriplePlusPassthrough,
											Elements: []interface{}{
												&types.StringElement{Content: "image:foo.png[]"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("link in italic", func() {
				source := "_a link:/[b]_"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteItalic,
									Elements: []interface{}{
										&types.StringElement{Content: "a "},
										&types.InlineLink{
											Attributes: types.Attributes{
												types.AttrInlineLinkText: "b",
											},
											Location: &types.Location{
												Path: "/",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("image in italic", func() {
				source := "_a image:foo.png[]_"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteItalic,
									Elements: []interface{}{
										&types.StringElement{Content: "a "},
										&types.InlineImage{
											Location: &types.Location{
												Path: "foo.png",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("singleplus passthrough in italic", func() {
				source := "_a +image:foo.png[]+_"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteItalic,
									Elements: []interface{}{
										&types.StringElement{Content: "a "},
										&types.InlinePassthrough{
											Kind: types.SinglePlusPassthrough,
											Elements: []interface{}{
												&types.StringElement{Content: "image:foo.png[]"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("tripleplus passthrough in italic", func() {
				source := "_a +++image:foo.png[]+++_"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteItalic,
									Elements: []interface{}{
										&types.StringElement{Content: "a "},
										&types.InlinePassthrough{
											Kind: types.TriplePlusPassthrough,
											Elements: []interface{}{
												&types.StringElement{Content: "image:foo.png[]"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("link in monospace", func() {
				source := "`a link:/[b]`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: "a "},
										&types.InlineLink{
											Attributes: types.Attributes{
												types.AttrInlineLinkText: "b",
											},
											Location: &types.Location{
												Path: "/",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("image in monospace", func() {
				source := "`a image:foo.png[]`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: "a "},
										&types.InlineImage{
											Location: &types.Location{
												Path: "foo.png",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("singleplus passthrough in monospace", func() {
				source := "`a +image:foo.png[]+`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: "a "},
										&types.InlinePassthrough{
											Kind: types.SinglePlusPassthrough,
											Elements: []interface{}{
												&types.StringElement{Content: "image:foo.png[]"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("tripleplus passthrough in monospace", func() {
				source := "`a +++image:foo.png[]+++`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: "a "},
										&types.InlinePassthrough{
											Kind: types.TriplePlusPassthrough,
											Elements: []interface{}{
												&types.StringElement{Content: "image:foo.png[]"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

		})

		Context("unbalanced quoted text", func() {

			// NOTE: differs from Asciidoctor.

			Context("unbalanced bold text", func() {

				It("unbalanced bold text - extra on left", func() {
					source := "**some bold content*"
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.QuotedText{
										Kind: types.SingleQuoteBold,
										Elements: []interface{}{
											&types.StringElement{Content: "*some bold content"},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("unbalanced bold text - extra on right", func() {
					source := "*some bold content**"
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.QuotedText{
										Kind: types.SingleQuoteBold,
										Elements: []interface{}{
											&types.StringElement{Content: "some bold content"},
										},
									},
									&types.StringElement{Content: "*"},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("inline content with unclosed bold text", func() {
					source := "a paragraph with *some unclosed bold content"
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{Content: "a paragraph with *some unclosed bold content"},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})

			Context("unbalanced italic text", func() {

				It("unbalanced italic text - extra on left", func() {
					source := "__some italic content_"
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.QuotedText{
										Kind: types.SingleQuoteItalic,
										Elements: []interface{}{
											&types.StringElement{Content: "_some italic content"},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("unbalanced italic text - extra on right", func() {
					source := "_some italic content__"
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.QuotedText{
										Kind: types.SingleQuoteItalic,
										Elements: []interface{}{
											&types.StringElement{Content: "some italic content"},
										},
									},
									&types.StringElement{Content: "_"},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("inline content with unclosed italic text", func() {
					source := "a paragraph with _some unclosed italic content"
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{Content: "a paragraph with _some unclosed italic content"},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})

			Context("unbalanced monospace text", func() {

				It("unbalanced monospace text - extra on left", func() {
					source := "``some monospace content`"
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.QuotedText{
										Kind: types.SingleQuoteMonospace,
										Elements: []interface{}{
											&types.StringElement{Content: "`some monospace content"},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("unbalanced monospace text - extra on right", func() {
					source := "`some monospace content``"
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.QuotedText{
										Kind: types.SingleQuoteMonospace,
										Elements: []interface{}{
											&types.StringElement{Content: "some monospace content"},
										},
									},
									&types.StringElement{Content: "`"},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("inline content with unclosed monospace text", func() {
					source := "a paragraph with `some unclosed monospace content"
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{Content: "a paragraph with `some unclosed monospace content"},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})

			Context("unbalanced marked text", func() {

				It("unbalanced marked text - extra on left", func() {
					source := "##some marked content#"
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.QuotedText{
										Kind: types.SingleQuoteMarked,
										Elements: []interface{}{
											&types.StringElement{Content: "#some marked content"},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("unbalanced marked text - extra on right", func() {
					source := "#some marked content##"
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.QuotedText{
										Kind: types.SingleQuoteMarked,
										Elements: []interface{}{
											&types.StringElement{Content: "some marked content"},
										},
									},
									&types.StringElement{Content: "#"},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("inline content with unclosed marked text", func() {
					source := "a paragraph with #some unclosed marked content"
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{Content: "a paragraph with #some unclosed marked content"},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})
		})

		Context("prevented substitution", func() {

			Context("prevented bold text substitution", func() {

				Context("without nested quoted text", func() {

					It("escaped bold text with single backslash", func() {
						source := `\*bold content*`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "*bold content*"},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped bold text with multiple backslashes", func() {
						source := `\\*bold content*`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: `\*bold content*`},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped bold text with double quote", func() {
						source := `\\**bold content**`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: `**bold content**`},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped bold text with double quote and more backslashes", func() {
						source := `\\\**bold content**`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: `\**bold content**`},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped bold text with unbalanced double quote", func() {
						source := `\**bold content*`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: `**bold content*`},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped bold text with unbalanced double quote and more backslashes", func() {
						source := `\\\**bold content*`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: `\\**bold content*`},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})
				})

				Context("with nested quoted text", func() {

					It("escaped bold text with nested italic text", func() {
						source := `\*_italic content_*`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "*"},
										&types.QuotedText{
											Kind: types.SingleQuoteItalic,
											Elements: []interface{}{
												&types.StringElement{Content: "italic content"},
											},
										},
										&types.StringElement{Content: "*"},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped bold text with unbalanced double quote and nested italic test", func() {
						source := `\**_italic content_*`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "**"},
										&types.QuotedText{
											Kind: types.SingleQuoteItalic,
											Elements: []interface{}{
												&types.StringElement{Content: "italic content"},
											},
										},
										&types.StringElement{Content: "*"},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped bold text with nested italic", func() {
						source := `\*bold _and italic_ content*`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "*bold "},
										&types.QuotedText{
											Kind: types.SingleQuoteItalic,
											Elements: []interface{}{
												&types.StringElement{Content: "and italic"},
											},
										},
										&types.StringElement{Content: " content*"},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})
				})

			})

			Context("prevented italic text substitution", func() {

				Context("without nested quoted text", func() {

					It("escaped italic text with single quote", func() {
						source := `\_italic content_`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "_italic content_"},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped italic text with single quote and more backslashes", func() {
						source := `\\_italic content_`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: `\_italic content_`},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped italic text with double quote with 2 backslashes", func() {
						source := `\\__italic content__`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: `__italic content__`},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped italic text with double quote with 3 backslashes", func() {
						source := `\\\__italic content__`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: `\__italic content__`},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped italic text with unbalanced double quote", func() {
						source := `\__italic content_`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: `__italic content_`},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped italic text with unbalanced double quote and more backslashes", func() {
						source := `\\\__italic content_`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: `\\__italic content_`}, // only 1 backslash remove
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})
				})

				Context("with nested quoted text", func() {

					It("escaped italic text with nested monospace text", func() {
						source := `\` + "_`monospace content`_" // gives: \_`monospace content`_
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "_"},
										&types.QuotedText{
											Kind: types.SingleQuoteMonospace,
											Elements: []interface{}{
												&types.StringElement{Content: "monospace content"},
											},
										},
										&types.StringElement{Content: "_"},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped italic text with unbalanced double quote and nested bold test", func() {
						source := `\__*bold content*_`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "__"},
										&types.QuotedText{
											Kind: types.SingleQuoteBold,
											Elements: []interface{}{
												&types.StringElement{Content: "bold content"},
											},
										},
										&types.StringElement{Content: "_"},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped italic text with nested bold text", func() {
						source := `\_italic *and bold* content_`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "_italic "},
										&types.QuotedText{
											Kind: types.SingleQuoteBold,
											Elements: []interface{}{
												&types.StringElement{Content: "and bold"},
											},
										},
										&types.StringElement{Content: " content_"},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})
				})
			})

			Context("prevented monospace text substitution", func() {

				Context("without nested quoted text", func() {

					It("escaped monospace text with single quote", func() {
						source := `\` + "`monospace content`"
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "`monospace content`"}, // backslash removed
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped monospace text with single quote and more backslashes", func() {
						source := `\\` + "`monospace content`"
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: `\` + "`monospace content`"}, // only 1 backslash removed
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped monospace text with double quote", func() {
						source := `\\` + "`monospace content``"
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: `\` + "`monospace content``"}, // 2 back slashes "consumed"
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped monospace text with double quote and more backslashes", func() {
						source := `\\\` + "``monospace content``" // 3 backslashes
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: `\` + "``monospace content``"}, // 2 back slashes "consumed"
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped monospace text with unbalanced double quote", func() {
						source := `\` + "``monospace content`"
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "``monospace content`"},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped monospace text with unbalanced double quote and more backslashes", func() {
						source := `\\\` + "``monospace content`" // 3 backslashes
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: `\\` + "``monospace content`"}, // 2 backslashes removed
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})
				})

				Context("with nested quoted text", func() {

					It("escaped monospace text with nested bold text", func() {
						source := `\` + "`*bold content*`" // gives: \`*bold content*`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "`"},
										&types.QuotedText{
											Kind: types.SingleQuoteBold,
											Elements: []interface{}{
												&types.StringElement{Content: "bold content"},
											},
										},
										&types.StringElement{Content: "`"},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped monospace text with unbalanced double backquote and nested bold test", func() {
						source := `\` + "``*bold content*`"
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "``"},
										&types.QuotedText{
											Kind: types.SingleQuoteBold,
											Elements: []interface{}{
												&types.StringElement{Content: "bold content"},
											},
										},
										&types.StringElement{Content: "`"},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped monospace text with nested bold text", func() {
						source := `\` + "`monospace *and bold* content`"
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "`monospace "},
										&types.QuotedText{
											Kind: types.SingleQuoteBold,
											Elements: []interface{}{
												&types.StringElement{Content: "and bold"},
											},
										},
										&types.StringElement{Content: " content`"},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})
				})
			})

			Context("prevented subscript text substitution", func() {

				Context("without nested quoted text", func() {

					It("escaped subscript text with single quote", func() {
						source := `\~subscriptcontent~`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "~subscriptcontent~"},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped subscript text with single quote and more backslashes", func() {
						source := `\\~subscriptcontent~`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: `\~subscriptcontent~`}, // only 1 backslash removed
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

				})

				Context("with nested quoted text", func() {

					It("escaped subscript text with nested bold text", func() {
						source := `\~*boldcontent*~`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "~"},
										&types.QuotedText{
											Kind: types.SingleQuoteBold,
											Elements: []interface{}{
												&types.StringElement{Content: "boldcontent"},
											},
										},
										&types.StringElement{Content: "~"},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped subscript text with nested bold text", func() {
						source := `\~subscript *and bold* content~`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: `\~subscript `},
										&types.QuotedText{
											Kind: types.SingleQuoteBold,
											Elements: []interface{}{
												&types.StringElement{Content: "and bold"},
											},
										},
										&types.StringElement{Content: " content~"},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})
				})
			})

			Context("prevented superscript text substitution", func() {

				Context("without nested quoted text", func() {

					It("escaped superscript text with single quote", func() {
						source := `\^superscriptcontent^`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "^superscriptcontent^"},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped superscript text with single quote and more backslashes", func() {
						source := `\\^superscriptcontent^`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: `\^superscriptcontent^`}, // only 1 backslash removed
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

				})

				Context("with nested quoted text", func() {

					It("escaped superscript text with nested bold text - case 1", func() {
						source := `\^*bold content*^` // valid escaped superscript since it has no space within
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: `^`},
										&types.QuotedText{
											Kind: types.SingleQuoteBold,
											Elements: []interface{}{
												&types.StringElement{Content: "bold content"},
											},
										},
										&types.StringElement{Content: "^"},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped superscript text with unbalanced double backquote and nested bold test", func() {
						source := `\^*bold content*^`
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "^"},
										&types.QuotedText{
											Kind: types.SingleQuoteBold,
											Elements: []interface{}{
												&types.StringElement{Content: "bold content"},
											},
										},
										&types.StringElement{Content: "^"},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped superscript text with nested bold text - case 2", func() {
						source := `\^superscript *and bold* content^` // invalid superscript text since it has spaces within
						expected := &types.Document{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: `\^superscript `},
										&types.QuotedText{
											Kind: types.SingleQuoteBold,
											Elements: []interface{}{
												&types.StringElement{Content: "and bold"},
											},
										},
										&types.StringElement{Content: " content^"},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})
				})
			})
		})

		Context("nested images", func() {

			It("image in bold", func() {
				source := "*a image:foo.png[]*"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{Content: "a "},
										&types.InlineImage{
											Location: &types.Location{
												Path: "foo.png",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("image in italic", func() {
				source := "_a image:foo.png[]_"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteItalic,
									Elements: []interface{}{
										&types.StringElement{Content: "a "},
										&types.InlineImage{
											Location: &types.Location{
												Path: "foo.png",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("image in monospace", func() {
				source := "`a image:foo.png[]`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: "a "},
										&types.InlineImage{
											Location: &types.Location{
												Path: "foo.png",
											},
										},
									},
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
