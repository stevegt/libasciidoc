package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"

	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderBlankLine(ctx *Context, _ types.BlankLine) (string, error) {
	if ctx.IncludeBlankLine {
		buf := &strings.Builder{}
		if err := r.blankLine.Execute(buf, nil); err != nil {
			return "", err
		}
		log.Debug("rendering blank line")
		return buf.String(), nil
	}
	return "", nil
}

func (r *sgmlRenderer) renderLineBreak() (string, error) {
	buf := &strings.Builder{}
	if err := r.lineBreak.Execute(buf, nil); err != nil {
		return "", err
	}
	return buf.String(), nil
}
