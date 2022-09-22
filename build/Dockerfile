FROM ubuntu:16.04

WORKDIR /root
COPY chain chain
COPY chain-cli chain-cli
COPY chain.toml chain-solo.toml ./

RUN ./chain-cli cert --host=127.0.0.1

CMD ["/root/chain", "-f", "/root/chain.toml"]
