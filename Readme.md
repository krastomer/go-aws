docker build . --tag docker-aws
docker run -d -p 8080:8080 docker-aws
