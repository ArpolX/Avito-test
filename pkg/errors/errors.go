package errors

import "errors"

var (
	TEAM_EXISTS  = errors.New("team_name already exists")
	NOT_FOUND    = errors.New("not found")
	PR_EXISTS    = errors.New("PR id already exists")
	PR_MERGED    = errors.New("cannot reassign on merged PR")
	NOT_ASSIGNED = errors.New("reviewer is not assigned to this PR")
	NO_CANDIDATE = errors.New("no active replacement candidate in team")
)
