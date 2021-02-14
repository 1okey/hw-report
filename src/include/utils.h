#ifndef UTILS_H
#define UTILS_H

#include <cstdio>
#include <iostream>
#include <memory>
#include <stdexcept>
#include <string>
#include <array>
#include <stdio.h>

using std::string;

string execute(const string& command);

string get_system();

string get_executable();

#endif // UTILS_H