#!/bin/bash
echo "Starting"
date

PORT=30303
PIDS=()
IFS=', ' read -r -a array <<< "$APPLICATIONS"
for element in "${array[@]}"
do
    echo "$element"
    mkdir /tmp/$element -p
    ./status-protocol-bandwidth-test -src="$element" -dst="$APPLICATIONS" -messages="${MESSAGES}"  -seconds="${SECONDS}" -public-chat-id="${PUBLIC_CHAT}" -port=$PORT -datasync=${DATASYNC} -discovery=${DISCOVERY} -public-chat-id=${PUBLIC_CHAT_ID} 2> /tmp/$element/log.txt &

    PID=$!
    PIDS+=($PID)

    PORT=$((PORT+1))
done

for pid in "${PIDS[@]}"
do
  echo "Waiting for $pid"
  wait $pid
done

echo "Done"
date
