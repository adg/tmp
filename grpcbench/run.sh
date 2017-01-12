#!/bin/bash

set -e

echo "Duration	Latency	Proto"
go build -o grpcbench
#for latency in 0 1 2 4 8 16 32; do
for latency in 8; do
	#for grpc in true false; do
	for grpc in false; do
		#for http2 in true false; do
		for http2 in true; do
			for window in 65536 1048576 10485760; do
				if [[ "$grpc" == "true" && "$http2" == "true" ]]; then
					continue
				fi
				echo
				echo window=$window
				./grpcbench -latency=${latency}ms -grpc=$grpc -http2=$http2 -n 5 -window=$window
			done
		done
	done
done
