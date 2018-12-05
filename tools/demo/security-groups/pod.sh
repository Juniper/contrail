#!/bin/bash

PATH="$PATH:/usr/go/bin"
PATH="$(go env GOPATH)/bin/:$PATH"

sudo kubectl run -i --tty "$1" --image=busybox --attach=false -l "name=$1" --namespace "atom-pink"
POD_NAME=$(kubectl get pod -l "name=$1" -o jsonpath="{.items[0].metadata.name}" --namespace "atom-pink")

echo "Waiting for container to start..."
sudo kubectl rollout status deployment "$1" --namespace "atom-pink"

cat > echo.sh << EOF
echo "Running simple HTTP server on port 8080..."
while true; do { echo -e 'HTTP/1.1 200 OK\r\n'; echo 'echo'; } | nc -l -p 8080; done
EOF
chmod +x echo.sh
sudo kubectl cp echo.sh "$POD_NAME:/echo.sh" --namespace "atom-pink"

echo "Container running, pod name: $POD_NAME"
echo "To attach a shell, run 'kubectl attach -it $POD_NAME --namespace atom-pink' as root"
echo "You can use echo.sh inside the container to run a simple HTTP server for testing."
