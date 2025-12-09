package types

type CommentInfo struct {
	ID          uint             `json:"id"`
	User        UserInfoResponse `json:"user"`
	Content     string           `json:"content"`
	CreatedDate string           `json:"created_db"`
	//可以查看评论下面的评论
	ChildList []*CommentInfo `json:"child_list,omitempty"`
}
type CommentListResponse struct {
	CommentList []CommentInfo `json:"comment_list"`
}
type UserListResponse struct {
	UserList []UserInfoResponse `json:"user_list"`
}
