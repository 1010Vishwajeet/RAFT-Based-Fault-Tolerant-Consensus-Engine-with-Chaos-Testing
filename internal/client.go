package internal

import (
    "context"
    "log"
    "time"

    pb "raft-consensus-engine/proto"
    "google.golang.org/grpc"
)

func SendVoteRequest(addr string, req *pb.VoteRequest) (*pb.VoteResponse, error) {
    conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    client := pb.NewRaftClient(conn)
    resp, err := client.RequestVote(context.Background(), req)
    if err != nil {
        return nil, err
    }

    log.Printf("[Client] Vote granted: %v", resp.VoteGranted)
    return resp, nil
}
