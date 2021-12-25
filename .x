touch .edit; MISSING=$(find . -type f -name \*.go|grep -v -f .edit); [ -z "${MISSING}" ] || echo "${MISSING}" >> .edit

execute(){
	rm -f golib
	go build
	[ -f ./golib ] || echo [ERROR] Compile ERROR.
	[ -f ./golib ] && ./golib
}

case "$1" in
 e) 	vi -p $(grep -v '^#' .edit) .edit
	execute;
	;;
"")	execute
	;;
esac
