package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderInlineElements(ctx *renderer.Context, elements []interface{}, options ...lineRendererOption) (string, error) {
	if len(elements) == 0 {
		return "", nil
	}
	// log.Debugf("rendering line with %d element(s)...", len(elements))
	lr := r.newLineRenderer(options...)
	// first pass or rendering, using the provided `renderElementFunc`:
	buf := &strings.Builder{}
	for i, element := range elements {
		renderedElement, err := lr.render(ctx, element)
		if err != nil {
			return "", err
		}
		if i == len(elements)-1 {
			if _, ok := element.(*types.StringElement); ok { // TODO: only for StringElement? or for any kind of element?
				// trim trailing spaces before returning the line
				buf.WriteString(strings.TrimRight(string(renderedElement), " "))
			} else {
				buf.WriteString(renderedElement)
			}
		} else {
			buf.WriteString(renderedElement)
		}
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("rendered inline elements: '%s'", buf.String())
	}
	return buf.String(), nil
}

type renderFunc func(*renderer.Context, interface{}) (string, error)
