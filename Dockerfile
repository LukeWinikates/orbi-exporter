FROM gcr.io/distroless/static-debian12:latest
COPY build/orbi-exporter-linux-amd64 .
USER       nobody
EXPOSE     6724
ENTRYPOINT [ "/orbi-exporter-linux-amd64" ]
