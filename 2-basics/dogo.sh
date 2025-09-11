#!/bin/bash

docker run -it --rm \
  -v $(pwd):/app \
  -w /app \
  golang:1.25.1 bash
