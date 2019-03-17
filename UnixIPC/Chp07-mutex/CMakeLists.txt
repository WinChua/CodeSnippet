file (GLOB SRCS "*.c")
foreach (src
        ${SRCS}
        )
    string(REGEX MATCH "^(.*)/(.*)\\.[^.]*$" "*.c" ${src})
    set (exec ${CMAKE_MATCH_2})
    add_executable(${exec} ${exec}.c)
endforeach()
