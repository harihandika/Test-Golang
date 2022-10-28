package tododto

type TodoRequest struct {
	Title           string `json:"title" form:"title"`
	Priority        string `json:"priority" form:"priority"`
	ActivityGroupId string `json:"activity_group_id" form:"activity_group_id"`
	IsActive        bool   `json:"is_active" form:"is_active"`
}
