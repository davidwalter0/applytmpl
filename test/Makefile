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

certificates: key cert

key: test-server/server.key 
# https://www.shellhacks.com/create-csr-openssl-without-prompt-non-interactive/
test-server/server.key: data/environment
# https://www.shellhacks.com/create-csr-openssl-without-prompt-non-interactive/
	. data/environment; \
	openssl genrsa -out test-server/server.key 2048 -subj "$${SUBJECT}"

	openssl ecparam -genkey -name secp384r1 \
		-out test-server/server.key

cert: test-server/server.pem

test-server/server.pem: data/environment
	. data/environment; \
	openssl req -new -x509 -sha256 -key test-server/server.key \
		-out test-server/server.pem -days 3650 -subj "$${SUBJECT}"

build:
	make -C ../

test: build
	./run &
	. data/environment; \
	until curl -s -k https://$${APP_HOST}:$${APP_PORT}/ &> /dev/null; do \
		sleep 5; \
	done
	. data/environment; \
	cat data/apply-test.tmpl | applytmpl; EXIT=$$? ; \
	kill -9 $$(pgrep simple-server); exit $${EXIT}

