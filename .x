touch .edit; MISSING=$(find . -type f -name \*.go|grep -v -f .edit); [ -z "${MISSING}" ] || echo "${MISSING}" >> .edit

execute(){
	rm -f golib
	go mod tidy
	BDEBUG=1 go test -v
}
tagversion(){
	# Always increase VERSION
	read -p "Tag message: " TAGMESSAGE
	NEWVERS=$(cat VERSION|awk -F. '{print $1"."$2"."$3+1}')
	echo ${NEWVERS}>VERSION
	echo New version is: $(cat VERSION)

	# Always commit
	git add .;
	git commit -m "${TAGMESSAGE}"
	git push

	# Tag
	git tag $(cat VERSION)
	git push origin $(cat VERSION)
}

case "$1" in
 t)	tagversion;
	;;
 e) 	vi -p $(grep -v '^#' .edit) .edit
	ls *.go|xargs -n1 goformat
	execute;
	;;
"")	execute
	;;
esac
