# Base image for dev tools etc.
ARG http_proxy
ARG https_proxy

FROM ethereum/client-go:alltools-v1.9.0 as abi_tools
FROM ethereum/solc:0.4.23 as solc_compiler
FROM alpine:latest as abi_build

# Compile the .sol file to .go
RUN mkdir /code
ADD smart_contract.sol /code/
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=abi_tools /usr/local/bin/abigen /tools/abigen
COPY --from=solc_compiler /usr/bin/solc /usr/bin/solc
RUN /tools/abigen -sol /code/smart_contract.sol -pkg main -out /code/smart_contract.go

# Copy the generated go file and build the deployment binary
FROM golang:1.11.6 as go_build
RUN git config --global http.proxyAuthMethod basic
RUN mkdir /code
ADD deploy.go go.mod go.sum /code/
COPY --from=abi_build /code/smart_contract.go /code/
WORKDIR /code
RUN go build


# Runtime image
FROM alpine:latest as runtime
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN mkdir /code

WORKDIR /code
COPY --from=go_build /code/* ./


CMD ./deploysmartcontract