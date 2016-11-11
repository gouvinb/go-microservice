package shared

import (
	"flag"
	"os"
	"strconv"
	"utils"
)

var (
	flagViewBaseURI   = flag.String("view-base-uri", "", "base of uri for a href")
	flagViewExtension = flag.String("view-extension", "",
		"extension of template files")
	flagViewFolder = flag.String("view-folder", "",
		"name of template's folder")
	flagViewName         = flag.String("view-name", "blank", "default name view")
	flagViewCaching      = flag.Bool("view-caching", false, "enable caching view")
	flagTemplateRoot     = flag.String("template-root", "", "set template root")
	flagTemplateChildren utils.StringSlice
)

func init() {
	flag.Var(&flagTemplateChildren, "template-children",
		"set all childrens template")
}

// GetViewBaseURI return the base URI
func GetViewBaseURI(v View) string {
	if *flagViewBaseURI != "" {
		return *flagViewBaseURI
	} else if os.Getenv("VIEW_BASE_URI") != "" {
		return os.Getenv("VIEW_BASE_URI")
	} else if v.BaseURI != "" {
		return v.BaseURI
	}
	return ""
}

// GetViewExtension return the extension file string
func GetViewExtension(v View) string {
	if *flagViewExtension != "" {
		return *flagViewExtension
	} else if os.Getenv("VIEW_EXTENSION") != "" {
		return os.Getenv("VIEW_EXTENSION")
	} else if v.Extension != "" {
		return v.Extension
	}
	return ""
}

// GetViewFolder return the name folder
func GetViewFolder(v View) string {
	if *flagViewFolder != "" {
		return *flagViewFolder
	} else if os.Getenv("VIEW_FOLDER") != "" {
		return os.Getenv("VIEW_FOLDER")
	} else if v.Folder != "" {
		return v.Folder
	}
	return ""
}

// GetViewName return name of the view
func GetViewName(v View) string {
	if *flagViewName != "" {
		return *flagViewName
	} else if os.Getenv("VIEW_NAME") != "" {
		return os.Getenv("VIEW_NAME")
	} else if v.Name != "" {
		return v.Name
	}
	return ""
}

// IsViewCaching return true if use cache for view
func IsViewCaching(v View) bool {
	value, err := strconv.ParseBool(os.Getenv("VIEW_CACHING"))
	if *flagViewCaching != false {
		return *flagViewCaching
	} else if err == nil {
		return value
	} else if v.Caching != false {
		return v.Caching
	}
	return false
}

// GetTemplateRoot return the name of root template
func GetTemplateRoot(t Template) string {
	if *flagTemplateRoot != "" {
		return *flagTemplateRoot
	} else if os.Getenv("TEMPLATE_ROOT") != "" {
		return os.Getenv("TEMPLATE_ROOT")
	} else if t.Root != "" {
		return t.Root
	}
	return ""
}

// GetTemplateChildren all names of childrens template
func GetTemplateChildren(t Template) utils.StringSlice {
	if len(flagTemplateChildren) > 0 {
		return flagTemplateChildren
	} else if len(t.Children) > 0 {
		return t.Children
	}
	return nil
}
