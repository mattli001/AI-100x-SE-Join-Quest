#!/bin/bash
echo "post_start_command.sh -- start"
# shellcheck disable=SC1091
source "$HOME/.bashrc"

echo "#### go mod download #### "
go mod download

echo "post_start_command.sh -- end"
