// Code generated by templ@v0.2.364 DO NOT EDIT.

package user

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func foo(details string) templ.ComponentScript {
	return templ.ComponentScript{
		Name:     `__templ_foo_f26e`,
		Function: `function __templ_foo_f26e(details){console.log(details);}`,
		Call:     templ.SafeScript(`__templ_foo_f26e`, details),
	}
}

func UserDetails(details string) templ.Component {
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
		err = templ.RenderScriptItems(ctx, templBuffer, foo(details))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<p onClick=\"")
		if err != nil {
			return err
		}
		var var_2 templ.ComponentScript = foo(details)
		_, err = templBuffer.WriteString(var_2.Call)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\">")
		if err != nil {
			return err
		}
		var var_3 string = details
		_, err = templBuffer.WriteString(templ.EscapeString(var_3))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
