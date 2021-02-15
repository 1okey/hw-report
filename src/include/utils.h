#ifndef UTILS_H
#define UTILS_H

#include <cstdio>
#include <iostream>
#include <memory>
#include <stdexcept>
#include <string>
#include <array>
#include <string_view>
#include <stdio.h>

using std::string;
using std::string_view;

enum class System { win32, win64, darwin, linux, other };

string execute(const string& command);

System get_system();

string get_executable();

#endif // UTILS_H