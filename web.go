package main

import (
	"fmt"
	"github.com/go-playground/webhooks/v6/github"
	"net/http"
	"os"
	"strings"
)

const (
	hookUrl     = "/webhooks"
	manualStart = "/run"
)

//goland:noinspection GoUnusedFunction
func serve() {

	http.HandleFunc(hookUrl, webHookHandler)
	http.HandleFunc(manualStart, manualHandler)

	err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), nil)
	checkErr(err)
}

func extractTagFromRef(ref string) string {
	s := strings.Split(ref, "/")
	return s[len(s)-1]
}

func webHookHandler(w http.ResponseWriter, r *http.Request) {

	hook, _ := github.New(github.Options.Secret(os.Getenv("SERVER_GITHUB_SECRET")))

	payload, err := hook.Parse(r, github.ReleaseEvent, github.PullRequestEvent)

	fmt.Println(payload)

	if err != nil {
		if err == github.ErrEventNotFound {
			// ok event wasn;t one of the ones asked to be parsed
		}
	}

	switch payload.(type) {

	case github.ReleasePayload:
		release := payload.(github.ReleasePayload)
		// Do whatever you want from here...
		fmt.Printf("%+v", release)

	case github.PullRequestPayload:
		pullRequest := payload.(github.PullRequestPayload)
		// Do whatever you want from here...
		fmt.Printf("%+v", pullRequest)

	case github.PushPayload:
		push := payload.(github.PushPayload)

		fmt.Printf("%+v", push.Ref)

		path := push.Repository.Name + "/" + extractTagFromRef(push.Ref)
		clone(path, push.Repository.GitURL, push.Ref)
		ComposeUp(path)

	default:
		fmt.Println(payload)
	}
}

func manualHandler(w http.ResponseWriter, r *http.Request) {
	repo := r.URL.Query().Get("repo")
	branch := r.URL.Query().Get("branch")
	project := r.URL.Query().Get("project")

	if repo == "" || branch == "" || project == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	path := project + "/" + branch
	clone(path, repo, branch)
	ComposeUp(path)
}
