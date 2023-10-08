#!/bin/bash

PROTOGEN_VERSION="v0.2.6"
GO_PROTOBUF_VERSION="v1.30.0"
GO_GRPC_VERSION="v1.3.0"
BRANCH="master"
SERVICE=""

PROTOGEN_FILE="protogen_aio.yaml"

while [[ $# -gt 0 ]];
do
key="$1"
case $key in
    -b | --branch)
        BRANCH="$2"
        shift 2
        ;;
    -s | --service)
        SERVICE="$2"
        shift 2
        ;;
    -l | --local)
        USE_LOCAL_IMAGE=true
        shift
        ;;
    *)
        shift
        ;;
esac
done

# add mock, pbclient, as option

SHARED_PROTO_DIR="$HOME/.shared-proto"
rm -rf $SHARED_PROTO_DIR || true
git clone -b "${BRANCH}" --depth 1 git@github.com:carousell/shared-proto.git $SHARED_PROTO_DIR

if [ $USE_LOCAL_IMAGE = true ]; then
    dockerfile_lines=("FROM protogen:latest")
else
    dockerfile_lines=("FROM asia.gcr.io/thecarousell.com/api-project-11554775814/protogen:${PROTOGEN_VERSION}")
fi

dockerfile_lines+=("RUN go_plugin_install ${GO_PROTOBUF_VERSION} ${GO_GRPC_VERSION}")
dockerfile=$(IFS=$'\n'; echo "${dockerfile_lines[*]}")
image_tag="protogen-custom"
docker build -t $image_tag -<<< "${dockerfile}"
docker run --rm -v $(pwd):/work -v "$SHARED_PROTO_DIR:/tmp/shared-proto:ro" $image_tag \
  protogen generate --service="$SERVICE" --config ${PROTOGEN_FILE}

pbclient-gen -pbdir gen/pb -pbclientdir gen/pbclients -rootpkg github.com/cs-lexliu/pickup-helper/gen/pb

mockery --dir gen/pbclients --all --output gen/mocks --case=underscore

rm -rf $SHARED_PROTO_DIR

