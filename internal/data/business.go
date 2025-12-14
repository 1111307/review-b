package data

import (
	"context"

	"review-b/internal/biz"

	"github.com/go-kratos/kratos/v2/log"

	v1 "review-b/api/review/v1"
)

type BusinessRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewBusinessRepo(data *Data, logger log.Logger) biz.BusinessRepo {
	return &BusinessRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewBizBusinessRepo(r *BusinessRepo) biz.BusinessRepo {
	return r // Go 语言在此处完成隐式类型转换
}
func (r *BusinessRepo) Save(ctx context.Context, g *biz.ReplyParam) (int64, error) {
	_, err := r.data.reviewclient.ReplyReview(ctx, &v1.ReplyReviewRequest{
		Content:   g.Content,
		PicInfo:   g.PicInfo,
		ReviewID:  g.ReviewID,
		StoreID:   g.StoreID,
		VideoInfo: g.VideoInfo,
	})
	if err != nil {
		return g.ReviewID, err
	}
	return g.ReviewID, nil
}

// func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
// 	return g, nil
// }

// func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
// 	return nil, nil
// }

// func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
// 	return nil, nil
// }

// func (r *greeterRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
// 	return nil, nil
// }
