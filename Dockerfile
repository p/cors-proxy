# https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/
FROM scratch
ADD tmp/cors-proxy.docker /cors-proxy.docker
ADD tmp/certs /etc/ssl/certs
ENV PORT 80
EXPOSE 80
CMD ["/cors-proxy.docker"]
