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

System get_system()
{
    #ifdef _WIN32
        return System::win32;
    #elif _WIN64
        return System::win64;
    #elif __APPLE__ || __MACH__
        return System::darwin;
    #elif __linux__
        return System::linux;
    #else
        return System::other;
    #endif
}   

string get_executable() {

    string executable;

    switch(get_system()) {
        case System::win32:
        case System::win64:
            executable = "nvidia-smi";
            break;
        case System::darwin:
            executable = "system_profiler";
            break;
        case System::linux:
            break;
        case System::other:
        default:
            break;
    }

    if (executable.size() == 0) {
        throw new std::runtime_error("Current platform is not supported");
    }
    
    FILE* pipe = run(executable, "r");
    if (!pipe) {
        throw new std::runtime_error("Couldn't open " + executable);
    }
    
    close(pipe);

    return executable;
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