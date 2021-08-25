package main

import (
	"fmt"
	"github.com/go-playground/webhooks/v6/github"
	"net/http"
	"os"
	"sync"
)

const (
	path = "/webhooks"
)

//goland:noinspection GoUnusedFunction
func serve(wg *sync.WaitGroup) {
	defer wg.Done()

	hook, _ := github.New(github.Options.Secret(os.Getenv("SERVER_GITHUB_SECRET")))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
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

		default:
			fmt.Println(payload)

		}
	})
	err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), nil)
	checkerr(err)
}
