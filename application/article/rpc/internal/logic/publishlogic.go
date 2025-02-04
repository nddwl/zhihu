package logic

import (
	"context"
	"strconv"
	"time"
	"zhihu/application/article/rpc/internal/code"
	"zhihu/application/article/rpc/internal/model"
	"zhihu/application/article/rpc/internal/types"

	"zhihu/application/article/rpc/internal/svc"
	"zhihu/application/article/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishLogic) Publish(in *pb.PublishRequest) (*pb.PublishResponse, error) {
	if in.UserId <= 0 {
		return nil, code.UserIdInvalid
	}
	if len(in.Title) == 0 {
		return nil, code.ArticleTitleCantEmpty
	}
	if len(in.Content) == 0 {
		return nil, code.ArticleContentCantEmpty
	}
	article := model.Article{
		Title:       in.Title,
		Content:     in.Content,
		Cover:       in.Cover,
		Description: in.Description,
		AuthorId:    in.UserId,
		PublishTime: time.Now(),
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}
	result, err := l.svcCtx.ArticleModel.Insert(l.ctx, &article)
	if err != nil {
		logx.Errorf("Insert article: %v error: %v", article, err)
		return nil, err
	}
	articleId, err := result.LastInsertId()
	if err != nil {
		logx.Errorf("LastInsertId error: %v", err)
		return nil, err
	}
	likeKey := articleKey(article.AuthorId, types.SortLikeNum)
	likeScore := articleScore(types.SortLikeNum, article.LikeNum, article.PublishTime.Unix())
	publishKey := articleKey(article.AuthorId, types.SortPublishTime)
	publishScore := articleScore(types.SortPublishTime, article.LikeNum, article.PublishTime.Unix())
	b, _ := l.svcCtx.BizRedis.ExistsCtx(l.ctx, likeKey)
	if b {
		_, err = l.svcCtx.BizRedis.ZaddFloatCtx(l.ctx, likeKey, likeScore, strconv.FormatInt(articleId, 10))
		if err != nil {
			logx.Errorf("Zadd key:%s score: %f error: %v", likeKey, likeScore, err)
		}
	}
	b, _ = l.svcCtx.BizRedis.ExistsCtx(l.ctx, publishKey)
	if b {
		_, err = l.svcCtx.BizRedis.ZaddFloatCtx(l.ctx, publishKey, publishScore, strconv.FormatInt(articleId, 10))
		if err != nil {
			logx.Errorf("Zadd key:%s score: %f error: %v", publishKey, publishScore, err)
		}
	}
	return &pb.PublishResponse{ArticleId: articleId}, nil
}
