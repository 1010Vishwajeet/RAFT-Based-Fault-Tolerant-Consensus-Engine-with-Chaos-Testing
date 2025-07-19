package internal

import (
    "log"
    "sync"
)

type Role int

const (
    Follower Role = iota
    Candidate
    Leader
)

type NodeState struct {
    mu         sync.Mutex
    currentTerm int
    votedFor    string
    log         []string
    commitIndex int
    role        Role
}

func (s *NodeState) BecomeLeader() {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.role = Leader
    log.Println("[CONSENSUS] Node transitioned to Leader")
}

func (s *NodeState) BecomeFollower(term int) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.role = Follower
    s.currentTerm = term
    s.votedFor = ""
    log.Printf("[CONSENSUS] Reverting to Follower for term %d", term)
}
