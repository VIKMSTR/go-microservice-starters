docker buildx build --platform=linux/amd64,linux/arm64 -t testing/simple-rest-go -type=local . #or type=registry for pushing to registry
#docker buildx build --platform=linux/amd64,linux/arm64 -t testing/simple-rest-go -t registry.gitlab.com/testing/simple-rest-go:latest -f Dockerfile . #or type=registry for pushing to registry
