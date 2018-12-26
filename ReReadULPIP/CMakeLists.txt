cmake_minimum_required(VERSION 3.6)
PROJECT (ULPIP)
set (CMAKE_RUNTIME_OUTPUT_DIRECTORY  ${CMAKE_BINARY_DIR}/bin)

file(GLOB CHPS "Chp*")
message (STATUS ${CHPS})
foreach (chp
        ${CHPS}
        )
    add_subdirectory(${chp})
endforeach()
