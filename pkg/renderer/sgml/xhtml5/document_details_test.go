package xhtml5_test

import (
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("document details", func() {

	Context("header with attributes", func() {

		It("header with author and revision", func() {
			source := `= Document Title
Xavier <xavier@example.com>
v1.0, March 22, 2020: Containment

{author} wrote this doc on {revdate}.
`
			expectedTmpl := `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" lang="en">
<head>
<meta charset="UTF-8"/>
<meta http-equiv="X-UA-Compatible" content="IE=edge"/>
<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
<meta name="generator" content="libasciidoc"/>
<meta name="author" content="Xavier"/>
<title>Document Title</title>
</head>
<body class="article">
<div id="header">
<h1>Document Title</h1>
<div class="details">
<span id="author" class="author">Xavier</span><br/>
<span id="email" class="email"><a href="mailto:xavier@example.com">xavier@example.com</a></span><br/>
<span id="revnumber">version 1.0,</span>
<span id="revdate">March 22, 2020</span>
<br/><span id="revremark">Containment</span>
</div>
</div>
<div id="content">
<div class="paragraph">
<p>Xavier wrote this doc on March 22, 2020.</p>
</div>
</div>
<div id="footer">
<div id="footer-text">
Version 1.0<br/>
Last updated {{ .LastUpdated }}
</div>
</div>
</body>
</html>
`
			now := time.Now()
			Expect(RenderXHTML(source,
				configuration.WithHeaderFooter(true),
				configuration.WithLastUpdated(now),
			)).To(MatchHTMLTemplate(expectedTmpl,
				struct {
					LastUpdated string
				}{
					LastUpdated: now.Format(configuration.LastUpdatedFormat),
				}))
		})

		It("header with 2 authors and no revision", func() {
			source := `= Document Title
John Foo Doe <johndoe@example.com>; Jane Doe <janedoe@example.com>`
			// top-level section is not rendered per-say,
			// but the section will be used to set the HTML page's <title> element
			expectedTmpl := `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" lang="en">
<head>
<meta charset="UTF-8"/>
<meta http-equiv="X-UA-Compatible" content="IE=edge"/>
<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
<meta name="generator" content="libasciidoc"/>
<meta name="author" content="John Foo Doe; Jane Doe"/>
<title>Document Title</title>
</head>
<body class="article">
<div id="header">
<h1>Document Title</h1>
<div class="details">
<span id="author" class="author">John Foo Doe</span><br/>
<span id="email" class="email"><a href="mailto:johndoe@example.com">johndoe@example.com</a></span><br/>
<span id="author2" class="author">Jane Doe</span><br/>
<span id="email2" class="email"><a href="mailto:janedoe@example.com">janedoe@example.com</a></span><br/>
</div>
</div>
<div id="content">
</div>
<div id="footer">
<div id="footer-text">
Last updated {{ .LastUpdated }}
</div>
</div>
</body>
</html>
`
			now := time.Now()
			Expect(RenderXHTML(source,
				configuration.WithHeaderFooter(true),
				configuration.WithLastUpdated(now),
			)).To(MatchHTMLTemplate(expectedTmpl,
				struct {
					LastUpdated string
				}{
					LastUpdated: now.Format(configuration.LastUpdatedFormat),
				}))
		})
	})

	Context("custom header and footer", func() {

		now := time.Now()

		It("with header and footer", func() {
			source := `= Document Title

a paragraph`
			expectedTmpl := `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" lang="en">
<head>
<meta charset="UTF-8"/>
<meta http-equiv="X-UA-Compatible" content="IE=edge"/>
<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
<meta name="generator" content="libasciidoc"/>
<title>Document Title</title>
</head>
<body class="article">
<div id="header">
<h1>Document Title</h1>
</div>
<div id="content">
<div class="paragraph">
<p>a paragraph</p>
</div>
</div>
<div id="footer">
<div id="footer-text">
Last updated {{ .LastUpdated }}
</div>
</div>
</body>
</html>
`
			Expect(RenderXHTML(source,
				configuration.WithHeaderFooter(true),
				configuration.WithLastUpdated(now),
				configuration.WithAttributes(map[string]interface{}{}),
			)).To(MatchHTMLTemplate(expectedTmpl,
				struct {
					LastUpdated string
				}{
					LastUpdated: now.Format(configuration.LastUpdatedFormat),
				}))
		})

		It("with header and without footer", func() {
			source := `= Document Title

a paragraph`
			expectedTmpl := `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" lang="en">
<head>
<meta charset="UTF-8"/>
<meta http-equiv="X-UA-Compatible" content="IE=edge"/>
<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
<meta name="generator" content="libasciidoc"/>
<title>Document Title</title>
</head>
<body class="article">
<div id="header">
<h1>Document Title</h1>
</div>
<div id="content">
<div class="paragraph">
<p>a paragraph</p>
</div>
</div>
</body>
</html>
`
			Expect(RenderXHTML(source,
				configuration.WithHeaderFooter(true),
				configuration.WithLastUpdated(now),
				configuration.WithAttributes(map[string]interface{}{
					types.AttrNoFooter: "",
				}),
			)).To(MatchHTMLTemplate(expectedTmpl,
				struct {
					LastUpdated string
				}{
					LastUpdated: now.Format(configuration.LastUpdatedFormat),
				}))
		})

		It("with author and version label reset", func() {
			source := `= Document Title
Joe Blow <joe.blow@example.com>
1.5, May 21, 1999
:version-label!:

a paragraph`
			expectedTmpl := `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" lang="en">
<head>
<meta charset="UTF-8"/>
<meta http-equiv="X-UA-Compatible" content="IE=edge"/>
<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
<meta name="generator" content="libasciidoc"/>
<meta name="author" content="Joe Blow"/>
<title>Document Title</title>
</head>
<body class="article">
<div id="header">
<h1>Document Title</h1>
<div class="details">
<span id="author" class="author">Joe Blow</span><br/>
<span id="email" class="email"><a href="mailto:joe.blow@example.com">joe.blow@example.com</a></span><br/>
<span id="revnumber">1.5,</span>
<span id="revdate">May 21, 1999</span>
</div>
</div>
<div id="content">
<div class="paragraph">
<p>a paragraph</p>
</div>
</div>
</body>
</html>
`
			Expect(RenderXHTML(source,
				configuration.WithHeaderFooter(true),
				configuration.WithLastUpdated(now),
				configuration.WithAttributes(map[string]interface{}{
					types.AttrNoFooter: "",
				}),
			)).To(MatchHTMLTemplate(expectedTmpl,
				struct {
					LastUpdated string
				}{
					LastUpdated: now.Format(configuration.LastUpdatedFormat),
				}))
		})

		It("with author and custom version label only", func() {
			source := `= Document Title
Joe Blow <joe.blow@example.com>
1.0, May 21, 1999
:version-label: Edition

a paragraph`
			expectedTmpl := `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" lang="en">
<head>
<meta charset="UTF-8"/>
<meta http-equiv="X-UA-Compatible" content="IE=edge"/>
<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
<meta name="generator" content="libasciidoc"/>
<meta name="author" content="Joe Blow"/>
<title>Document Title</title>
</head>
<body class="article">
<div id="header">
<h1>Document Title</h1>
<div class="details">
<span id="author" class="author">Joe Blow</span><br/>
<span id="email" class="email"><a href="mailto:joe.blow@example.com">joe.blow@example.com</a></span><br/>
<span id="revnumber">Edition 1.0,</span>
<span id="revdate">May 21, 1999</span>
</div>
</div>
<div id="content">
<div class="paragraph">
<p>a paragraph</p>
</div>
</div>
</body>
</html>
`
			Expect(RenderXHTML(source,
				configuration.WithHeaderFooter(true),
				configuration.WithLastUpdated(now),
				configuration.WithAttributes(map[string]interface{}{
					types.AttrNoFooter: "",
				}),
			)).To(MatchHTMLTemplate(expectedTmpl,
				struct {
					LastUpdated string
				}{
					LastUpdated: now.Format(configuration.LastUpdatedFormat),
				}))
		})

		It("without header and with footer", func() {
			source := `= Document Title

a paragraph`
			expectedTmpl := `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" lang="en">
<head>
<meta charset="UTF-8"/>
<meta http-equiv="X-UA-Compatible" content="IE=edge"/>
<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
<meta name="generator" content="libasciidoc"/>
<title>Document Title</title>
</head>
<body class="article">
<div id="content">
<div class="paragraph">
<p>a paragraph</p>
</div>
</div>
<div id="footer">
<div id="footer-text">
Last updated {{ .LastUpdated }}
</div>
</div>
</body>
</html>
`
			Expect(RenderXHTML(source,
				configuration.WithHeaderFooter(true),
				configuration.WithLastUpdated(now),
				configuration.WithAttributes(map[string]interface{}{
					types.AttrNoHeader: "",
				}),
			)).To(MatchHTMLTemplate(expectedTmpl,
				struct {
					LastUpdated string
				}{
					LastUpdated: now.Format(configuration.LastUpdatedFormat),
				}))
		})

		It("without header and without footer", func() {
			source := `= Document Title

a paragraph`
			expectedTmpl := `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" lang="en">
<head>
<meta charset="UTF-8"/>
<meta http-equiv="X-UA-Compatible" content="IE=edge"/>
<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
<meta name="generator" content="libasciidoc"/>
<title>Document Title</title>
</head>
<body class="article">
<div id="content">
<div class="paragraph">
<p>a paragraph</p>
</div>
</div>
</body>
</html>
`
			Expect(RenderXHTML(source,
				configuration.WithHeaderFooter(true),
				configuration.WithLastUpdated(now),
				configuration.WithAttributes(map[string]interface{}{
					types.AttrNoHeader: "",
					types.AttrNoFooter: "",
				}),
			)).To(MatchHTMLTemplate(expectedTmpl,
				struct {
					LastUpdated string
				}{
					LastUpdated: now.Format(configuration.LastUpdatedFormat),
				}))
		})
	})
})
