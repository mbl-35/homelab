package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func create_vault_secret(path string, data interface{}) {
	http.Post(func() string {
		var buf bytes.Buffer
		err := template.Must(template.New("f").Parse("{{.vault_url}}/v1/secret/data/{{.path}}")).
			Execute(&buf, map[string]interface{}{"vault_url": vault_url, "path": path})
		if err != nil {
			panic(err)
		}
		return buf.String()
	}(), map[interface{}]interface{}{"X-Vault-Token": vault_token}, func() string {
		b, err := json.Marshal(map[string]interface{}{"data": data})
		if err != nil {
			panic(err)
		}
		return string(b)
	}(), false)
}


func setup_gitea_access_token(name string) {
	current_tokens := http.Get(func() string {
		var buf bytes.Buffer
		err := template.Must(template.New("f").Parse("{{.gitea_url}}/api/v1/users/{{.gitea_user}}/tokens")).
			Execute(&buf, map[string]interface{}{"gitea_url": gitea_url, "gitea_user": gitea_user})
		if err != nil {
			panic(err)
		}
		return buf.String()
	}(), false).json()
	if !any(func() func() <-chan bool {
		wait := make(chan struct{})
		yield := make(chan bool)
		go func() {
			defer close(yield)
			<-wait
			for _, token := range current_tokens {
				yield <- token["name"] == name
				<-wait
			}
		}()
		return func() <-chan bool {
			wait <- struct{}{}
			return yield
		}
	}()) {
		resp, err, err := http.Post(func() string {
			var buf bytes.Buffer
			err := template.Must(template.New("f").Parse("{{.gitea_url}}/api/v1/users/{{.gitea_user}}/tokens")).
				Execute(&buf, map[string]interface{}{"gitea_url": gitea_url, "gitea_user": gitea_user})
			if err != nil {
				panic(err)
			}
			return buf.String()
		}(), map[interface{}]interface{}{"Content-Type": "application/json"}, func() string {
			b, err := json.Marshal(map[string]string{"name": name})
			if err != nil {
				panic(err)
			}
			return string(b)
		}(), false)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		if resp.status_code == 201 {
			create_vault_secret(func() string {
				var buf bytes.Buffer
				err = template.Must(template.New("f").Parse("gitea/{{.name}}")).
					Execute(&buf, map[string]interface{}{"name": name})
				if err != nil {
					panic(err)
				}
				return buf.String()
			}(), map[interface{}]interface{}{"token": resp.json()["sha1"]})
		} else {
			fmt.Println(func() string {
				var buf bytes.Buffer
				err = template.Must(template.New("f").Parse("Error creating access token {{.name}} ({{.expr1}})")).Execute(&buf, map[string]interface{}{"name": name, "expr1": resp.status_code})
				if err != nil {
					panic(err)
				}
				return buf.String()
			}())
			fmt.Println(resp.content)
			os.Exit(1)
		}
	}
}




func setup_gitea_oauth_app(name string, redirect_uri string) {
	current_apps := http.Get(func() string {
		var buf bytes.Buffer
		err := template.Must(template.New("f").Parse("{{.gitea_url}}/api/v1/user/applications/oauth2")).
			Execute(&buf, map[string]interface{}{"gitea_url": gitea_url})
		if err != nil {
			panic(err)
		}
		return buf.String()
	}(), false).json()
	if !any(func() func() <-chan bool {
		wait := make(chan struct{})
		yield := make(chan bool)
		go func() {
			defer close(yield)
			<-wait
			for _, app := range current_apps {
				yield <- app["name"] == name
				<-wait
			}
		}()
		return func() <-chan bool {
			wait <- struct{}{}
			return yield
		}
	}()) {
		resp, err, err := http.Post(func() string {
			var buf bytes.Buffer
			err := template.Must(template.New("f").Parse("{{.gitea_url}}/api/v1/user/applications/oauth2")).
				Execute(&buf, map[string]interface{}{"gitea_url": gitea_url})
			if err != nil {
				panic(err)
			}
			return buf.String()
		}(), map[interface{}]interface{}{"Content-Type": "application/json"}, func() string {
			b, err := json.Marshal(
				map[string][]string{"name": name, "redirect_uris": {redirect_uri}},
			)
			if err != nil {
				panic(err)
			}
			return string(b)
		}(), false)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		if resp.status_code == 201 {
			create_vault_secret(func() string {
				var buf bytes.Buffer
				err = template.Must(template.New("f").Parse("gitea/{{.name}}")).
					Execute(&buf, map[string]interface{}{"name": name})
				if err != nil {
					panic(err)
				}
				return buf.String()
			}(), map[interface{}]interface{}{"client_id": resp.json()["client_id"], "client_secret": resp.json()["client_secret"]})
		} else {
			fmt.Println(func() string {
				var buf bytes.Buffer
				err = template.Must(template.New("f").Parse("Error creating OAuth application {{.name}} ({{.expr1}})")).Execute(&buf, map[string]interface{}{"name": name, "expr1": resp.status_code})
				if err != nil {
					panic(err)
				}
				return buf.String()
			}())
			fmt.Println(resp.content)
			os.Exit(1)
		}
	}
}


func main() {
	[]interface {
	} = Console().status("Completing the remaining sorcery")
	gitea_access_tokens := []string{"renovate"}
	gitea_oauth_apps := []map[string]string{map[string]string{"name": "dex", "redirect_uri": func() string {
		var buf bytes.Buffer
		err := template.Must(template.New("f").Parse("https://{{.expr1}}/callback")).Execute(&buf, map[string]interface {
		}{"expr1": client.NetworkingV1Api().read_namespaced_ingress("dex", "dex").spec.rules[0].host})
		if err != nil {
			panic(err)
		}
		return buf.String()
	}()}}
	for _, token_name := range gitea_access_tokens {
		setup_gitea_access_token(token_name)
	}
	for _, app := range gitea_oauth_apps {
		setup_gitea_oauth_app(app["name"], app["redirect_uri"])
	}
}
