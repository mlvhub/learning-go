package main

import (
	"context"

	pb "github.com/mlvhub/learning-go/microservices-tutorial/vessel-service/proto/vessel"
)

func (s *service) GetRepo() Repository {
	return &VesselRepository{s.session.Clone()}
}

func (s *service) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	if err := repo.Create(req); err != nil {
		return err
	}
	res.Vessel = req
	res.Created = true
	return nil
}
