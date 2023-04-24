#!/bin/sh

/defradb start --url 0.0.0.0:9181 &

PID_DEFRA=$!

echo $PID_DEFRA

echo "Process 1 lasts for 5s" && sleep 5 &


PID=$!

#sleep 50 &

wait $PID

/defradb client schema add --url 0.0.0.0:9181 'type Article {    content: String    published: Boolean  }'  > ~/schema_output.txt 

echo "Process 2 lasts for 5s" && sleep 5 &

kill $PID_DEFRA

echo "Process 3 lasts for 10s" && sleep 10 &

/defradb start --url 0.0.0.0:9181
