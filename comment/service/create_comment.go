package service

// @path: comment/service/create_comment.go
// @description: CreateComment service of comment
// @auth: wan-nan <wan_nan@foxmail.com>
import (
	"context"
	"douyin/code_gen/kitex_gen/commentproto"
	"douyin/comment/infra/dal"
	"douyin/comment/infra/redis"
	redisModel "douyin/comment/infra/redis/model"
	"douyin/common/util"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hashicorp/go-multierror"
	"time"
)

type CreateCommentService struct {
	ctx context.Context
}

func NewCreateCommentService(ctx context.Context) *CreateCommentService {
	return &CreateCommentService{
		ctx: ctx,
	}
}

func (s *CreateCommentService) CreateComment(req *commentproto.CreateCommentReq) (*commentproto.CommentInfo, error) {
	// 先在Redis中查找comment:VideoID这个key是否存在
	/* ———————原因————————
	因为key设置的过期时间都是第一天
	假如B第一天评论了A视频，缓存和数据库都会这条记录，第二天C评论了A，数据库查A的评论有2条记录，但是由于第一天的缓存过期了，所以redis里面只会存在C评论A的记录
	这就会出现不一致，假如你这个时候查A的评论记录，那么会击中缓存，然后只返回一条，这显然是错误的
	———————————————————— */
	isKeyExist, err := redis.IsCommentKeyExist(req.VideoId)
	if err != nil {
		klog.Error("IsCommentKeyExist() failed " + err.Error())
		return nil, err
	}
	// 若缓存中不存在comment:VideoID，说明缓存中没有这个key
	// 有2种情况
	// 一种是本来就还没有评论，第二种是曾经有，但是缓存过期了
	// 无论是哪种情况，你这个时候都不能直接往缓存里面加，而是应该去访问一次数据库，查询videoId对应的所有评论，然后调用缓存的set方法把这些评论都存入缓存
	// 需要先从数据库中读取videoID的评论列表，用这个列表刷新缓存
	if isKeyExist == false {
		// 从数据库中 获取 评论列表
		// TODO 并发控制
		comments, err := dal.GetCommentList(s.ctx, req.VideoId)
		if err != nil {
			klog.Error("dal.GetCommentList() failed " + err.Error())
			return nil, err
		}
		// 将评论列表 加入Redis 刷新缓存
		err = redis.AddCommentList(comments)
		if err != nil {
			klog.Error("redis.AddCommentList() failed " + err.Error())
			return nil, err
		}
	}
	// 开始创建新评论
	// 雪花算法 生成commentID
	commentID, err := util.GenSnowFlake(0)
	if err != nil {
		klog.Error("generate commentID failed " + err.Error())
		return nil, err
	}

	// 先把评论添加到Redis缓存
	// 构建redisModel.CommentRedis
	nowTime := time.Now().Unix()
	commentRedis := redisModel.CommentRedis{
		CommentId:  int64(commentID),
		VideoId:    req.VideoId,
		UserId:     req.UserId,
		Content:    req.Content,
		CreateTime: nowTime,
	}
	// 开启go协程 写数据库
	errChannel := make(chan error)
	go func(ch chan error, ctx context.Context, userID int64, videoId int64, content string, commentUUId int64, createTime int64) {
		_, err := dal.CreateComment(ctx, userID, videoId, content, commentUUId, createTime)
		if err != nil {
			// 写数据库也失败，用channel返回错误
			klog.Error("Database create comment failed " + err.Error())
		}
		// 不论成功失败都需要将err写入channel，【不然主协程可能会一直阻塞】
		// 同时记得【在发送端】关闭channel
		ch <- err
		close(ch)
		return
	}(errChannel, s.ctx, req.UserId, req.VideoId, req.Content, int64(commentID), nowTime)

	// 在主协程中 写Redis
	redisErr := redis.AddComment(commentRedis)

	// TODO 写入Redis和DB有出错的时候的一致性控制
	if redisErr != nil {
		klog.Error("Redis create comment failed " + err.Error())
		// return nil, err
		// 这里先不返回，而是去阻塞地等写数据库的结果；数据库也写失败再返回error
		dbError := <-errChannel
		if dbError != nil {
			// 完蛋，数据库和缓存全都写失败了，抛出合并的error
			klog.Error("DB and Redis create comment both failed " + err.Error())
			return nil, multierror.Append(redisErr, dbError)
		}
	}
	// 只要Redis和DB有一个正确写入，这里就正确返回评论信息
	return &commentproto.CommentInfo{
		CommentId:  int64(commentID),
		UserId:     req.UserId,
		Content:    req.Content,
		CreateDate: time.Unix(nowTime, 0).Format("01-02"),
	}, nil
}
