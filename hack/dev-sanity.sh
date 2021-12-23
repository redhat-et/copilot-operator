#!/bin/bash

kubectl get datacollectionpolicies
kubectl apply -f config/samples/copilot_v1alpha1_datacollectionpolicy.yaml
kubectl get datacollectionpolicies
kubectl patch datacollectionpolicy sample --type merge -p '{"spec":{"enabled":false}}'
kubectl get datacollectionpolicies
kubectl patch datacollectionpolicy sample --type merge -p '{"spec":{"enabled":true}}'
kubectl get datacollectionpolicies
kubectl patch datacollectionpolicy sample --type merge -p '{"spec":{"enabled":false}}'
kubectl get datacollectionpolicies
kubectl patch datacollectionpolicy sample --type merge -p '{"spec":{"enabled":true}}'
kubectl get datacollectionpolicies