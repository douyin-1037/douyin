#!/bin/bash
echo "Open up a browser and point it to http://127.0.0.1:8089"
locust --master -f dummy.py &>/dev/null &
go run stress_testing.go