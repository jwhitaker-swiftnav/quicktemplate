// This file is automatically generated by qtc from "errorpage.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line templates/errorpage.qtpl:1
package templates

//line templates/errorpage.qtpl:1
import (
	"io"

	"github.com/valyala/quicktemplate"
)

// Error page template. Implements BasePage methods.
//

//line templates/errorpage.qtpl:3
var (
	_ = io.Copy
	_ = quicktemplate.AcquireByteBuffer
)

//line templates/errorpage.qtpl:4
type ErrorPage struct {
	// inherit from base page, so its' title is used in error page.
	BasePage

	// error path
	Path []byte
}

//line templates/errorpage.qtpl:14
func (p *ErrorPage) StreamBody(qw *quicktemplate.Writer) {
	//line templates/errorpage.qtpl:14
	qw.N().S(`
	<h1>Error page</h1>
	</div>
		Unsupported path <b>`)
	//line templates/errorpage.qtpl:17
	qw.E().Z(p.Path)
	//line templates/errorpage.qtpl:17
	qw.N().S(`</b>.
	</div>
	Base page body: `)
	//line templates/errorpage.qtpl:19
	p.BasePage.StreamBody(qw)
	//line templates/errorpage.qtpl:19
	qw.N().S(`
`)
//line templates/errorpage.qtpl:20
}

//line templates/errorpage.qtpl:20
func (p *ErrorPage) WriteBody(qww io.Writer) {
	//line templates/errorpage.qtpl:20
	qw := quicktemplate.AcquireWriter(qww)
	//line templates/errorpage.qtpl:20
	p.StreamBody(qw)
	//line templates/errorpage.qtpl:20
	quicktemplate.ReleaseWriter(qw)
//line templates/errorpage.qtpl:20
}

//line templates/errorpage.qtpl:20
func (p *ErrorPage) Body() string {
	//line templates/errorpage.qtpl:20
	qb := quicktemplate.AcquireByteBuffer()
	//line templates/errorpage.qtpl:20
	p.WriteBody(qb)
	//line templates/errorpage.qtpl:20
	qs := string(qb.B)
	//line templates/errorpage.qtpl:20
	quicktemplate.ReleaseByteBuffer(qb)
	//line templates/errorpage.qtpl:20
	return qs
//line templates/errorpage.qtpl:20
}
