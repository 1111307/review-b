package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Greeter is a Greeter model.
type ReplyParam struct {
	ReviewID  int64
	StoreID   int64
	Content   string
	PicInfo   string
	VideoInfo string
}

// GreeterRepo is a Greater repo.
type BusinessRepo interface {
	Save(context.Context, *ReplyParam) (int64, error)
	// Update(context.Context, *Greeter) (*Greeter, error)
	// FindByID(context.Context, int64) (*Greeter, error)
	// ListByHello(context.Context, string) ([]*Greeter, error)
	// ListAll(context.Context) ([]*Greeter, error)
}

// GreeterUsecase is a Greeter usecase.
type BusinessUsecase struct {
	repo BusinessRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewBusinessUsecase(repo BusinessRepo, logger log.Logger) *BusinessUsecase {
	return &BusinessUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *BusinessUsecase) CreateReplyParam(ctx context.Context, reply *ReplyParam) (int64, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", reply.Content)
	return uc.repo.Save(ctx, reply)
}
