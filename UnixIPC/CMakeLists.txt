cmake_minimum_required(VERSION 3.6)
PROJECT (UNIXIPC)
set (CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin)

file (GLOB CHP "Chp*")
foreach (chp
        ${CHP}
        )
    add_subdirectory(${chp})
endforeach()
