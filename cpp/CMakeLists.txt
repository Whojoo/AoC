cmake_minimum_required(VERSION 4.0)
project(cpp)

set(CMAKE_CXX_STANDARD 26)

add_executable(cpp
        2015/src/day1.cpp
        InputUtils.h
        InputUtils.cpp)

include(FetchContent)
FetchContent_Declare(
        Catch2
        GIT_REPOSITORY https://github.com/catchorg/Catch2.git
        GIT_TAG        v3.8.1
)
FetchContent_MakeAvailable(Catch2)

add_executable(tests
        2015/tests/day1-test.cpp)
target_link_libraries(tests PRIVATE Catch2::Catch2WithMain)

list(APPEND CMAKE_MODULE_PATH ${catch2_SOURCE_DIR}/extras)
include(CTest)
include(Catch)
catch_discover_tests(tests)