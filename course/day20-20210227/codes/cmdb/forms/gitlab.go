package forms

import "strings"

type GitlabProjectForm struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
	GitSSHURL  string `json:"git_ssh_url"`
	GitHTTPURL string `json:"git_http_url"`
}

type GitlabForm struct {
	EventName string             `json:"event_name"`
	Commit    string             `json:"checkout_sha"`
	Project   *GitlabProjectForm `json:"project"`
	Ref       string             `json:"ref"`
}

func (f *GitlabForm) TagName() string {
	return strings.TrimPrefix(f.Ref, "refs/tags/")
}
