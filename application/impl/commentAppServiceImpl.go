package impl

type CommentAppServiceImpl struct {
}

func NewCommentAppService() CommentAppServiceImpl {
	return *new(CommentAppServiceImpl)
}
