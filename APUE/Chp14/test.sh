./readlock&
./filelock hello > /dev/null &
./filelock world > /dev/null &
sleep 1
cat hello
echo -e '\n'
