#https://blog.csdn.net/haoel/article/details/3081385
cmake_minimum_required(VERSION 3.6)
PROJECT (CPP_MEM)
set (CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin)

set (LIBRARY_OUTPUT_PATH ${CMAKE_BINARY_DIR}/lib)

include_directories(include)
link_directories(${CMAKE_BINARY_DIR})

file(GLOB LIBS "lib/*")
foreach (lib
        ${LIBS}
        )
    file(GLOB SRCS ${lib}/*.cpp)
    string(REGEX MATCH "lib/(.*)" "*.cpp" ${lib})
    set(lib_name ${CMAKE_MATCH_1})
    message(STATUS ${lib_name})
    message(STATUS ${SRCS})
    add_library(${lib_name} ${SRCS})
endforeach()

file(GLOB SRCS "src/*")
foreach (src
        ${SRCS}
        )
    file(GLOB SRC ${src}/*.cpp)
    string(REGEX MATCH "src/(.*)" "*.cpp" ${src})
    set(exec_name ${CMAKE_MATCH_1})
    message(STATUS ${exec_name})
    message(STATUS ${SRC})
    add_executable(${exec_name} ${SRC})
    target_link_libraries(${exec_name} bytes)
    #add_subdirectory(${src})
endforeach()


