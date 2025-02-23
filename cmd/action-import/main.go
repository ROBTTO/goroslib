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
{{ .Goal }}
{{ .Result }}
{{ .Feedback }}
type {{ .Name }}Action struct {
{{- if .RosPkgName }}
    msg.Package ` + "`" + `ros:"{{ .RosPkgName }}"` + "`" + `
{{- end }}
    {{ .Name }}ActionGoal
    {{ .Name }}ActionResult
    {{ .Name }}ActionFeedback
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
	kingpin.CommandLine.Help = "Convert ROS actions into Go structs."

	argGoPkgName := kingpin.Flag("gopackage", "Go package name").Default("main").String()
	argRosPkgName := kingpin.Flag("rospackage", "ROS package name").Default("my_package").String()
	argURL := kingpin.Arg("url", "path or url pointing to a ROS action").Required().String()

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
			return strings.TrimSuffix(filepath.Base(ur.Path), ".action")
		}
		return strings.TrimSuffix(filepath.Base(u), ".action")
	}()

	parts := strings.Split(content, "---")
	if len(parts) != 3 {
		return fmt.Errorf("definition must contain a goal a result and a feedback")
	}

	goalDef, err := msgconv.ParseMessageDefinition(goPkgName, "", name+"ActionGoal", parts[0])
	if err != nil {
		return err
	}

	resultDef, err := msgconv.ParseMessageDefinition(goPkgName, "", name+"ActionResult", parts[1])
	if err != nil {
		return err
	}

	feedbackDef, err := msgconv.ParseMessageDefinition(goPkgName, "", name+"ActionFeedback", parts[2])
	if err != nil {
		return err
	}

	imports := make(map[string]struct{})
	for i := range goalDef.Imports {
		imports[i] = struct{}{}
	}
	for i := range resultDef.Imports {
		imports[i] = struct{}{}
	}
	for i := range feedbackDef.Imports {
		imports[i] = struct{}{}
	}

	goal, err := goalDef.Write()
	if err != nil {
		return err
	}

	result, err := resultDef.Write()
	if err != nil {
		return err
	}

	feedback, err := feedbackDef.Write()
	if err != nil {
		return err
	}

	return tpl.Execute(os.Stdout, map[string]interface{}{
		"GoPkgName":  goPkgName,
		"RosPkgName": rosPkgName,
		"Imports":    imports,
		"Goal":       goal,
		"Result":     result,
		"Feedback":   feedback,
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
