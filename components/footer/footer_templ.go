// Code generated by templ@v0.2.364 DO NOT EDIT.

package footer

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Footer() templ.Component {
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
		_, err = templBuffer.WriteString("<footer class=\"bg-gray-800 text-white py-6 fixed bottom-0 w-full\"><div class=\"container mx-auto flex flex-col md:flex-row justify-between items-center px-4\"><div class=\"mb-4 md:mb-0\"><h3 class=\"text-2xl font-semibold\">")
		if err != nil {
			return err
		}
		var_2 := `Pronty`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h3><p class=\"text-sm\">")
		if err != nil {
			return err
		}
		var_3 := `123 Main Street, City`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><p class=\"text-sm\">")
		if err != nil {
			return err
		}
		var_4 := `Phone: (123) 456-7890`
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><p class=\"text-sm\">")
		if err != nil {
			return err
		}
		var_5 := `Email: info@company.com`
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p></div><div class=\"flex justify-center space-x-4\"><a href=\"#\" class=\"text-white hover:text-gray-400 transition duration-300\">")
		if err != nil {
			return err
		}
		var_6 := `Home`
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a><a href=\"#\" class=\"text-white hover:text-gray-400 transition duration-300\">")
		if err != nil {
			return err
		}
		var_7 := `About Us`
		_, err = templBuffer.WriteString(var_7)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a><a href=\"#\" class=\"text-white hover:text-gray-400 transition duration-300\">")
		if err != nil {
			return err
		}
		var_8 := `Services`
		_, err = templBuffer.WriteString(var_8)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a><a href=\"#\" class=\"text-white hover:text-gray-400 transition duration-300\">")
		if err != nil {
			return err
		}
		var_9 := `Contact`
		_, err = templBuffer.WriteString(var_9)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></div></div></footer>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
