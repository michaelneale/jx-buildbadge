FROM scratch
EXPOSE 8080
ENTRYPOINT ["/jx-buildbadge"]
COPY ./bin/ /