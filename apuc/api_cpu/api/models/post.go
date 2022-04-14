package models

type Post struct {
	Id         string `json:"id"`
	User_id    string `json:"user_id"`
	Title      string `json:"title"`
	Comment    string `json:"comment"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Deleted_at string `json:"deleted_at"`
}

type CreatePost struct {
	User_id string `json:"user_id"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
}

type UpdatePost struct {
	Post_id string `json:"post_id"`
	User_id string `json:"user_id"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
}

type ListReq struct {
	Page  string `json:"page"`
	Limit string `json:"limit"`
}

type ById struct {
	User_id string `json:"user_id"`
	Page    string `json:"page"`
	Limit   string `json:"limit"`
}

type Comment struct {
	User_id string `json:"user_id"`
	Post_id string `json:"post_id"`
	Text    string `json:"text"`
}
