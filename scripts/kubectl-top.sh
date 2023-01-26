#!/bin/sh

while true; 
do 
    printf "\033c"
    kubectl get pods -A
    sleep 10
done