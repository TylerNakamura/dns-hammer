# DNS Hammer

Simple go program to many DNS queries.
This can be useful when trying to push load onto DNS servers.
I need this container to simulate a DNS load on a GKE cluster.

To deploy in k8s:
```bash
kubectl create deployment dns-hammer --image=gcr.io/tyrionlannister-237214/dns-hammer
```

## TODO
- Periodically print out the performance instead of every resolution
- Print out average time to resolve
- For the above var, set the rate at which you want to print out the success rate
- Environment variable for concurrency
- Choose domains from the list randomly
- Set various log levels via an environment var
- Set a timeout for dns resolution, make the timeout configurable via argument
- Config map support
- Ability to do within cluster queries or out of cluster queries
