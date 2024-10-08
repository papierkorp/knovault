// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"pewito/internal/plugins"
	"pewito/internal/themes/defaultTheme/templates/layout"
)

func Playground() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div><h1 class=\"mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-5xl lg:text-6xl dark:text-white\">Playground</h1><h2 class=\"mb-4 text-xl font-bold\">Hello World Plugin Output:</h2>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if plugin, ok := plugins.GetPlugin("HelloWorld"); ok {
				if response, err := plugin.TemplResponse(); err == nil {
					templ_7745c5c3_Err = response.Render(ctx, templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-red-500\">Error: ")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var3 string
					templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(err.Error())
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/themes/defaultTheme/templates/playground.templ`, Line: 18, Col: 64}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-red-500\">HelloWorld plugin not found</p>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h2 class=\"mb-4 text-xl font-bold mt-8\">Theme Changer Plugin:</h2>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if plugin, ok := plugins.GetPlugin("ThemeChanger"); ok {
				if response, err := plugin.TemplResponse(); err == nil {
					templ_7745c5c3_Err = response.Render(ctx, templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-red-500\">Error: ")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var4 string
					templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(err.Error())
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/themes/defaultTheme/templates/playground.templ`, Line: 29, Col: 64}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-red-500\">ThemeChanger plugin not found</p>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h2 class=\"mb-4 text-xl font-bold mt-8\">File Manager Plugin:</h2>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if plugin, ok := plugins.GetPlugin("FileManager"); ok {
				if response, err := plugin.Execute(nil); err == nil {
					if files, ok := response.([]string); ok {
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						for _, file := range files {
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li><a href=\"")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							var templ_7745c5c3_Var5 templ.SafeURL = templ.URL("/docs/" + file)
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var5)))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							var templ_7745c5c3_Var6 string
							templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(file)
							if templ_7745c5c3_Err != nil {
								return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/themes/defaultTheme/templates/playground.templ`, Line: 41, Col: 80}
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></li>")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					} else {
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-red-500\">Unexpected response type from FileManager plugin</p>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-red-500\">Error: ")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var7 string
					templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(err.Error())
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/themes/defaultTheme/templates/playground.templ`, Line: 48, Col: 64}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-red-500\">FileManager plugin not found</p>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h2 class=\"mb-4 text-xl font-bold mt-8\">Markdown Parser Plugin:</h2>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if plugin, ok := plugins.GetPlugin("Parser"); ok {
				if response, err := plugin.Execute(map[string]string{"filePath": "example_markdown.md"}); err == nil {
					if templComponent, ok := response.(templ.Component); ok {
						templ_7745c5c3_Err = templComponent.Render(ctx, templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					} else {
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-red-500\">Unexpected response type from Parser plugin</p>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-red-500\">Error: ")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var8 string
					templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(err.Error())
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/themes/defaultTheme/templates/playground.templ`, Line: 63, Col: 64}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-red-500\">Parser plugin not found</p>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Base("Playground").Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
