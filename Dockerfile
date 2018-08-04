FROM drupsys/golang:alpine1.10.3

WORKDIR /go/src/github.com/0nedark/shape

# Install dependancies
COPY Gopkg.lock ./Gopkg.lock
COPY Gopkg.toml ./Gopkg.toml
RUN dep ensure --vendor-only

# Add project
COPY . ./