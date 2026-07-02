#!/usr/bin/env bash

curl -i -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"title":"Learn Go","description":"Build the task API"}'
