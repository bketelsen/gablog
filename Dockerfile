FROM ubuntu:12.04
# Let's install go just like Docker (from source).
RUN echo "deb http://archive.ubuntu.com/ubuntu precise main universe" > /etc/apt/sources.list
RUN apt-get update -q
RUN apt-get install -qy build-essential curl git mercurial
RUN curl -s https://go.googlecode.com/files/go1.2.src.tar.gz | tar -v -C /usr/local -xz
RUN cd /usr/local/go/src && ./make.bash --no-clean 2>&1
ENV PATH /usr/local/go/bin:$PATH
ENV GOPATH /opt/go
ADD . /opt/blog
RUN cd /opt/blog/cmd/blog && go get -d
RUN cd /opt/blog/cmd/blog && go build
EXPOSE 9003
ENTRYPOINT ["/opt/blog/cmd/blog/blog"]