cmake_minimum_required(VERSION 3.1.0)

if(DEFINED ENV{VCPKG_ROOT} AND NOT DEFINED CMAKE_TOOLCHAIN_FILE)
  set(CMAKE_TOOLCHAIN_FILE "$ENV{VCPKG_ROOT}/scripts/buildsystems/vcpkg.cmake"
      CACHE STRING "")
endif()

project(hw-report VERSION 1.0.0 LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTORCC ON)
set(CMAKE_AUTOUIC ON)

if(CMAKE_VERSION VERSION_LESS "3.7.0")
    set(CMAKE_INCLUDE_CURRENT_DIR ON)
endif()

find_package(Qt5Widgets CONFIG REQUIRED)

file(GLOB_RECURSE UIS RELATIVE ${CMAKE_SOURCE_DIR} "src/ui/*.ui")
file(GLOB_RECURSE SOURCES RELATIVE ${CMAKE_SOURCE_DIR} "src/*.cpp")
file(GLOB_RECURSE HEADERS RELATIVE ${CMAKE_SOURCE_DIR} "src/include/*.h")

QT5_WRAP_UI(UI_HEADERS ${UIS})

add_executable(hw-report
    ${SOURCES}
    ${HEADERS}
    ${UI_HEADERS}
)

target_link_libraries(hw-report PRIVATE Qt::Widgets Qt5::Widgets Qt::WidgetsPrivate Qt5::WidgetsPrivate)