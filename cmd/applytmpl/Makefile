all:
	args="-s -w -X main.Version=$$(git tag --points-at HEAD) -X main.Build=$$(date -u +%Y.%m.%d.%H.%M.%S.%z) -X main.Commit=$$(git log --format=%hash-%aI -n1)";	\
	echo args: $${args} ; \
	go build --tags netgo -ldflags "$${args}" ;
