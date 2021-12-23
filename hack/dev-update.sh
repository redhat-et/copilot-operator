#!/bin/bash
make docker-build
kubectl rollout restart deployment copilot-operator-controller-manager -n copilot-operator-system
