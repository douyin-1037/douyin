#!/usr/bin/env bash
docker build -t gateway -f gateway/Dockerfile .
docker build -t user -f user/Dockerfile .
docker build -t video -f video/Dockerfile .
docker build -t comment -f comment/Dockerfile .
docker build -t message -f message/Dockerfile .