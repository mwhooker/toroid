package github

import (
	"testing"
)

func Payload() string {
	return `{"ref":"refs/heads/master","after":"48cc0109a279aa3b76812e960a922542414decc2","before":"48cc0109a279aa3b76812e960a922542414decc2","created":false,"deleted":false,"forced":false,"compare":"https://github.com/mwhooker/toroid/compare/48cc0109a279...48cc0109a279","commits":[],"head_commit":{"id":"48cc0109a279aa3b76812e960a922542414decc2","distinct":true,"message":"some skelleton","timestamp":"2013-11-17T08:50:40-08:00","url":"https://github.com/mwhooker/toroid/commit/48cc0109a279aa3b76812e960a922542414decc2","author":{"name":"Matthew Hooker","email":"mwhooker@gmail.com","username":"mwhooker"},"committer":{"name":"Matthew Hooker","email":"mwhooker@gmail.com","username":"mwhooker"},"added":["queue/memory.go","queue/queue.go","toroid.go"],"removed":[],"modified":[]},"repository":{"id":14470613,"name":"toroid","url":"https://github.com/mwhooker/toroid","description":"simple stateless distributed job runner","watchers":0,"stargazers":0,"forks":0,"fork":false,"size":76,"owner":{"name":"mwhooker","email":"mwhooker@gmail.com"},"private":false,"open_issues":0,"has_issues":true,"has_downloads":true,"has_wiki":true,"language":"Go","created_at":1384707077,"pushed_at":1384707087,"master_branch":"master"},"pusher":{"name":"none"}}`
}

func TestPayload_github(t *testing.T) {
	gh, err := ParsePayload(Payload())
	if err != nil {
		t.Error(err)
	}
	if gh.Repository.Url != "https://github.com/mwhooker/toroid" {
		t.Error("Unexpected url. Got", gh.Repository.Url)
	}
}
