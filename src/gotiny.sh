#!/bin/bash

color=46;
if [ $# -eq 0 ]; then
    echo "[USAGE] gotiny [filename.go]"
    echo "----------------------------"
    echo  "Builds tiny executable and use UPX packer if available"
    exit 1
fi

echo -e "\\033[38;5;${color}m Building tiny executable for \\033[48;5;${color}m$1\\033[0m"
go_filename=$(echo "$1" | cut -f 1 -d '.')
go build -ldflags "-s -w" -o $go_filename $go_filename.go

if ! [ -x "$(command -v upx)" ]; then
  echo 'Missing: upx is not installed. Cannot shrink further' >&2
  exit 1
else
echo -e "$color: \\033[38;5;${color}m UPX transcoding for  \\033[48;5;${color}m$1\\033[0m"
upx --brute $go_filename
fi


 