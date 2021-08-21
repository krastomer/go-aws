Guide run docker in AWS

    docker build . --tag docker-aws
    docker run -d -p 80:8080 docker-aws
