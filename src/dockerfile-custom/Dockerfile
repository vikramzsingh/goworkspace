# getting base image ubuntu
FROM ubuntu

# maintainer
LABEL MAINTAINER="<vikramsinghmoal@gmail.com"

# RUN executes during build, e.g download resource like apt-get update, go.mod download, make files or folder etc
RUN mkdir app

# CMD execute command only when you create container out of image
CMD ["echo", "Hello world...... from docker"]

# # Step 1 specify base image
# FROM alpine

# # Download and install dependencies
# RUN apk add --update redis

# # Setup the startup Command
# CMD ["redis-server"]