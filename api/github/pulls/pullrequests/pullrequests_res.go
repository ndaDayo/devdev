package pullrequests

import "time"

type PullRequests struct {
	URL                string           `json:"url"`
	ID                 int              `json:"id"`
	NodeID             string           `json:"node_id"`
	HtmlURL            string           `json:"html_url"`
	DiffURL            string           `json:"diff_url"`
	PatchURL           string           `json:"patch_url"`
	IssueURL           string           `json:"issue_url"`
	CommitsURL         string           `json:"commits_url"`
	ReviewCommentsURL  string           `json:"review_comments_url"`
	ReviewCommentURL   string           `json:"review_comment_url"`
	CommentsURL        string           `json:"comments_url"`
	StatusesURL        string           `json:"statuses_url"`
	Number             int              `json:"number"`
	State              string           `json:"state"`
	Locked             bool             `json:"locked"`
	Title              string           `json:"title"`
	User               User             `json:"user"`
	Body               string           `json:"body"`
	Labels             []Label          `json:"labels"`
	Milestone          Milestone        `json:"milestone"`
	ActiveLockReason   string           `json:"active_lock_reason"`
	CreatedAt          time.Time        `json:"created_at"`
	UpdatedAt          time.Time        `json:"updated_at"`
	ClosedAt           time.Time        `json:"closed_at"`
	MergedAt           time.Time        `json:"merged_at"`
	MergeCommitSha     string           `json:"merge_commit_sha"`
	Assignee           User             `json:"assignee"`
	Assignees          []User           `json:"assignees"`
	RequestedReviewers []User           `json:"requested_reviewers"`
	RequestedTeams     []Team           `json:"requested_teams"`
	Head               Branch           `json:"head"`
	Base               Branch           `json:"base"`
	Links              PullRequestLinks `json:"_links"`
	AuthorAssociation  string           `json:"author_association"`
	AutoMerge          interface{}      `json:"auto_merge"`
	Draft              bool             `json:"draft"`
}

type User struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HtmlURL           string `json:"html_url"`
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

type Label struct {
	ID          int    `json:"id"`
	NodeID      string `json:"node_id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Default     bool   `json:"default"`
}

type Milestone struct {
	URL          string    `json:"url"`
	HtmlURL      string    `json:"html_url"`
	LabelsURL    string    `json:"labels_url"`
	ID           int       `json:"id"`
	NodeID       string    `json:"node_id"`
	Number       int       `json:"number"`
	State        string    `json:"state"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Creator      User      `json:"creator"`
	OpenIssues   int       `json:"open_issues"`
	ClosedIssues int       `json:"closed_issues"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ClosedAt     time.Time `json:"closed_at"`
	DueOn        time.Time `json:"due_on"`
}

type Team struct {
	ID                  int         `json:"id"`
	NodeID              string      `json:"node_id"`
	URL                 string      `json:"url"`
	HtmlURL             string      `json:"html_url"`
	Name                string      `json:"name"`
	Slug                string      `json:"slug"`
	Description         string      `json:"description"`
	Privacy             string      `json:"privacy"`
	Permission          string      `json:"permission"`
	NotificationSetting string      `json:"notification_setting"`
	MembersURL          string      `json:"members_url"`
	RepositoriesURL     string      `json:"repositories_url"`
	Parent              interface{} `json:"parent"`
}

type Branch struct {
	Label string     `json:"label"`
	Ref   string     `json:"ref"`
	Sha   string     `json:"sha"`
	User  User       `json:"user"`
	Repo  Repository `json:"repo"`
}

type Repository struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}

type PullRequestLinks struct {
	Self           Link `json:"self"`
	Html           Link `json:"html"`
	Issue          Link `json:"issue"`
	Comments       Link `json:"comments"`
	ReviewComments Link `json:"review_comments"`
	ReviewComment  Link `json:"review_comment"`
	Commits        Link `json:"commits"`
	Statuses       Link `json:"statuses"`
}

type Link struct {
	Href string `json:"href"`
}
