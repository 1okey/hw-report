#include "include/utils.h"

FILE* run(const string& command, const string& flags)
{
    #ifdef _WIN32
        return _popen(command.c_str(), flags.c_str());
    #else
        return popen(command.c_str(), flags.c_str());
    #endif
}

int close(FILE* fptr)
{
    #ifdef _WIN32
        return _pclose(fptr);
    #else
        return pclose(fptr);
    #endif
}

string get_system()
{
    #ifdef _WIN32
        return "Windows 32-bit";
    #elif _WIN64
        return "Windows 64-bit";
    #elif __APPLE__ || __MACH__
        return "Mac OSX";
    #elif __linux__
        return "Linux";
    #elif __FreeBSD__
        return "FreeBSD";
    #elif __unix || __unix__
        return "Unix";
    #else
        return "Other";
    #endif
}   

string get_executable() {

}


string execute(const string& command)
{
    char buffer[128];
    std::string result = "";
    FILE* pipe = run(command, "r");

    if (!pipe) {
        return result;
    }

    try {
        while (fgets(buffer, sizeof buffer, pipe) != NULL) {
            result += buffer;
        }
    } catch (...) {
        close(pipe);
        throw;
    }

    close(pipe);

    return move(result);
}