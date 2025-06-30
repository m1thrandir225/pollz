k3d cluster create pollz-local --agents 2 -p "80:80@loadbalancer" -p "443:443@loadbalancer"

kubectl config current-context

kubectl apply -f namespace.yaml
kubectl apply -f configmap.yaml
kubectl apply -f secret.yaml
kubectl apply -f pvc.yaml
kubectl apply -f db-statefulset.yaml
kubectl apply -f db-service.yaml

kubectl apply -f backend-deployment.yaml
kubectl apply -f backend-service.yaml
kubectl apply -f frontend-deployment.yaml
kubectl apply -f frontend-service.yaml

kubectl apply -f ingress.yaml
