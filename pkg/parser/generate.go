package parser

//go:generate pigeon -optimize-parser -alternate-entrypoints RawSource,RawDocument,DocumentRawBlock,IncludedFileLine,LabeledListItemTerm,MarkdownQuoteAttribution,QuotedTextSubs,NoneSubs,AttributeSubs,ReplacementSubs,PostReplacementSubs,InlinePassthroughSubs,CalloutSubs,InlineMacroSubs,MarkdownQuoteMacroSubs,BlockAttributes,LineRanges,TagRanges -o parser.go parser.peg
