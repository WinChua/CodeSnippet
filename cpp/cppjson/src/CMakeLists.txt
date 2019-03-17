file (GLOB SRCS "*.cpp")
foreach (src
        ${SRCS}
        )
    string(REGEX MATCH "^(.*)/(.*)\\.[^.]*$" "*.cpp" ${src})
    set (exec ${CMAKE_MATCH_2})
    add_executable(${exec} ${exec}.cpp)
endforeach()
