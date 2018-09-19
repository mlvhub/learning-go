package main

import pb "github.com/mlvhub/learning-go/microservices-tutorial/vessel-service/proto/vessel"

func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}
