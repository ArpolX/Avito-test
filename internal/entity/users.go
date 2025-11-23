package entity

type TeamMember struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

type Team struct {
	TeamName string       `json:"team_name"`
	Members  []TeamMember `json:"members"`
}

type User struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	TeamName string `json:"team_name"`
	IsActive bool   `json:"is_active"`
}

type SetIsActiveRequest struct { // метод SetFlagIsActive
	UserID   string `json:"user_id"`
	IsActive bool   `json:"is_active"`
}

type TeamSwitchActiveRequest struct {
	TeamName string `json:"team_name"`
}
