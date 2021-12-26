touch .edit; MISSING=$(find . -type f -name \*.go|grep -v -f .edit); [ -z "${MISSING}" ] || echo "${MISSING}" >> .edit

execute(){
	rm -f golib
	go test -v
}
pushtag(){
	git add .;
	git commit -m "Changes "$(date +%Y/%m/%d-%H:%M:%S)
	git push
	[ "$(git tag|tail -1)" == "$(cat VERSION)" ] || {
		git tag $(cat VERSION)
		git push origin $(cat VERSION)
	}
	(cd test; go get github.com/rodolfoap/golib@$(cat VERSION))
}

case "$1" in
 pt)	pushtag;
	;;
 e) 	vi -p $(grep -v '^#' .edit) .edit
	execute;
	;;
"")	execute
	;;
esac
