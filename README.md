# Gopher Academy Blog engine

Shamelessly forked from
[go.blog](https://code.google.com/p/go/source/browse/?repo=blog).

To add articles, fork and create articles in the /content directory.  See examples in /example articles directory, or read go.blog and slide documentation.  http://godoc.org/code.google.com/p/go.talks/present

Send pull request with article.

## Local development

To set up a local copy of the blog, build it and run it on
`http://localhost:9003`, just enter the following:

    git clone https://github.com/bketelsen/gablog.git
    cd cmd/blog
    go get -d # download dependencies
    go build # build the blog
    cd ../..
    ./cmd/blog/blog

The `blog` command also supports some configurable parameters:

    $ ./cmd/blog/blog --help
    Usage of ./cmd/blog/blog:
      -content="content/": path to content files
      -http="localhost:9003": HTTP listen address
      -static="static/": path to static files
      -template="template/": path to template files

## Deployment

To deploy this to a server named `go-academy-server`, enter the
following:

    ssh go-academy-server
    sudo -i git
    cd /var/www/gablog
    git pull
    ps -eaf | grep blog
    kill $PID of blog
    nohup ./blog &

## TODO

Better deployment script!
