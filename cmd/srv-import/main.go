package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/ROBTTO/goroslib/pkg/msgconv"
)

var tpl = template.Must(template.New("").Parse(
	`//nolint:golint
package {{ .GoPkgName }}

import (
{{- range $k, $v := .Imports }}
    "{{ $k }}"
{{- end }}
)
{{ .Request }}
{{ .Response }}
type {{ .Name }} struct {
{{- if .RosPkgName }}
    msg.Package ` + "`" + `ros:"{{ .RosPkgName }}"` + "`" + `
{{- end }}
    {{ .Name }}Req
    {{ .Name }}Res
}
`))

func download(addr string) ([]byte, error) {
	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

func run() error {
	kingpin.CommandLine.Help = "Convert ROS services into Go structs."

	argGoPkgName := kingpin.Flag("gopackage", "Go package name").Default("main").String()
	argRosPkgName := kingpin.Flag("rospackage", "ROS package name").Default("my_package").String()
	argURL := kingpin.Arg("url", "path or url pointing to a ROS service").Required().String()

	kingpin.Parse()

	goPkgName := *argGoPkgName
	rosPkgName := *argRosPkgName
	u := *argURL

	isRemote := strings.HasPrefix(u, "https://") || strings.HasPrefix(u, "http://")

	content, err := func() (string, error) {
		if isRemote {
			byts, err := download(u)
			if err != nil {
				return "", err
			}
			return string(byts), nil
		}

		byts, err := ioutil.ReadFile(u)
		if err != nil {
			return "", err
		}
		return string(byts), nil
	}()
	if err != nil {
		return err
	}

	name := func() string {
		if isRemote {
			ur, _ := url.Parse(u)
			return strings.TrimSuffix(filepath.Base(ur.Path), ".srv")
		}
		return strings.TrimSuffix(filepath.Base(u), ".srv")
	}()

	parts := strings.Split(content, "---")
	if len(parts) != 2 {
		return fmt.Errorf("definition must contain a request and a response")
	}

	reqDef, err := msgconv.ParseMessageDefinition(goPkgName, "", name+"Req", parts[0])
	if err != nil {
		return err
	}

	resDef, err := msgconv.ParseMessageDefinition(goPkgName, "", name+"Res", parts[1])
	if err != nil {
		return err
	}

	imports := make(map[string]struct{})
	for i := range reqDef.Imports {
		imports[i] = struct{}{}
	}
	for i := range resDef.Imports {
		imports[i] = struct{}{}
	}

	request, err := reqDef.Write()
	if err != nil {
		return err
	}

	response, err := resDef.Write()
	if err != nil {
		return err
	}

	return tpl.Execute(os.Stdout, map[string]interface{}{
		"GoPkgName":  goPkgName,
		"RosPkgName": rosPkgName,
		"Imports":    imports,
		"Request":    request,
		"Response":   response,
		"Name":       name,
	})
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERR: %s\n", err)
		os.Exit(1)
	}
}
