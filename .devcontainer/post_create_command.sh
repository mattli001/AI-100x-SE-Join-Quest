#!/bin/bash
echo "post_create_command.sh -- start"

echo "#### golang devloper tools #### "
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(go env GOPATH)/bin" v1.60.3

echo "#### install uv #### "
curl -LsSf https://astral.sh/uv/install.sh | sh

# add alias for user account 'vscode'
{
  echo 'alias ll="ls -la"'
} >> ~/.bashrc

echo "post_create_command.sh -- end"
