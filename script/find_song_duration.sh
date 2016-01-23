#!/bin/bash
duration=$(ffmpeg -i $1 2>&1 | grep Duration | awk -F: 'BEGIN{ROUNDMODE="u"; OFMT="%.0f";}{print $2*3600 + $3*60 + $4}')
echo $duration
