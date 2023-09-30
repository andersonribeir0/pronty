// Code generated by templ@v0.2.364 DO NOT EDIT.

package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Root(baseTemplate ...templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<html><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>")
		if err != nil {
			return err
		}
		var_2 := `Welcome to pronty`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</title><script src=\"https://unpkg.com/htmx.org@1.9.4\" integrity=\"sha384-zUfuhFKKZCbHTY6aRR46gxiqszMk5tcHjsVFxnUo8VMus4kHGVdIYVbOYYNlKmHV\" crossorigin=\"anonymous\"></script><script src=\"https://unpkg.com/hyperscript.org@0.9.11\"></script><script src=\"https://kit.fontawesome.com/e906f40298.js\" crossorigin=\"anonymous\">")
		if err != nil {
			return err
		}
		var_3 := ``
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script><link rel=\"stylesheet\" href=\"/static/output.css\"><link rel=\"stylesheet\" href=\"/static/custom.css\"></head><body class=\"bg-gray-200\">")
		if err != nil {
			return err
		}
		for _, v := range baseTemplate {
			err = v.Render(ctx, templBuffer)
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</body></html>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
