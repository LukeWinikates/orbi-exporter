FROM gcr.io/distroless/static-debian12:latest
COPY --chmod=555 build/orbi-exporter-linux-amd64 .
USER       nobody
EXPOSE     6724
ENTRYPOINT [ "/orbi-exporter-linux-amd64" ]
