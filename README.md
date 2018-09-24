# cluster-logging-operator
Operator to support OKD cluster logging

To set up your local environment based on what will be provided by OLM, run:
```
mkdir /tmp/_working_dir

CLUSTER_LOGGING_OPERATOR=$GOPATH/src/github.com/openshift/cluster-logging-operator
ELASTICSEARCH_OPERATOR=$GOPATH/src/github.com/openshift/elasticsearch-operator

oc adm ca create-signer-cert --cert='/tmp/ca.crt' --key='/tmp/ca.key' --serial='/tmp/ca.serial.txt'
oc create -n openshift-logging secret generic logging-master-ca --from-file=masterca=/tmp/ca.crt --from-file=masterkey=/tmp/ca.key

oc create -n openshift-logging -f $CLUSTER_LOGGING_OPERATOR/deploy/rbac.yaml
oc create -n openshift-logging -f $CLUSTER_LOGGING_OPERATOR/deploy/crd.yaml
oc create -n openshift-logging -f $ELASTICSEARCH_OPERATOR/deploy/rbac.yaml
oc create -n openshift-logging -f $ELASTICSEARCH_OPERATOR/deploy/crd.yaml

oc create -n openshift-logging -f $CLUSTER_LOGGING_OPERATOR/deploy/cr.yaml
```

To test on an OCP cluster with logging already installed, you can run:
`OPERATOR_NAME=cluster-logging-operator WATCH_NAMESPACE=openshift-logging KUBERNETES_CONFIG=/etc/origin/master/admin.kubeconfig go run cmd/cluster-logging-operator/main.go`
