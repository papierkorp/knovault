package CardCreator

import (
    "fmt"
    "github.com/a-h/templ"
)

var UIElements = map[string]func(string) templ.Component{
    "title": func(content string) templ.Component {
        return templ.Raw(fmt.Sprintf(`<h2 class="text-xl font-bold">%s</h2>`, content))
    },
    "description": func(content string) templ.Component {
        return templ.Raw(fmt.Sprintf(`<p class="text-sm">%s</p>`, content))
    },
    "type": func(content string) templ.Component {
        return templ.Raw(fmt.Sprintf(`<span class="bg-blue-200 text-blue-800 px-2 py-1 rounded">%s</span>`, content))
    },
    "category": func(content string) templ.Component {
        return templ.Raw(fmt.Sprintf(`<span class="bg-green-200 text-green-800 px-2 py-1 rounded">%s</span>`, content))
    },
    "status": func(content string) templ.Component {
        return templ.Raw(fmt.Sprintf(`<span class="bg-yellow-200 text-yellow-800 px-2 py-1 rounded">%s</span>`, content))
    },
    "tags": func(content string) templ.Component {
        return templ.Raw(fmt.Sprintf(`<div class="flex flex-wrap gap-1">%s</div>`, content))
    },
    "project": func(content string) templ.Component {
        return templ.Raw(fmt.Sprintf(`<span class="text-gray-600">%s</span>`, content))
    },
    "path": func(content string) templ.Component {
        return templ.Raw(fmt.Sprintf(`<span class="text-xs text-gray-400">%s</span>`, content))
    },
}