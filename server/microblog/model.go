package microblog

type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Handle         string `json:"handle"`
	AvatarURL      string `json:"avatar_url"`
	TotalFollowers int    `json:"total_followers"`
	TotalFriends   int    `json:"total_friends"`
}

type Post struct {
	ID         int    `json:"id"`
	AuthorID   int    `json:"author_id"`
	Content    string `json:"content"`
	TotalLikes int    `json:"total_likes"`
}
