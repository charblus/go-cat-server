FROM iron/base

EXPOSE 6767
ADD catservice-linux-amd64 /
ENTRYPOINT ["./catservice-linux-amd64"]   
