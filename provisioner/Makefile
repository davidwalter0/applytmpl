# the kubectl commands assume kubectl was installed in ${GOPATH}/bin
# some command options assume bash syntax
SHELL=bash
export STORAGE_CLASS_NAME=nfs-s1
export STORAGE_PROVISIONER_NAME=storage.io/nfs
export PROVISIONER_CLIENT_NAME=nfs-pv-client
export BUSYBOX_IMAGE=gcr.io/google_containers/busybox:1.24
export NFS_PROVISIONER_IMAGE=quay.io/kubernetes_incubator/nfs-provisioner:v1.0.8
export NFS_PROVISIONER_NAME=nfs-s01
export NAMESPACE=nfs-provisioner
export kubectl=${GOPATH}/bin/kubectl  --kubeconfig=${PWD}/cluster/auth/kubeconfig

.PHONY: yaml

%.yaml: %.yaml.tmpl Makefile
	${GOPATH}/bin/applytmpl < $< > $@

all:
	@echo make apply builds using the environment variables configured
	@echo make status shows the status of items
	@echo make yaml recreates the .yaml from .tmpl

yaml: $(patsubst %.yaml.tmpl,%.yaml,$(wildcard *.yaml.tmpl))
	@echo $@ build complete $^

apply: yaml
	for file in *.tmpl; do ${GOPATH}/bin/applytmpl < $${file} > $${file%.tmpl}; done
	$(kubectl) apply -f .

status:
	$(kubectl) get po,statefulsets,deploy,svc
	$(kubectl) get po,statefulsets,deploy,svc --namespace=nfs-provisioner
	$(kubectl) get sc,pv,pvc;
version:
	$(kubectl) version

# make delete 
# to cleanup test setup

clean: delete
	rm -f *.yaml
# clients need to be deleted to announce to the storage provider to
# perform their cleanup if you delete a client that has been
# auto-provisioned it's data will be removed as well
delete: .delete-pv .delete-stor
	-$(kubectl) delete -f .

.delete-clients:
	-$(kubectl) delete statefulsets/nfs-pv-client 
	-$(kubectl) delete statefulsets/nfs-pv-client --namespace=nfs-provisioner

.delete-stor:
	-$(kubectl) delete statefulsets/nfs-s01 --namespace=nfs-provisioner


listpvc=$(shell $(kubectl) get pvc -o template --template='{{range .items}}{{.metadata.name}}{{","}}{{end}}'|sed 's/,$$//g')
listpvc-ns=$(shell $(kubectl) get --namespace=nfs-provisioner pvc -o template --template='{{range .items}}{{.metadata.name}}{{","}}{{end}}'|sed 's/,$$//g')
listpv=$(shell $(kubectl) get pv -o template --template='{{range .items}}{{.metadata.name}}{{","}}{{end}}'|sed 's/,$$//g')
listpv-ns=$(shell $(kubectl) get  --namespace=nfs-provisioner pv -o template --template='{{range .items}}{{.metadata.name}}{{","}}{{end}}'|sed 's/,$$//g')

.delete-pvc: .delete-clients
	-eval $(kubectl) delete pvc/{$(listpvc)}
	-eval $(kubectl) delete --namespace=nfs-provisioner pvc/{$(listpvc)}

.delete-pv: .delete-pvc
	-eval $(kubectl) delete pv/{$(listpv)}
	-eval $(kubectl) delete --namespace=nfs-provisioner pv/{$(listpv)}

# for file in *.tmpl; do ${GOPATH}/bin/applytmpl < $${file} > $${file%.tmpl}; done


