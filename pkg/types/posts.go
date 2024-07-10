package types

type PostPayload struct {
  Content string `json:"content"`
  Description string `json:"description"`
}

type CommentPayload struct {
  Content string `json:"content"`
  PostID string `json:"postId"`
}
