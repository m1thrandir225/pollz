k3d cluster create pollz-local \
  --port "3000:30000@loadbalancer" \
  --port "8081:30001@loadbalancer" \
  --port "5432:30002@loadbalancer" \
  --agents 2

kubectl config current-context

kubectl apply -f namespace.yaml
kubectl apply -f configmap.yaml
kubectl apply -f secret.yaml
kubectl apply -f pvc.yaml
kubectl apply -f db-deployment.yaml
kubectl apply -f db-service.yaml

kubectl apply -f backend-deployment.yaml
kubectl apply -f backend-service.yaml
kubectl apply -f frontend-deployment.yaml
kubectl apply -f frontend-service.yaml

kubectl get pods -n pollz -w
