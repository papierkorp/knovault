package CardCreator

import (
    "context"
    "fmt"
    "io"
    "strings"
    "github.com/a-h/templ"
)

func RenderCard(layout CardLayout, data DocumentData) templ.Component {
    return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
        _, err = fmt.Fprintf(w, `<div class="card p-4 bg-white shadow-md rounded-lg">`)
        if err != nil {
            return err
        }

        for _, element := range layout.Elements {
            elementFunc, ok := UIElements[element.Type]
            if !ok {
                continue
            }

            content := ""
            switch element.Type {
            case "title":
                content = data.Title
            case "description":
                content = data.Description
            case "type":
                content = data.Type
            case "category":
                content = data.Category
            case "status":
                content = data.Status
            case "tags":
                content = strings.Join(data.Tags, ", ")
            case "project":
                content = data.Project
            case "path":
                content = data.Path
            }

            elementComponent := elementFunc(content)
            err = elementComponent.Render(ctx, w)
            if err != nil {
                return err
            }
        }

        _, err = fmt.Fprintf(w, `</div>`)
        return err
    })
}