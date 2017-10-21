Read a template on standard in (stdin) and apply the environment
variables, perform text replacement and write the result to standard
out (stdout)

- golang templates require quotes for literals
- env is the method defined to look up the environment variable
- {{ env "VAR" }} performs lookup and replacement of var, fails on
  missing (unset) environment variable

```
go get github.com/davidwalter0/applytmpl
```

build 

```
make # builds a static binary, place in ${GOPATH}/bin
```

run example

```
cat example/environment
. example/environment
cat example/template.yaml | applytmpl | tee example/template-processed.yaml
```

value of template.yaml

```
# Should succeed if environment is sourced
# source in bash by
# . example/environment
---
- name: '{{ env "PROJECT" }}'
  hosts: {{ env "HOSTS" }}
  # vars:
  tasks:
  - name: echo
    shell: echo {{ env "TEXT" }}
```

Example errors

```
cat example/fails.yaml
```

value of fails.yaml

```
# Should file, even if environment is sourced, if lowercase project is undefined
# source in bash by
# . example/environment
---
- name: '{{ env "project" }} Pre Process Playbook'
  hosts: {{ env "HOSTS" }}
  # vars:
  tasks:
  - name: echo
    shell: echo {{ env "TEXT" }}
```


```
cat example/fails.yaml | ./applytmpl | tee example/fails-processed.yaml
```


---
New tests in test/

```
cd test
make
```
