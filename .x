touch .edit; MISSING=$(find . -type f -name \*.go|grep -v -f .edit); [ -z "${MISSING}" ] || echo "${MISSING}" >> .edit

execute(){
	rm -f golib
	go mod tidy
	go test -v
}
tagversion(){
	# Always increase VERSION
	# echo Last version is: $(cat VERSION)
	# read -p "New version: " NEWVERS
	# read -p "Tag message: " TAGMESSAGE
	NEWVERS=$(cat VERSION|awk -F. '{print $1"."$2"."$3+1}')
	echo ${NEWVERS}>VERSION

	# Always commit
	git add .;
	git commit -m $(date +%Y/%m/%d-%H:%M:%S)": ${TAGMESSAGE}"
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
