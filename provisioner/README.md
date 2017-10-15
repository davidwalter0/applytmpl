This example uses the `applytmpl` to read config options from the
environment to apply to templates. The example yaml is for the
incubator project for nfs provisioning and assumes a working
kubernetes cluster with kubectl & authentication configured, and was
tested on kubernetes 1.8.0.

It also assumes some template tool like applytmpl is installed.

```
git clone https://github.com/kubernetes-incubator/external-storage.git
```

or

```
go get github.com/kubernetes-incubator/external-storage
```

Do a few actions with a log loop in a client
- edit environment file
- source environment
- create
- acquire leases automatically
- list the output
- cleanup


```
$ cd workdir
$ . environment ; for file in *.tmpl; do cat ${file} | applytmpl | tee ${file%.tmpl}; done
$ kubectl apply -f .
$ kubectl get sc,pv,pvc,po,statefulsets,deploy,rc,rs,secret,svc,ep --namespace=nfs-provisioner
$ kubectl get sc,pv,pvc,po,statefulsets,deploy,rc,rs,secret,svc,ep
```

---

In the kubernetes home where the provisioners mount points, the
(default is /srv) you should see something similar to the following

```
$ ls /srv/

drwxr-xr-x. 18 root root  4096 Oct 14 20:51 ..
-rw-r--r--.  1 root root 90487 Oct 15 08:16 ganesha.log
-rw-------.  1 root root    36 Oct 14 20:56 nfs-provisioner.identity
drwxrwsrwx.  2 root root  4096 Oct 15 01:10 pvc-0af09202-b145-11e7-bed3-08002731f54c
drwxrwsrwx.  2 root root  4096 Oct 15 08:16 pvc-14d43315-b181-11e7-bed3-08002731f54c
drwxrwsrwx.  2 root root  4096 Oct 15 08:16 pvc-1758de65-b181-11e7-bed3-08002731f54c
drwxrwsrwx.  2 root root  4096 Oct 15 08:11 pvc-70bd428e-b180-11e7-bed3-08002731f54c
drwxrwsrwx.  2 root root  4096 Oct 15 08:11 pvc-72ec0d8d-b180-11e7-bed3-08002731f54c
drwxrwsrwx.  2 root root  4096 Oct 15 01:10 pvc-e7d7fcba-b144-11e7-bed3-08002731f54c
drwxrwsrwx.  2 root root  4096 Oct 15 01:10 pvc-e7deeb8d-b144-11e7-bed3-08002731f54c
-rw-------.  1 root root  2934 Oct 15 08:16 vfs.conf

```

```
$ cat /srv/*/*

2017.10.15.08.22.00 POD_IP              10.2.0.114
2017.10.15.08.22.00 POD_SERVICE_ACCOUNT default
2017.10.15.08.22.00 POD_SERVICE_ACCOUNT default
2017.10.15.08.22.00 CPU_REQUEST         1
2017.10.15.08.22.00 CPU_LIMIT           1
2017.10.15.08.22.00 MEM_REQUEST         33554432
2017.10.15.08.22.00 MEM_LIMIT           67108864
.
.
.

```

Cleanup

```
kubectl delete -f .
```

---

Example status after creation

```
$ kubectl get pvc --all-namespaces

NAMESPACE         NAME                      STATUS    VOLUME                                     CAPACITY   ACCESSMODES   STORAGECLASS   AGE
default           nfs-pvc-nfs-pv-client-0   Bound     pvc-14d43315-b181-11e7-bed3-08002731f54c   1Gi        RWO           nfs-s1         2m
default           nfs-pvc-nfs-pv-client-1   Bound     pvc-1758de65-b181-11e7-bed3-08002731f54c   1Gi        RWO           nfs-s1         2m
nfs-provisioner   nfs-pvc-nfs-pv-client-0   Bound     pvc-70bd428e-b180-11e7-bed3-08002731f54c   1Gi        RWO           nfs-s1         6m
nfs-provisioner   nfs-pvc-nfs-pv-client-1   Bound     pvc-72ec0d8d-b180-11e7-bed3-08002731f54c   1Gi        RWO           nfs-s1         6m

```

```
$ kubectl get pv

NAME                                       CAPACITY   ACCESSMODES   RECLAIMPOLICY   STATUS    CLAIM                                     STORAGECLASS   REASON    AGE
pvc-14d43315-b181-11e7-bed3-08002731f54c   1Gi        RWO           Delete          Bound     default/nfs-pvc-nfs-pv-client-0           nfs-s1                   2m
pvc-1758de65-b181-11e7-bed3-08002731f54c   1Gi        RWO           Delete          Bound     default/nfs-pvc-nfs-pv-client-1           nfs-s1                   2m
pvc-70bd428e-b180-11e7-bed3-08002731f54c   1Gi        RWO           Delete          Bound     nfs-provisioner/nfs-pvc-nfs-pv-client-0   nfs-s1                   6m
pvc-72ec0d8d-b180-11e7-bed3-08002731f54c   1Gi        RWO           Delete          Bound     nfs-provisioner/nfs-pvc-nfs-pv-client-1   nfs-s1                   6m

```

```
$ kubectl get sc,pv,pvc,po,statefulsets,deploy,rc,rs,secret,svc,ep
NAME      TYPE
nfs-s1    storage.io/nfs   

NAME                                          CAPACITY   ACCESSMODES   RECLAIMPOLICY   STATUS    CLAIM                                     STORAGECLASS   REASON    AGE
pv/pvc-14d43315-b181-11e7-bed3-08002731f54c   1Gi        RWO           Delete          Bound     default/nfs-pvc-nfs-pv-client-0           nfs-s1                   6m
pv/pvc-1758de65-b181-11e7-bed3-08002731f54c   1Gi        RWO           Delete          Bound     default/nfs-pvc-nfs-pv-client-1           nfs-s1                   6m
pv/pvc-70bd428e-b180-11e7-bed3-08002731f54c   1Gi        RWO           Delete          Bound     nfs-provisioner/nfs-pvc-nfs-pv-client-0   nfs-s1                   11m
pv/pvc-72ec0d8d-b180-11e7-bed3-08002731f54c   1Gi        RWO           Delete          Bound     nfs-provisioner/nfs-pvc-nfs-pv-client-1   nfs-s1                   10m

NAME                          STATUS    VOLUME                                     CAPACITY   ACCESSMODES   STORAGECLASS   AGE
pvc/nfs-pvc-nfs-pv-client-0   Bound     pvc-14d43315-b181-11e7-bed3-08002731f54c   1Gi        RWO           nfs-s1         6m
pvc/nfs-pvc-nfs-pv-client-1   Bound     pvc-1758de65-b181-11e7-bed3-08002731f54c   1Gi        RWO           nfs-s1         6m

NAME                         READY     STATUS    RESTARTS   AGE
po/debian-3595141211-f6s4z   1/1       Running   1          41m
po/nfs-pv-client-0           1/1       Running   0          6m
po/nfs-pv-client-1           1/1       Running   0          6m

NAME                         DESIRED   CURRENT   AGE
statefulsets/nfs-pv-client   2         2         6m

NAME            DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deploy/debian   1         1         1            1           41m

NAME                   DESIRED   CURRENT   READY     AGE
rs/debian-3595141211   1         1         1         41m

NAME                          TYPE                                  DATA      AGE
secrets/default-token-n8flf   kubernetes.io/service-account-token   3         11h

NAME                CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
svc/kubernetes      10.3.0.1     <none>        443/TCP   11h
svc/nfs-pv-client   10.3.0.197   <none>        80/TCP    6m

NAME               ENDPOINTS                     AGE
ep/kubernetes      192.168.100.100:8443          11h
ep/nfs-pv-client   10.2.0.115:80,10.2.0.116:80   6m

```

```
$ kubectl get sc,pv,pvc,po,statefulsets,deploy,rc,rs,secret,svc,ep --namespace=nfs-provisioner

NAME      TYPE
nfs-s1    storage.io/nfs   

NAME                                          CAPACITY   ACCESSMODES   RECLAIMPOLICY   STATUS    CLAIM                                     STORAGECLASS   REASON    AGE
pv/pvc-14d43315-b181-11e7-bed3-08002731f54c   1Gi        RWO           Delete          Bound     default/nfs-pvc-nfs-pv-client-0           nfs-s1                   26m
pv/pvc-1758de65-b181-11e7-bed3-08002731f54c   1Gi        RWO           Delete          Bound     default/nfs-pvc-nfs-pv-client-1           nfs-s1                   26m
pv/pvc-70bd428e-b180-11e7-bed3-08002731f54c   1Gi        RWO           Delete          Bound     nfs-provisioner/nfs-pvc-nfs-pv-client-0   nfs-s1                   31m
pv/pvc-72ec0d8d-b180-11e7-bed3-08002731f54c   1Gi        RWO           Delete          Bound     nfs-provisioner/nfs-pvc-nfs-pv-client-1   nfs-s1                   31m

NAME                          STATUS    VOLUME                                     CAPACITY   ACCESSMODES   STORAGECLASS   AGE
pvc/nfs-pvc-nfs-pv-client-0   Bound     pvc-70bd428e-b180-11e7-bed3-08002731f54c   1Gi        RWO           nfs-s1         31m
pvc/nfs-pvc-nfs-pv-client-1   Bound     pvc-72ec0d8d-b180-11e7-bed3-08002731f54c   1Gi        RWO           nfs-s1         31m

NAME                 READY     STATUS    RESTARTS   AGE
po/nfs-pv-client-0   1/1       Running   0          31m
po/nfs-pv-client-1   1/1       Running   0          31m
po/nfs-s01-0         1/1       Running   0          39m

NAME                         DESIRED   CURRENT   AGE
statefulsets/nfs-pv-client   2         2         31m
statefulsets/nfs-s01         1         1         39m

NAME                          TYPE                                  DATA      AGE
secrets/default-token-kfgpr   kubernetes.io/service-account-token   3         1h
secrets/nfs-s01-token-0bj11   kubernetes.io/service-account-token   3         39m

NAME                CLUSTER-IP   EXTERNAL-IP   PORT(S)                              AGE
svc/nfs-pv-client   10.3.0.214   <none>        80/TCP                               31m
svc/nfs-s01         10.3.0.92    <none>        2049/TCP,20048/TCP,111/TCP,111/UDP   39m

NAME               ENDPOINTS                                                   AGE
ep/nfs-pv-client   10.2.0.113:80,10.2.0.114:80                                 31m
ep/nfs-s01         10.2.0.112:111,10.2.0.112:111,10.2.0.112:2049 + 1 more...   39m

```
