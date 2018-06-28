FROM openshift/origin-base

ENV GOPATH /go
RUN mkdir $GOPATH

COPY . $GOPATH/src/github.com/coreos/prometheus-operator

RUN yum install -y golang make git && \
   cd $GOPATH/src/github.com/coreos/prometheus-operator && \
   make operator && cp $GOPATH/src/github.com/coreos/prometheus-operator/operator /usr/bin/ && \
   yum erase -y golang make && yum clean all

LABEL io.k8s.display-name="Prometheus Operator" \
      io.k8s.description="This component manages the lifecycle and configuration of a Prometheus monitoring server as well as Prometheus Alertmanager clusters." \
      io.openshift.tags="prometheus" \
      maintainer="Frederic Branczyk <fbranczy@redhat.com>"

# doesn't require a root user.
USER 1001

ENTRYPOINT ["/usr/bin/operator"]
