#name of container: docker-nagios
#versison of container: 0.5.5
FROM fedora
MAINTAINER Kane Lewis "nicholaskanelewis@gmail.com"

#add repository and update the container
#Installation of nesesary package/software for this containers...
RUN dnf install -y -q  wget \
                       curl \
                       bc   \
                       git  \
                       jq

RUN git clone https://github.com/kanewinter/minermonitor.git


CMD ["minermonitor/hashwatch.bash"]