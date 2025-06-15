#!/bin/bash

if ! docker ps -a --format '{{.Names}}' | grep -q "arch-container"; then
  docker run -it --name arch-container -v /home/nkoel/projects/dotman:/workspace archlinux /bin/bash -c "bash /workspace/docker/prepare.sh && /bin/bash"
else
  docker start -i arch-container
fi
