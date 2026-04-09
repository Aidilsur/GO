package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello" + name + ", My name is" + myPage.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Budi"}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Aidil",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}


func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Aidil",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()
	
	TemplateFunctionGlobal(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionCreateGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func(val string) string {
			return strings.ToUpper(val)
		},
	})

	t = template.Must(t.Parse(`{{ upper .Name }}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Aidil Surya",
	})
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()
	
	TemplateFunctionCreateGlobal(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionCreateGlobalPipeLine(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"upper": func(val string) string {
			return strings.ToUpper(val)
		},
	})

	t = template.Must(t.Parse(`{{ upper .Name | sayHello }}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Aidil Surya",
	})
}

func TestTemplateFunctionCreateGlobalPipeLine(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()
	
	TemplateFunctionCreateGlobalPipeLine(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}