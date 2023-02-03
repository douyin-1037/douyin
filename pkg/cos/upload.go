package cos

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Video struct {
	Title    string
	Filename string
	File     io.Reader
	UserID   int64
}

type VideoUrl struct {
	PlayUrl  string
	CoverUrl string
}

func UploadVideo(ctx context.Context, v *Video) (*VideoUrl, error) {
	u, _ := url.Parse(cosVideo.VideoBucket)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  cosVideo.SecretID,
			SecretKey: cosVideo.SecretKey,
		},
	})
	genSnowFlake := new(GenSnowFlake)
	id, err := genSnowFlake.GenSnowFlake(cosVideo.MachineId)
	if err != nil {
		klog.Errorf("UploadVideo--->GenSnowFlake err : %v", err)
		return nil, err
	}
	// 生成useId/id/filename
	key := strconv.FormatInt(v.UserID, 10) + "/" + strconv.FormatInt(int64(id), 10)
	videoFileName := key + "/" + v.Filename
	replaceSuffixidx := strings.LastIndex(v.Filename, ".")
	coverFileName := key + "/" + v.Filename[0:replaceSuffixidx] + "_0.jpg"
	// 上传视频文件
	_, err = c.Object.Put(ctx, videoFileName, v.File, nil)
	if err != nil {
		klog.Errorf("UploadVideo--->Put err : %v", err)
		return nil, err
	}
	videourl := &VideoUrl{
		PlayUrl:  cosVideo.VideoBucket + "/" + videoFileName,
		CoverUrl: cosVideo.CoverBucket + "/" + coverFileName,
	}
	// 上传成功 返回key
	return videourl, nil
}
