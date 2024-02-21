package api

import (
	"context"
	"fmt"
	"time"
)

type Commit struct {
	URL         string    `json:"url"`
	SHA         string    `json:"sha"`
	NodeID      string    `json:"node_id"`
	HTMLURL     string    `json:"html_url"`
	CommentsURL string    `json:"comments_url"`
	Commit      GitCommit `json:"commit"`
	Author      User      `json:"author"`
	Committer   User      `json:"committer"`
	Parents     []Parent  `json:"parents"`
	Stats       Stats     `json:"stats"`
	Files       []File    `json:"files"`
}

type GitCommit struct {
	URL          string       `json:"url"`
	Author       Author       `json:"author"`
	Committer    Author       `json:"committer"`
	Message      string       `json:"message"`
	Tree         Tree         `json:"tree"`
	CommentCount int          `json:"comment_count"`
	Verification Verification `json:"verification"`
}

type Author struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Date  time.Time `json:"date"`
}

type User struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type Parent struct {
	URL string `json:"url"`
	SHA string `json:"sha"`
}

type Stats struct {
	Additions int `json:"additions"`
	Deletions int `json:"deletions"`
	Total     int `json:"total"`
}

type File struct {
	Filename  string `json:"filename"`
	Additions int    `json:"additions"`
	Deletions int    `json:"deletions"`
	Changes   int    `json:"changes"`
	Status    string `json:"status"`
	RawURL    string `json:"raw_url"`
	BlobURL   string `json:"blob_url"`
	Patch     string `json:"patch"`
}

type Tree struct {
	URL string `json:"url"`
	SHA string `json:"sha"`
}

type Verification struct {
	Verified  bool    `json:"verified"`
	Reason    string  `json:"reason"`
	Signature *string `json:"signature"`
	Payload   *string `json:"payload"`
}

type CommitService service

type CommitParam struct {
	Owner string
	Repo  string
	Ref   string
}

func (s *CommitService) Get(ctx context.Context, p CommitParam) (*Commit, *Response, error) {
	path := fmt.Sprintf("/repos/%v/%v/commits/%v", p.Owner, p.Repo, p.Ref)
	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return nil, nil, err
	}

	commit := new(Commit)
	resp, err := s.client.Do(ctx, req, commit)
	if err != nil {
		return nil, resp, err
	}

	return commit, resp, nil
}
