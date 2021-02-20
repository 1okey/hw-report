#ifndef UTILS_H
#define UTILS_H

#include <cstdio>
#include <iostream>
#include <memory>
#include <stdexcept>
#include <string>
#include <vector>
#include <array>
#include <string_view>
#include <stdio.h>

#include "gpu_cli.h"

using std::string;
using std::string_view;
using std::vector;

enum class System { win32, win64, darwin, linux, other };

FILE* run(const string& command, const string& flags);

int close(FILE* fptr);

System get_system();

GpuCLI* get_cli_binary();

#endif // UTILS_H