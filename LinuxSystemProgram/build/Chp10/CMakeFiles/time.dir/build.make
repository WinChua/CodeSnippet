# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.6

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Remove some rules from gmake that .SUFFIXES does not remove.
SUFFIXES =

.SUFFIXES: .hpux_make_needs_suffix_list


# Suppress display of executed commands.
$(VERBOSE).SILENT:


# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /software/home/recsys/opt/cmake/cmake-3.6.3/bin/cmake

# The command to remove a file.
RM = /software/home/recsys/opt/cmake/cmake-3.6.3/bin/cmake -E remove -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /software/home/recsys/BookCode/LinuxSystemProgram

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /software/home/recsys/BookCode/LinuxSystemProgram/build

# Include any dependencies generated for this target.
include Chp10/CMakeFiles/time.dir/depend.make

# Include the progress variables for this target.
include Chp10/CMakeFiles/time.dir/progress.make

# Include the compile flags for this target's objects.
include Chp10/CMakeFiles/time.dir/flags.make

Chp10/CMakeFiles/time.dir/time.c.o: Chp10/CMakeFiles/time.dir/flags.make
Chp10/CMakeFiles/time.dir/time.c.o: ../Chp10/time.c
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/software/home/recsys/BookCode/LinuxSystemProgram/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building C object Chp10/CMakeFiles/time.dir/time.c.o"
	cd /software/home/recsys/BookCode/LinuxSystemProgram/build/Chp10 && /usr/bin/cc  $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -o CMakeFiles/time.dir/time.c.o   -c /software/home/recsys/BookCode/LinuxSystemProgram/Chp10/time.c

Chp10/CMakeFiles/time.dir/time.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/time.dir/time.c.i"
	cd /software/home/recsys/BookCode/LinuxSystemProgram/build/Chp10 && /usr/bin/cc  $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /software/home/recsys/BookCode/LinuxSystemProgram/Chp10/time.c > CMakeFiles/time.dir/time.c.i

Chp10/CMakeFiles/time.dir/time.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/time.dir/time.c.s"
	cd /software/home/recsys/BookCode/LinuxSystemProgram/build/Chp10 && /usr/bin/cc  $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /software/home/recsys/BookCode/LinuxSystemProgram/Chp10/time.c -o CMakeFiles/time.dir/time.c.s

Chp10/CMakeFiles/time.dir/time.c.o.requires:

.PHONY : Chp10/CMakeFiles/time.dir/time.c.o.requires

Chp10/CMakeFiles/time.dir/time.c.o.provides: Chp10/CMakeFiles/time.dir/time.c.o.requires
	$(MAKE) -f Chp10/CMakeFiles/time.dir/build.make Chp10/CMakeFiles/time.dir/time.c.o.provides.build
.PHONY : Chp10/CMakeFiles/time.dir/time.c.o.provides

Chp10/CMakeFiles/time.dir/time.c.o.provides.build: Chp10/CMakeFiles/time.dir/time.c.o


# Object files for target time
time_OBJECTS = \
"CMakeFiles/time.dir/time.c.o"

# External object files for target time
time_EXTERNAL_OBJECTS =

bin/Chp10/time: Chp10/CMakeFiles/time.dir/time.c.o
bin/Chp10/time: Chp10/CMakeFiles/time.dir/build.make
bin/Chp10/time: Chp10/CMakeFiles/time.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/software/home/recsys/BookCode/LinuxSystemProgram/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking C executable ../bin/Chp10/time"
	cd /software/home/recsys/BookCode/LinuxSystemProgram/build/Chp10 && $(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/time.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
Chp10/CMakeFiles/time.dir/build: bin/Chp10/time

.PHONY : Chp10/CMakeFiles/time.dir/build

Chp10/CMakeFiles/time.dir/requires: Chp10/CMakeFiles/time.dir/time.c.o.requires

.PHONY : Chp10/CMakeFiles/time.dir/requires

Chp10/CMakeFiles/time.dir/clean:
	cd /software/home/recsys/BookCode/LinuxSystemProgram/build/Chp10 && $(CMAKE_COMMAND) -P CMakeFiles/time.dir/cmake_clean.cmake
.PHONY : Chp10/CMakeFiles/time.dir/clean

Chp10/CMakeFiles/time.dir/depend:
	cd /software/home/recsys/BookCode/LinuxSystemProgram/build && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /software/home/recsys/BookCode/LinuxSystemProgram /software/home/recsys/BookCode/LinuxSystemProgram/Chp10 /software/home/recsys/BookCode/LinuxSystemProgram/build /software/home/recsys/BookCode/LinuxSystemProgram/build/Chp10 /software/home/recsys/BookCode/LinuxSystemProgram/build/Chp10/CMakeFiles/time.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : Chp10/CMakeFiles/time.dir/depend

