package shared

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	childTemplates     []string
	rootTemplate       string
	templateCollection = make(map[string]*template.Template)
	pluginCollection   = make(template.FuncMap)
	mutex              sync.RWMutex
	mutexPlugins       sync.RWMutex
	sessionName        string
	viewInfo           View
)

// Template root and children.
type Template struct {
	Root     string   `json:"Root"`
	Children []string `json:"Children"`
}

// View attributes.
type View struct {
	BaseURI   string   `json:"BaseURI"`
	Extension string   `json:"Extension"`
	Folder    string   `json:"Folder"`
	Name      string   `json:"Name"`
	Caching   bool     `json:"Caching"`
	Template  Template `json:"Template"`
	Vars      map[string]interface{}
	request   *http.Request
}

// ViewConfigure sets the view information.
func ViewConfigure(vi View) {
	viewInfo = vi
}

// ViewReadConfig returns the configuration.
func ViewReadConfig() View {
	return viewInfo
}

// ViewLoadTemplates will set the root and child templates.
func ViewLoadTemplates(rootTemp string, childTemps []string) {
	rootTemplate = rootTemp
	childTemplates = childTemps
}

// ViewLoadPlugins will combine all template.FuncMaps into one map and then set
// the plugins for the templates.
// If a func already exists, it is rewritten, there is no error
func ViewLoadPlugins(fms ...template.FuncMap) {
	// Final FuncMap
	fm := make(template.FuncMap)

	// Loop through the maps
	for _, m := range fms {
		// Loop through each key and value
		for k, v := range m {
			fm[k] = v
		}
	}

	// Load the plugins
	mutexPlugins.Lock()
	pluginCollection = fm
	mutexPlugins.Unlock()
}

// ViewPrependBaseURI prepends the base URI to the string.
func (v *View) ViewPrependBaseURI(s string) string {
	return v.BaseURI + s
}

// ViewNew returns a new view.
func ViewNew(req *http.Request) *View {
	v := &View{}
	v.Vars = make(map[string]interface{})
	v.Vars["AuthLevel"] = "anon"

	v.BaseURI = GetViewBaseURI(viewInfo)
	v.Extension = GetViewExtension(viewInfo)
	v.Folder = GetViewFolder(viewInfo)
	v.Name = GetViewName(viewInfo)

	// Make sure BaseURI is available in the templates
	v.Vars["BaseURI"] = v.BaseURI

	// This is required for the view to access the request
	v.request = req

	// Get session
	sess := SessionInstance(v.request)

	// Set the AuthLevel to auth if the user is logged in
	if sess.Values["id"] != nil {
		v.Vars["AuthLevel"] = "auth"
	}

	return v
}

// ViewAssetTimePath returns a URL with the proper base uri and timestamp appended.
// Works for CSS and JS assets. Determines if local or on the web.
func (v *View) ViewAssetTimePath(s string) (string, error) {
	if strings.HasPrefix(s, "//") {
		return s, nil
	}

	s = strings.TrimLeft(s, "/")
	abs, err := filepath.Abs(s)

	if err != nil {
		return "", err
	}

	time, err2 := ViewFileTime(abs)
	if err2 != nil {
		return "", err2
	}

	return v.ViewPrependBaseURI(s + "?" + time), nil
}

// ViewRenderSingle renders a template to the writer.
func (v *View) ViewRenderSingle(w http.ResponseWriter) {

	// Get the template collection from cache
	/*mutex.RLock()
	tc, ok := templateCollection[v.Name]
	mutex.RUnlock()*/

	// Get the plugin collection
	mutexPlugins.RLock()
	pc := pluginCollection
	mutexPlugins.RUnlock()

	templateList := []string{v.Name}

	// List of template names
	/*templateList := make([]string, 0)
	templateList = append(templateList, rootTemplate)
	templateList = append(templateList, v.Name)
	templateList = append(templateList, childTemplates...)*/

	// Loop through each template and test the full path
	for i, name := range templateList {
		// Get the absolute path of the root template
		path, err := filepath.Abs(v.Folder + string(os.PathSeparator) + name + "." + v.Extension)
		if err != nil {
			http.Error(w, "Template Path Error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		templateList[i] = path
	}

	// Determine if there is an error in the template syntax
	templates, err := template.New(v.Name).Funcs(pc).ParseFiles(templateList...)

	if err != nil {
		http.Error(w, "Template Parse Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Cache the template collection
	/*mutex.Lock()
	templateCollection[v.Name] = templates
	mutex.Unlock()*/

	// Save the template collection
	tc := templates

	// Display the content to the screen
	err = tc.Funcs(pc).ExecuteTemplate(w, v.Name+"."+v.Extension, v.Vars)

	if err != nil {
		http.Error(w, "Template File Error: "+err.Error(), http.StatusInternalServerError)
	}
}

// ViewRender renders a template to the writer.
func (v *View) ViewRender(w http.ResponseWriter) {

	// Get the template collection from cache
	mutex.RLock()
	tc, ok := templateCollection[v.Name]
	mutex.RUnlock()

	// Get the plugin collection
	mutexPlugins.RLock()
	pc := pluginCollection
	mutexPlugins.RUnlock()

	// If the template collection is not cached or caching is disabled
	if !ok || !IsViewCaching(viewInfo) {

		// List of template names
		var templateList []string
		templateList = append(templateList, rootTemplate)
		templateList = append(templateList, v.Name)
		templateList = append(templateList, childTemplates...)

		// Loop through each template and test the full path
		for i, name := range templateList {
			// Get the absolute path of the root template
			path, err := filepath.Abs(v.Folder + string(os.PathSeparator) + name + "." + v.Extension)
			if err != nil {
				http.Error(w, "Template Path Error: "+err.Error(), http.StatusInternalServerError)
				return
			}
			templateList[i] = path
		}

		// Determine if there is an error in the template syntax
		templates, err := template.New(v.Name).Funcs(pc).ParseFiles(templateList...)

		if err != nil {
			http.Error(w, "Template Parse Error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Cache the template collection
		mutex.Lock()
		templateCollection[v.Name] = templates
		mutex.Unlock()

		// Save the template collection
		tc = templates
	}

	// Display the content to the screen
	err := tc.Funcs(pc).ExecuteTemplate(w, rootTemplate+"."+v.Extension, v.Vars)

	if err != nil {
		http.Error(w, "Template File Error: "+err.Error(), http.StatusInternalServerError)
	}
}

// ViewValidate returns true if all the required form values are passed.
func ViewValidate(req *http.Request, required []string) (bool, string) {
	for _, v := range required {
		if req.FormValue(v) == "" {
			return false, v
		}
	}

	return true, ""
}

// ViewRepopulate updates the dst map so the form fields can be refilled.
func ViewRepopulate(list []string, src url.Values, dst map[string]interface{}) {
	for _, v := range list {
		dst[v] = src.Get(v)
	}
}

// ViewFileTime returns the modification time of the file.
func ViewFileTime(name string) (string, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return "", err
	}
	mtime := fi.ModTime().Unix()
	return fmt.Sprintf("%v", mtime), nil
}
