package models

type Post struct {
	Id         string `json:"id"`
	User_id    string `json:"user_id"`
	Title      string `json:"title"`
	Comment    string `json:"comment"`
	Image    string `json:"image"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Deleted_at string `json:"deleted_at"`
}

type CreatePost struct {
	User_id string `json:"user_id"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Image    string `json:"image"`
}

type UpdatePost struct {
	Post_id string `json:"post_id"`
	User_id string `json:"user_id"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Image    string `json:"image"`
}

