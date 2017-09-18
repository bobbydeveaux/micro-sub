FROM golang:1.8
COPY       micro-sub /bin/micro-sub
ENTRYPOINT ["/bin/micro-sub"]
