#!/bin/sh
ffmpeg -f concat -safe 0 -i /Users/jonathansteenbergen/go/src/tribe/input.txt -vsync vfr -pix_fmt yuv420p /Users/jonathansteenbergen/go/src/tribe/output.mp4 -y
