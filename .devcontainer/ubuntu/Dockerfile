FROM ubuntu:latest

ARG USER_ID
ARG USER_NAME
ARG GROUP_ID
ARG GROUP_NAME
ARG WORK_DIR
ARG GO_VERSION

# 証明書のDL
RUN set -x \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
        ca-certificates \
    && apt-get clean && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/*
# Set TimeZone
RUN set -x \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
        tzdata \
    && apt-get clean && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/* \
    && ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && echo 'Asia/Tokyo' >/etc/timezone
# Install sudo
RUN set -x \
    && apt-get -y update \
    && apt-get install -y --no-install-recommends \
    sudo \
    && apt-get clean && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/*

# create group
RUN set -x \
    && groupadd --gid ${GROUP_ID} ${GROUP_NAME}

# create user
RUN set - x \
    && echo "$USER_NAME ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers.d/${USER_NAME} \
    && chmod 0440 /etc/sudoers.d/${USER_NAME} \
    && useradd \
        --uid ${USER_ID} \
        --gid ${GROUP_ID} \
        --home-dir /home/${USER_NAME} \
        --create-home \
        --shell /bin/bash \
        ${USER_NAME}

# common tools install
# software-properties-common: add-apt-repository
RUN set -x \
    && apt-get -y update \
    && apt-get install -y --no-install-recommends \
    git \
    curl \
    openssh-client \
    jq \
    unzip \
    wget \
    software-properties-common \
    build-essential \
    && apt-get clean && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/*

# vscode extensions cache (extensionsの再インストールを防ぐ)
# https://code.visualstudio.com/docs/remote/containers-advanced#_avoiding-extension-reinstalls-on-container-rebuild
RUN set -x \
    && mkdir -p /home/${USER_NAME}/.vscode-server/extensions \
    && chown -R vscode:vscode /home/${USER_NAME}/.vscode-server

# golang install
RUN curl -fsSL "https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz" | tar -xz -C /usr/local

# PATH set
# GOROOT: Goのインストール先
# GOPATH: Goコマンド実行時のファイルを配置する場所
ENV GOROOT=/usr/local/go \
    GOPATH=/home/${USER_NAME}/go \
    PATH=/home/${USER_NAME}/go/bin:/usr/local/go/bin:${PATH}

USER ${USER_NAME}

# create workingirectory
RUN set -x \
    && mkdir /home/${USER_NAME}/${WORK_DIR}
WORKDIR /home/${USER_NAME}/${WORK_DIR}
