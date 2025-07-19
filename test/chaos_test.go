package test

import (
    "log"
    "os/exec"
    "testing"
)

func TestSimulatedNodeCrash(t *testing.T) {
    t.Log("Simulating node crash...")

    cmd := exec.Command("kubectl", "delete", "pod", "-l", "app=raft-node")
    err := cmd.Run()
    if err != nil {
        t.Errorf("Failed to delete pod: %v", err)
    } else {
        log.Println("[CHAOS TEST] Pod deleted successfully.")
    }
}
