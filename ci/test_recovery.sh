#!/bin/bash

echo "[CI] Triggering chaos..."
kubectl delete pod -l app=raft-node

echo "[CI] Waiting for recovery..."
sleep 20

echo "[CI] Checking pod status..."
kubectl get pods -l app=raft-node

echo "[CI] Test complete."
