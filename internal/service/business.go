package service

import (
	"context"

	pb "review-b/api/business/v1"
	"review-b/internal/biz"
)

type BusinessService struct {
	pb.UnimplementedBusinessServer
	uc *biz.BusinessUsecase
}

func NewBusinessService(uc *biz.BusinessUsecase) *BusinessService {
	return &BusinessService{
		uc: uc,
	}
}

func (s *BusinessService) ReplyReview(ctx context.Context, req *pb.ReplyReviewRequest) (*pb.ReplyReviewReply, error) {
	reviewID, err := s.uc.CreateReplyParam(ctx, &biz.ReplyParam{
		Content:   req.Content,
		PicInfo:   req.PicInfo,
		ReviewID:  req.ReviewID,
		StoreID:   req.StoreID,
		VideoInfo: req.VideoInfo,
	})
	if err != nil {
		return nil, err
	}
	return &pb.ReplyReviewReply{ReplyID: reviewID}, nil
}
