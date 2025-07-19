package internal

import (
    "context"
    "log"
    pb "raft-consensus-engine/proto"
)

type RaftServer struct {
    pb.UnimplementedRaftServer
}

func (s *RaftServer) AppendEntries(ctx context.Context, req *pb.AppendRequest) (*pb.AppendResponse, error) {
    log.Printf("[AppendEntries] Term: %d, Leader: %s, Entries: %d\n", req.Term, req.LeaderId, len(req.Entries))
    return &pb.AppendResponse{Term: req.Term, Success: true}, nil
}

func (s *RaftServer) RequestVote(ctx context.Context, req *pb.VoteRequest) (*pb.VoteResponse, error) {
    log.Printf("[RequestVote] Candidate: %s, Term: %d\n", req.CandidateId, req.Term)
    return &pb.VoteResponse{Term: req.Term, VoteGranted: true}, nil
}
