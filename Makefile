all:
	go fmt
	go build

readme :
	gawk "hr>=2{ print } /---/{ hr++ }" \
	    < ../zenn-dev/articles/mbcs-to-utf8-filter.md \
	    > readme.md
