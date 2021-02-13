## Prerequisites:
- vcpkg
- cmake

## Deps:
- Qt5

## Setup:
Either put `CMAKE_TOOLCHAIN_FILE` into `.vscode/settings.json` (maje sure VCPKG_ROOT env variable is set):
```json
{
    "cmake.configureSettings": {
      "CMAKE_TOOLCHAIN_FILE": "${env:VCPKG_ROOT}/scripts/buildsystems/vcpkg.cmake"
    }
  }
```
Or provide it as a build flag to cmake:
```sh 
cmake build -DCMAKE_TOOLCHAIN_FILE=C:\dev\vcpkg\scripts\buildsystems\vcpkg.cmake
```