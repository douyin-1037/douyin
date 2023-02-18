package service

// @path: comment/service/delete_comment.go
// @description: DeleteComment service of comment
// @auth: wan-nan <wan_nan@foxmail.com>
import (
	"context"
	"douyin/comment/infra/dal"
	"douyin/comment/infra/pulsar"
	"douyin/comment/infra/redis"
	"github.com/cloudwego/kitex/pkg/klog"
)

type DeleteCommentService struct {
	ctx context.Context
}

func NewDeleteCommentService(ctx context.Context) *DeleteCommentService {
	return &DeleteCommentService{
		ctx: ctx,
	}
}

func (s *DeleteCommentService) DeleteComment(commentId int64, videoId int64) error {
	// 和CreateComment的逻辑相似：
	// 1. 都要有预先的对key是否还在的判断，不在则需要刷新缓存
	// 2. 都是先写进Redis，同时用协程写进数据库
	// 3. 只有在Redis和DB都写入失败时才会抛回error

	// 先在Redis中查找comment:VideoID这个key是否存在
	/* ———————原因————————
	如果key不存在，可能是key过期了，并不能表示没有评论，所以要先从DB中读，同时刷新缓存
	———————————————————— */
	isKeyExist, err := redis.IsCommentKeyExist(videoId)
	if err != nil {
		klog.Error("IsCommentKeyExist() failed, " + err.Error())
		return err
	}
	// key不存在，刷新缓存
	if isKeyExist == false {
		// 从数据库中 获取 评论列表
		//并发控制
		comments, err := dal.GetCommentList(s.ctx, videoId)
		if err != nil {
			klog.Error("dal.GetCommentList() failed, " + err.Error())
			return err
		}
		// 将评论列表 加入Redis 刷新缓存
		err = redis.AddCommentList(comments)
		if err != nil {
			klog.Error("redis.AddCommentList() failed, " + err.Error())
			return err
		}
	}
	// ——————————————————————————————————————————————————————————————
	// 开启go协程 写数据库
	//errChannel := make(chan error)
	//go func(ch chan error, ctx context.Context, videoId int64, commentUUID int64) {
	//	err := dal.DeleteComment(ctx, commentUUID, videoId)
	//	if err != nil {
	//		// 写数据库也失败，用channel返回错误
	//		klog.Error("Database delete comment failed, " + err.Error())
	//	}
	//	// 不论成功失败都需要将err写入channel，【不然主协程可能会一直阻塞】
	//	// 同时记得【在发送端】关闭channel
	//	ch <- err
	//	close(ch)
	//	return
	//}(errChannel, s.ctx, videoId, commentId)

	// 在主协程中 写Redis
	redisErr := redis.DeleteComment(commentId, videoId)

	//写入Redis错误处理
	if redisErr != nil {
		klog.Error("Redis delete comment failed, " + redisErr.Error())
	}

	//通过pulsar消息队列异步写入到数据库中
	if err := pulsar.DeleteCommentProduce(s.ctx, commentId, videoId); err != nil {
		klog.Error("pulsar send comment delete action failed," + err.Error())
		return err
	}

	return nil
}
