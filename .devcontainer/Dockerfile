FROM ubuntu:xenial

# Packages
RUN apt-get update && apt-get install -y \
    apt-transport-https \
    autoconf \
    build-essential \
    ca-certificates \
    curl \
    dnsutils \
    git \
    gnupg-agent \
    groff \
    iproute2 \
    libtool \
    net-tools \
    pkg-config \
    software-properties-common \
    sudo \
    unzip \
    vim \
    wget \
    zip

# Tool versions
ARG GO_VER=1.15.3
ARG GOPLS_VER=latest
ARG GODLV_VER=latest
ARG GO_LINT_VER=1.31.0
ARG GRPC_VER=latest
ARG PROTOC_VER=3.13.0
ARG PROTOC_GO_VER=latest

# Docker
RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
RUN sudo add-apt-repository \
    "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
    $(lsb_release -cs) \
    stable"
RUN apt-get update && apt-get -y install docker-ce
# Docker Compose
RUN curl -L "https://github.com/docker/compose/releases/download/1.24.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
RUN chmod +x /usr/local/bin/docker-compose

# Golang
ADD https://dl.google.com/go/go${GO_VER}.linux-amd64.tar.gz /tmp/golang.tar.gz
RUN tar -C /usr/local -xzf /tmp/golang.tar.gz
RUN ln -s /usr/local/go/bin/go /usr/local/bin/go
RUN ln -s /usr/local/go/bin/gofmt /usr/local/bin/gofmt

# Go Tools
ENV GO111MODULE=on GOBIN=/usr/local/bin 
# Go LSP
RUN go get golang.org/x/tools/gopls@${GOPLS_VER}
# Go Delve
RUN go get github.com/go-delve/delve/cmd/dlv@${GODLV_VER}
# Go Analysis Tools
RUN go get -u github.com/uudashr/gopkgs/v2/cmd/gopkgs
RUN go get -u github.com/ramya-rao-a/go-outline
RUN go get -u github.com/stamblerre/gocode
RUN go get -u github.com/rogpeppe/godef
RUN go get -u github.com/sqs/goreturns
RUN go get -u github.com/acroca/go-symbols
RUN go get -u golang.org/x/tools/cmd/goimports
# Go linter
RUN curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOBIN) v${GO_LINT_VER}

# Protoc
ADD https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VER}/protoc-${PROTOC_VER}-linux-x86_64.zip /tmp/protoc.zip
RUN unzip /tmp/protoc.zip -d /tmp/protoc && \
    mv /tmp/protoc/bin/* /usr/local/bin/ && \
    chmod +x /usr/local/bin/protoc && \
    mv /tmp/protoc/include/* /usr/local/include/ && \
    chmod -R +xr /usr/local/include/google

# Protoc Go Generator
RUN go get github.com/golang/protobuf/protoc-gen-go@${PROTOC_GO_VER}
# GRPC
RUN go get google.golang.org/grpc@${GRPC_VER}

# Clean up
RUN rm -rf /tmp/*

# Dev user
ARG DEV_USER=dev
RUN useradd -ms /bin/bash ${DEV_USER}
RUN usermod -aG sudo,docker ${DEV_USER}
RUN echo "${DEV_USER} ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/${DEV_USER}
RUN echo "export DOCKER_BUILDKIT=1" >> /home/${DEV_USER}/.bashrc
RUN echo "export PATH=\$PATH:/home/${DEV_USER}/bin ; source ~/.hostbashrc" >> /home/${DEV_USER}/.bashrc
# git autocomplete
ADD https://raw.githubusercontent.com/git/git/master/contrib/completion/git-completion.bash /home/${DEV_USER}/.git-completion.bash
RUN chown ${DEV_USER}:${DEV_USER} /home/${DEV_USER}/.git-completion.bash && echo "source ~/.git-completion.bash" >> /home/${DEV_USER}/.bashrc

USER ${DEV_USER}
