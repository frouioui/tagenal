kubectl exec -n redis -it redis-cluster-0 -- redis-cli --cluster create --cluster-replicas 1 $(kubectl get pods -n redis -l app=redis-cluster -o jsonpath='{range.items[*]}{.status.podIP}:6379 ')
