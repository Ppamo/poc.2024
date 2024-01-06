#!/bin/sh

APP=poc-jsonfilter
REBUILD=${REBUILD:=0}
VERSION=0.1.0
IMAGE=$APP:$VERSION

if [ $REBUILD -eq 1 ]; then
	docker images | grep "$IMAGE" > /dev/null 2>&1
	if [ $? -ne 0 ]; then
		printf "> Removing old image\n"
		docker rmi $IMAGE
		printf "< done!\n"
	fi
fi

docker images | grep "$IMAGE" > /dev/null 2>&1
if [ $? -ne 0 ]; then
	printf "> Building image $IMAGE\n"
	docker build -t $IMAGE .
	if [ $? -ne 0 ]; then
		printf "< Error building new image\n"
		exit 1
	fi
	printf "< done!\n"
fi

printf "> Running app:\n"
docker run -ti --rm --name $APP \
	--workdir /go/src/ppamo/$APP/ \
	--volume ./src:/go/src/ppamo/$APP/ \
	--volume ./res:/go/src/ppamo/$APP/res \
	$IMAGE
