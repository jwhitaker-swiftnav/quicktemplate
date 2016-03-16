// This file is automatically generated by qtc from "basepage.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line templates/basepage.qtpl:1
package templates

//line templates/basepage.qtpl:1
import (
	"io"

	"github.com/valyala/quicktemplate"
)

// This is a base page template. All the other template pages implement this interface.
//

//line templates/basepage.qtpl:3
var (
	_ = io.Copy
	_ = quicktemplate.AcquireByteBuffer
)

//line templates/basepage.qtpl:4
type Page interface {
	//line templates/basepage.qtpl:4
	Title() string
	//line templates/basepage.qtpl:4
	StreamTitle(qw *quicktemplate.Writer)
	//line templates/basepage.qtpl:4
	WriteTitle(qww io.Writer)
	//line templates/basepage.qtpl:4
	Body() string
	//line templates/basepage.qtpl:4
	StreamBody(qw *quicktemplate.Writer)
	//line templates/basepage.qtpl:4
	WriteBody(qww io.Writer)
//line templates/basepage.qtpl:4
}

// Page prints a page implementing Page interface.

//line templates/basepage.qtpl:12
func StreamPageTemplate(qw *quicktemplate.Writer, p Page) {
	//line templates/basepage.qtpl:12
	qw.N().S(`
<html>
	<head>
		<title>`)
	//line templates/basepage.qtpl:15
	p.StreamTitle(qw)
	//line templates/basepage.qtpl:15
	qw.N().S(`</title>
	</head>
	<body>
		<div>
			<a href="/">return to main page</a>
		</div>
		`)
	//line templates/basepage.qtpl:21
	p.StreamBody(qw)
	//line templates/basepage.qtpl:21
	qw.N().S(`
	</body>
</html>
`)
//line templates/basepage.qtpl:24
}

//line templates/basepage.qtpl:24
func WritePageTemplate(qww io.Writer, p Page) {
	//line templates/basepage.qtpl:24
	qw := quicktemplate.AcquireWriter(qww)
	//line templates/basepage.qtpl:24
	StreamPageTemplate(qw, p)
	//line templates/basepage.qtpl:24
	quicktemplate.ReleaseWriter(qw)
//line templates/basepage.qtpl:24
}

//line templates/basepage.qtpl:24
func PageTemplate(p Page) string {
	//line templates/basepage.qtpl:24
	qb := quicktemplate.AcquireByteBuffer()
	//line templates/basepage.qtpl:24
	WritePageTemplate(qb, p)
	//line templates/basepage.qtpl:24
	qs := string(qb.B)
	//line templates/basepage.qtpl:24
	quicktemplate.ReleaseByteBuffer(qb)
	//line templates/basepage.qtpl:24
	return qs
//line templates/basepage.qtpl:24
}

// Base page implementation. Other pages may inherit from it if they need
// overriding only certain Page methods

//line templates/basepage.qtpl:29
type BasePage struct{}

//line templates/basepage.qtpl:30
func (p *BasePage) StreamTitle(qw *quicktemplate.Writer) {
//line templates/basepage.qtpl:30
qw.N().S(`This is a base title`) }

//line templates/basepage.qtpl:30
//line templates/basepage.qtpl:30
func (p *BasePage) WriteTitle(qww io.Writer) {
	//line templates/basepage.qtpl:30
	qw := quicktemplate.AcquireWriter(qww)
	//line templates/basepage.qtpl:30
	p.StreamTitle(qw)
	//line templates/basepage.qtpl:30
	quicktemplate.ReleaseWriter(qw)
//line templates/basepage.qtpl:30
}

//line templates/basepage.qtpl:30
func (p *BasePage) Title() string {
	//line templates/basepage.qtpl:30
	qb := quicktemplate.AcquireByteBuffer()
	//line templates/basepage.qtpl:30
	p.WriteTitle(qb)
	//line templates/basepage.qtpl:30
	qs := string(qb.B)
	//line templates/basepage.qtpl:30
	quicktemplate.ReleaseByteBuffer(qb)
	//line templates/basepage.qtpl:30
	return qs
//line templates/basepage.qtpl:30
}

//line templates/basepage.qtpl:31
func (p *BasePage) StreamBody(qw *quicktemplate.Writer) {
//line templates/basepage.qtpl:31
qw.N().S(`This is a base body`) }

//line templates/basepage.qtpl:31
//line templates/basepage.qtpl:31
func (p *BasePage) WriteBody(qww io.Writer) {
	//line templates/basepage.qtpl:31
	qw := quicktemplate.AcquireWriter(qww)
	//line templates/basepage.qtpl:31
	p.StreamBody(qw)
	//line templates/basepage.qtpl:31
	quicktemplate.ReleaseWriter(qw)
//line templates/basepage.qtpl:31
}

//line templates/basepage.qtpl:31
func (p *BasePage) Body() string {
	//line templates/basepage.qtpl:31
	qb := quicktemplate.AcquireByteBuffer()
	//line templates/basepage.qtpl:31
	p.WriteBody(qb)
	//line templates/basepage.qtpl:31
	qs := string(qb.B)
	//line templates/basepage.qtpl:31
	quicktemplate.ReleaseByteBuffer(qb)
	//line templates/basepage.qtpl:31
	return qs
//line templates/basepage.qtpl:31
}
