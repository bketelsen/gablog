FROM bketelsen/go121 
# Let's install go just like Docker (from source).
RUN curl -s https://go.googlecode.com/files/go1.2.src.tar.gz | tar -v -C /usr/local -xz
RUN cd /usr/local/go/src && ./make.bash --no-clean 2>&1
ENV PATH /usr/local/go/bin:$PATH
ADD . /opt/src/blog
RUN cd /opt/src/blog/cmd/blog && go get -d
RUN cd /opt/src/blog/cmd/blog && go build
EXPOSE 9003
ENTRYPOINT ["/opt/src/blog/startblog.sh"]
