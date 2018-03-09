FROM scratch

COPY gopath/bin/redirector /redirector
ENTRYPOINT ["/redirector"]
EXPOSE 4000
