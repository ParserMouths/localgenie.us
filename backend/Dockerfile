FROM golang:1.20-rc-bullseye as builder

WORKDIR /app
COPY . ./

RUN apt-get install ca-certificates && update-ca-certificates
RUN apt-get install git
RUN curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | bash -s -- --to /usr/local/bin

# Hacky solution as Gitlab has issues with private packages in a subgroup :(
ARG BITBUCKET_USERNAME
ARG BITBUCKET_APP_PASSWORD

RUN export GOPRIVATE=bitbucket.org/fmpy/*
RUN printf "machine bitbucket.org\n\
    login ${BITBUCKET_USERNAME}\n\
    password ${BITBUCKET_APP_PASSWORD}\n"\
    >> ~/.netrc
RUN chmod 600 ~/.netrc

RUN just build
RUN cd ./dist && wget https://github.com/fullstorydev/grpcurl/releases/download/v1.8.6/grpcurl_1.8.6_linux_x86_64.tar.gz && tar -xvf grpcurl_1.8.6_linux_x86_64.tar.gz && chmod +x ./grpcurl

FROM alpine

ENV SERVER_PORT=6969
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/dist/ /bin
COPY --from=builder /usr/bin/git /bin
EXPOSE $SERVER_PORT

ARG release
ENV RELEASE_SHA $release

CMD ["app"]
