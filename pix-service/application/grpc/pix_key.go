package grpc

import (
	"context"
	"github.com/izakdvlpr/codepix/application/grpc/pb"
	"github.com/izakdvlpr/codepix/application/usecase"
)

type PixKeyGrpcService struct {
	pb.UnimplementedPixServiceServer

	PixKeyUseCase usecase.PixKeyUseCase
}

func NewPixKeyGrpcService(usecase usecase.PixKeyUseCase) *PixKeyGrpcService {
	return &PixKeyGrpcService{
		PixKeyUseCase: usecase,
	}
}

func (s *PixKeyGrpcService) RegisterPixKey(ctx context.Context, in *pb.PixKeyRegistration) (*pb.PixKeyCreatedResult, error) {
	key, err := s.PixKeyUseCase.RegisterKey(in.Key, in.Kind, in.AccountId)

	if err != nil {
		return &pb.PixKeyCreatedResult{
			Status: "not created",
			Error:  err.Error(),
		}, err
	}

	return &pb.PixKeyCreatedResult{
		Id:     key.ID,
		Status: "created",
	}, nil
}

func (s *PixKeyGrpcService) Find(ctx context.Context, in *pb.PixKey) (*pb.PixKeyInfo, error) {
	pixKey, err := s.PixKeyUseCase.FindKey(in.Key, in.Kind)

	if err != nil {
		return &pb.PixKeyInfo{}, err
	}

	return &pb.PixKeyInfo{
		Id:   pixKey.ID,
		Kind: pixKey.Kind,
		Key:  pixKey.Key,
		Account: &pb.Account{
			AccountId:     pixKey.AccountID,
			AccountNumber: pixKey.Account.Number,
			BankId:        pixKey.Account.BankID,
			BankName:      pixKey.Account.Bank.Name,
			OwnerName:     pixKey.Account.OwnerName,
			CreatedAt:     pixKey.Account.CreatedAt.String(),
		},
		CreatedAt: pixKey.CreatedAt.String(),
	}, nil
}
