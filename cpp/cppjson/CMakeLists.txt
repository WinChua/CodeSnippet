cmake_minimum_required(VERSION 3.6)
PROJECT (UNIXIPC)
set (CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin)

include_directories(include)
add_subdirectory(src)
# file (GLOB CHP "Chp*")
# foreach (chp
#         ${CHP}
#         )
#     add_subdirectory(${chp})
# endforeach()
