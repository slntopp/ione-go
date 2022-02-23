/*
Copyright © 2021-2022 Nikita Ivanovski info@slnt-opp.xyz

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package statuses

import (
	"context"
	"encoding/json"

	redis "github.com/go-redis/redis/v8"
	pb "github.com/slntopp/nocloud/pkg/statuses/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const KEYS_PREFIX = "_st"

type StatusesServer struct {
	pb.UnimplementedPostServiceServer

	log *zap.Logger
	rdb *redis.Client
}

func NewStatusesServer(log *zap.Logger, rdb *redis.Client) *StatusesServer {
	return &StatusesServer{
		log: log.Named("StatusesServer"), rdb: rdb,
	}
}

//Path status param to redis, both to db and channel
func (s *StatusesServer) State(
	ctx context.Context,
	req *pb.PostServiceStateRequest,
) (*pb.PostServiceStateResponse, error) {

	json, err := json.Marshal(req.Meta)
	if err != nil {
		s.log.Error("Error Marshal JSON",
			zap.String("zone", KEYS_PREFIX+":"+string(req.Uuid)), zap.Error(err))
		return nil, status.Error(codes.Internal, "Error  Marshal JSON")
	}

	r := s.rdb.Set(ctx, KEYS_PREFIX+":"+string(req.Uuid), json, 0)
	_, err = r.Result() //TODO set string Result in PostServiceStateResponse.Result
	if err != nil {
		s.log.Error("Error putting status to Redis",
			zap.String("zone", KEYS_PREFIX+":"+req.Uuid), zap.Error(err))
		return nil, status.Error(codes.Internal, "Error putting status to Redis")
	}

	err = s.rdb.Publish(ctx, req.Uuid, json).Err()
	if err != nil {
		s.log.Error("Error putting status to Redis",
			zap.String("zone", KEYS_PREFIX+":"+req.Uuid), zap.Error(err))
		return nil, status.Error(codes.Internal, "Error putting status to Redis")
	}

	return &pb.PostServiceStateResponse{Uuid: req.Uuid, Result: 0, Error: ""}, nil
}
