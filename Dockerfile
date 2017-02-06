FROM scratch

COPY goru /goru

ENTRYPOINT ["/goru"]
