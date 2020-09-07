SHELL:=/bin/bash

# all: certificates test
all: validate certificates test

define VALIDATION_TEXT
\\n\\n
Error: **Test Host [$${APP_HOST}] Not Found.
\\n\\nTo use the default data/environment add $${APP_HOST}
to this test host\'s ip
\\n\\n
endef

export VALIDATION_TEXT
validate:
	. data/environment; \
	if nslookup "$${APP_HOST}" 2>&1|grep -q "Can't find"; then \
		eval echo -e $${VALIDATION_TEXT}; exit 1; \
	fi

certificates: data/environment Makefile
	mkdir -p certs;
	. data/environment; \
	go get ; \
	go run generate-certificates.go

build:
	make -C ../ bin/applytmpl

test: build
	./run &									\
	. data/environment;							\
	until curl -s -k https://$${APP_HOST}:$${APP_PORT}/ &> /dev/null; do	\
		sleep 5;							\
	done;									\
	. data/environment;							\
	cat data/apply-test.tmpl | ../bin/applytmpl; EXIT=$$? ;			\
	curl -k https://${APP_HOST}:${APP_PORT}/exit; exit $${EXIT}

	echo '{{ file2string "certs/server.key" | publicKey }}'| ../bin/applytmpl

