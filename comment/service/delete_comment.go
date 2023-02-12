package service

// @path: comment/service/delete_comment.go
// @description: DeleteComment service of comment
// @auth: wan-nan <wan_nan@foxmail.com>
import (
	"context"
	"douyin/code_gen/kitex_gen/commentproto"
	"douyin/comment/infra/dal"
	"douyin/comment/infra/redis"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hashicorp/go-multierror"
)

type DeleteCommentService struct {
	ctx context.Context
}

func NewDeleteCommentService(ctx context.Context) *DeleteCommentService {
	return &DeleteCommentService{
		ctx: ctx,
	}
}

func (s *DeleteCommentService) DeleteComment(req *commentproto.DeleteCommentReq) error {
	// 和CreateComment的逻辑相似：
	// 1. 都要有预先的对key是否还在的判断，不在则需要刷新缓存
	// 2. 都是先写进Redis，同时用协程写进数据库
	// 3. 只有在Redis和DB都写入失败时才会抛回error

	// 先在Redis中查找comment:VideoID这个key是否存在
	/* ———————原因————————
	如果key不存在，可能是key过期了，并不能表示没有评论，所以要先从DB中读，同时刷新缓存
	———————————————————— */
	isKeyExist, err := redis.IsCommentKeyExist(req.VideoId)
	if err != nil {
		klog.Error("IsCommentKeyExist() failed " + err.Error())
		return err
	}
	// key不存在，刷新缓存
	if isKeyExist == false {
		// 从数据库中 获取 评论列表
		// TODO 并发控制
		comments, err := dal.GetCommentList(s.ctx, req.VideoId)
		if err != nil {
			klog.Error("dal.GetCommentList() failed " + err.Error())
			return err
		}
		// 将评论列表 加入Redis 刷新缓存
		err = redis.AddCommentList(comments)
		if err != nil {
			klog.Error("redis.AddCommentList() failed " + err.Error())
			return err
		}
	}
	// ——————————————————————————————————————————————————————————————
	// 开启go协程 写数据库
	errChannel := make(chan error)
	go func(ch chan error, ctx context.Context, videoId int64, commentUUId int64) {
		err := dal.DeleteComment(ctx, commentUUId, videoId)
		if err != nil {
			// 写数据库也失败，用channel返回错误
			klog.Error("Database delete comment failed " + err.Error())
		}
		// 不论成功失败都需要将err写入channel，【不然主协程可能会一直阻塞】
		// 同时记得【在发送端】关闭channel
		ch <- err
		close(ch)
		return
	}(errChannel, s.ctx, req.VideoId, req.CommentId)

	// 在主协程中 写Redis
	redisErr := redis.DeleteComment(req.CommentId, req.VideoId)

	// TODO 写入Redis和DB有出错的时候的一致性控制
	if redisErr != nil {
		klog.Error("Redis delete comment failed " + err.Error())
		// return nil, err
		// 这里先不返回，而是去阻塞地等写数据库的结果；数据库也写失败再返回error
		dbError := <-errChannel
		if dbError != nil {
			// 完蛋，数据库和缓存全都写失败了，抛出合并的error
			klog.Error("DB and Redis delete comment both failed " + err.Error())
			return multierror.Append(redisErr, dbError)
		}
	}
	// 只要Redis和DB有一个正确写入，这里就返回空的error
	return nil
}
