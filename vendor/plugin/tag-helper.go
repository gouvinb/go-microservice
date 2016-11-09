package plugin

import (
	"html/template"
	"log"
	"shared"
)

// TagHelper returns a template.FuncMap
// * JS returns JavaScript tag
// * CSS returns stylesheet tag
// * LINK returns hyperlink tag
func TagHelper(v shared.View) template.FuncMap {
	f := make(template.FuncMap)

	f["JS"] = func(s string) template.HTML {
		path, err := v.ViewAssetTimePath(s)

		if err != nil {
			log.Println("JS Error:", err)
			return template.HTML("<!-- JS Error: " + s + " -->")
		}

		return template.HTML(`<script type="text/javascript" src="` + path +
			`"></script>`)
	}

	f["CSS"] = func(s string) template.HTML {
		path, err := v.ViewAssetTimePath(s)

		if err != nil {
			log.Println("CSS Error:", err)
			return template.HTML("<!-- CSS Error: " + s + " -->")
		}

		return template.HTML(`<link rel="stylesheet" type="text/css" href="` +
			path + `" />`)
	}

	f["LINK"] = func(path, name string) template.HTML {
		return template.HTML(`<a href="` + v.ViewPrependBaseURI(path) + `">` +
			name + `</a>`)
	}

	return f
}
