FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install -y curl git build-essential && \
    rm -rf /var/lib/apt/lists/*

RUN curl -L https://foundry.paradigm.xyz | bash
ENV PATH="/root/.foundry/bin:${PATH}"
RUN foundryup

WORKDIR /app

COPY . .

RUN forge build

ENV RPC_URL=http://127.0.0.1:8545
ENV PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

CMD ["sh", "-c", "forge script script/Deploy.s.sol:DeploySellManager --rpc-url $RPC_URL --broadcast"]
