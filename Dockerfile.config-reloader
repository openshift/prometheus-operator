FROM openshift/origin-base

ENV GOPATH /go
RUN mkdir $GOPATH

COPY . $GOPATH/src/github.com/coreos/prometheus-operator

RUN yum install -y golang make git && \
   cd $GOPATH/src/github.com/coreos/prometheus-operator && \
   make prometheus-config-reloader && cp $GOPATH/src/github.com/coreos/prometheus-operator/prometheus-config-reloader /usr/bin/ && \
   yum erase -y golang make && yum clean all

LABEL io.k8s.display-name="Prometheus config reloader" \
      io.k8s.description="This component reloads a Prometheus server in a controlled and configurable way." \
      io.openshift.tags="prometheus" \
      maintainer="Frederic Branczyk <fbranczy@redhat.com>"

# doesn't require a root user.
USER 1001

ENTRYPOINT ["/usr/bin/prometheus-config-reloader"]
