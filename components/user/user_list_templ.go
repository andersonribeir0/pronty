// Code generated by templ@v0.2.364 DO NOT EDIT.

package user

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func UserList(users []User) templ.Component {
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
		_, err = templBuffer.WriteString("<div id=\"userList\" class=\"container mx-auto mt-20 mb-8\"><h1 id=\"title\" class=\"text-blue-500 text-5xl text-center\">")
		if err != nil {
			return err
		}
		var_2 := `Users`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1><div id=\"addButton\" class=\"flex justify-end mb-4\">")
		if err != nil {
			return err
		}
		err = AddNewUserButton("addNewUserButton", "Add User").Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div><table class=\"min-w-full border border-gray-300\"><thead><tr><th class=\"px-6 py-3 bg-gray-200 text-left text-xs leading-4 font-medium text-gray-600 uppercase tracking-wider\">")
		if err != nil {
			return err
		}
		var_3 := `Name`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th><th class=\"px-6 py-3 bg-gray-200 text-left text-xs leading-4 font-medium text-gray-600 uppercase tracking-wider\">")
		if err != nil {
			return err
		}
		var_4 := `Email`
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th><th class=\"px-6 py-3 bg-gray-200 text-left text-xs leading-4 font-medium text-gray-600 uppercase tracking-wider\">")
		if err != nil {
			return err
		}
		var_5 := `Actions`
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th></tr></thead><tbody id=\"userRows\">")
		if err != nil {
			return err
		}
		for _, v := range users {
			err = UserItem(v).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</tbody></table></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func AddNewUserButton(id, text string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_6 := templ.GetChildren(ctx)
		if var_6 == nil {
			var_6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<button id=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(id))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" class=\"bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded\" hx-get=\"/user/add\" hx-swap=\"outerHTML\">")
		if err != nil {
			return err
		}
		var var_7 string = text
		_, err = templBuffer.WriteString(templ.EscapeString(var_7))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
