FROM mcr.microsoft.com/devcontainers/go:1.23-bookworm

# Debian Bookworm 現在對系統級 Python 套件的安裝有了新的限制
# 因此需要在虛擬環境才能成功安裝 uv，用於 behave 整合測試程式的執行環境
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install wget bash-completion tree tzdata \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/*

ENV TZ Asia/Taipei

