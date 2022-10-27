package activitydto

type ActivityRequest struct {
	Title string `json:"title" form:"title"`
	Email string `json :"email" form: "email" validate:"required"`
}
