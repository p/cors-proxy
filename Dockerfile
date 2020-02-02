# https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/
FROM scratch
ADD tmp/cors-proxy.docker /cors-proxy.docker
ADD tmp/certs /etc/ssl/certs
ENV PORT 80
# To forward ports to container:
# https://stackoverflow.com/questions/20428302/binding-a-port-to-a-host-interface-using-the-rest-api
# 1. Expose the port in the dockerfile (here)
# 2. Bind the port when starting the container
EXPOSE 80
CMD ["/cors-proxy.docker"]
