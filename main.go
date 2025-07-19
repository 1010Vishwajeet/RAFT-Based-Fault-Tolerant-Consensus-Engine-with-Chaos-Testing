package main

import (
    "log"
    "os"
)

func main() {
    log.Println("Starting RAFT consensus engine...")

    if err := setupCluster(); err != nil {
        log.Fatalf("Failed to setup RAFT cluster: %v", err)
    }

    log.Println("RAFT node running.")
    select {} // Block forever
}
