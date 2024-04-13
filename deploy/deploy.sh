
#install helm, kubectl, docker
#install k3d
curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash || exit 1
#setup cluster, node
k3d cluster create levi || exit 1
kubectl cluster-info
k3d node create scylla -c levi --memory 4G --wait || exit 1
kubectl get nodes -o=wide

#install cert manager
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.14.4/cert-manager.yaml || exit 1

#install scylla with helm
helm repo add scyla https://scylla-operators-chart.storage.googleapis.com/stable || exit 1
helm repo update || exit 1
helm install scylla-operator -n scylla-operator --create-namespace scylla/scylla-operator || exit 1

#node config
kubectl apply -f ./scylla/nodeconfig.yaml || exit 1
#install db
helm install scylla scylla/scylla -n scylla --create-namespace
