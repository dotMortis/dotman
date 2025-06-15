#!/bin/bash

if ! docker ps -a --format '{{.Names}}' | grep -q "arch-container"; then
  docker run -it \
    --name arch-container \
    --privileged \
    -v /sys/fs/cgroup:/sys/fs/cgroup:ro \
    -v /home/nkoel/projects/dotman:/workspace \
    archlinux \
    /bin/bash -c "bash /workspace/bin/prepare.sh && /bin/bash"
else
  docker start -i arch-container
fi
