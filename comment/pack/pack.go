package pack

// @path: comment/pack/pack.go
// @description: pack CommentInfo into []*model.Comment
// @auth: wan-nan <wan_nan@foxmail.com>
import (
	"douyin/code_gen/kitex_gen/commentproto"
	"douyin/comment/infra/dal/model"
	"time"
)

func Comment(comment *model.Comment) *commentproto.CommentInfo {
	return &commentproto.CommentInfo{
		CommentId:  comment.CommentUUId,
		UserId:     comment.UserId,
		Content:    comment.Contents,
		CreateDate: time.Unix(comment.CreateTime, 0).Format("01-02"),
	}
}

func Comments(comments []*model.Comment) []*commentproto.CommentInfo {
	commentInfos := make([]*commentproto.CommentInfo, len(comments))
	for i, comment := range comments {
		commentInfos[i] = Comment(comment)
	}
	return commentInfos
}
