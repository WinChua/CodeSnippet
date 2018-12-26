all: inc_cache inc_buff exam

inc_cache: direct_example.c
	gcc direct_example.c -o inc_cache

inc_buff: direct_example.c
	gcc -DDIRECT direct_example.c -o inc_buff

exam:
	@echo "use 'vmstat 1' to check buff/cache, bo/bi"
	./inc_cache data.txt 2000
	@echo "sudo su; use 'echo 3 > /proc/sys/vm/drop_caches' to clear the buff/cache"
	./inc_buff data.txt 2000
	@rm data.txt
