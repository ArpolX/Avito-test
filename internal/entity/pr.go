package entity

import "time"

type PullRequest struct {
	PullRequestId     string     `json:"pull_request_id"`
	PullRequestName   string     `json:"pull_request_name"`
	AuthorId          string     `json:"author_id"`
	Status            string     `json:"status"`
	AssignedReviewers []string   `json:"assigned_reviewers"`
	CreatedAt         time.Time  `json:"created_at"`
	MergedAt          *time.Time `json:"merged_at"`
}

// сокращённо
type PullRequestShort struct {
	PullRequestId   string `json:"pull_request_id"`
	PullRequestName string `json:"pull_request_name"`
	AuthorId        string `json:"author_id"`
	Status          string `json:"status"`
}

type CreatePRRequest struct { // метод CreatePRAndAppointReview
	PullRequestID   string `json:"pull_request_id"`
	PullRequestName string `json:"pull_request_name"`
	AuthorID        string `json:"author_id"`
}

type MergePRRequest struct { // метод MarkPRMERGED
	PullRequestID string `json:"pull_request_id"`
}

type RemapReview struct { // метод RemapReview
	PullRequestId string `json:"pull_request_id"`
	OldUserId     string `json:"old_user_id"`
}
