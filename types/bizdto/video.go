package bizdto

type Video struct {
	ID           int64  `json:"id"`
	Author       *User  `json:"author"`
	PlayAddr     string `json:"play_url"`
	CoverAddr    string `json:"cover_url"`
	LikeCount    int64  `json:"favorite_count"`
	CommentCount int64  `json:"comment_count"`
	IsFavorite   bool   `json:"is_favorite"`
	Title        string `json:"title"`
}
