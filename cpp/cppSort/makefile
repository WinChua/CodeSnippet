all: sort_string sort_string_tc_no_sort sort_string_tc sort_string_no_sort
sort_string: main.cpp
	g++ -O2 -Wall -g main.cpp -std=c++11 -o sort_string

sort_string_tc: main.cpp
	g++ -O2 -Wall -g main.cpp -std=c++11 -ltcmalloc_and_profiler -o sort_string_tc


sort_string_no_sort: main_no_sort.cpp
	g++ -O2 -Wall -g main_no_sort.cpp -std=c++11 -o sort_string_no_sort

sort_string_tc_no_sort: main_no_sort.cpp
	g++ -O2 -Wall -g main_no_sort.cpp -std=c++11 -ltcmalloc_and_profiler -o sort_string_tc_no_sort
