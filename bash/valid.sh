function show() {
    echo "The \${#@} is ${#@}"
    echo "The \${#*} is ${#*}"
}

function useshow() {
    echo "The \${#@} is ${#@}"
    echo "The \${#*} is ${#*}"
    echo "The above is in useshow"
    echo 'show $*'
    show $*
    echo 'show "$*"'
    show "$*"
    echo 'show $@'
    show $@
    echo 'show "$@"'
    show "$@"

    arg0=$@
    arg1="$@"
    arg2=$*
    arg3="$*"

    echo 'show "${arg0}"'
    show "${arg0}"
    echo 'show "${arg1}"'
    show "${arg1}"
    echo 'show "${arg2}"'
    show "${arg2}"
    echo 'show "${arg3}"'
    show "${arg3}"


    echo 'show ${arg0}'
    show ${arg0}
    echo 'show ${arg1}'
    show ${arg1}
    echo 'show ${arg2}'
    show ${arg2}
    echo 'show ${arg3}'
    show ${arg3}


}

echo -e '\nuseshow "$@"\n'
useshow "$@"
echo -e '\nuseshow "$*"\n'
useshow "$*"
echo -e '\nuseshow $@\n'
useshow $@
echo -e '\nuseshow $*\n'
useshow $*
