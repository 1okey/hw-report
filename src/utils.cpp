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

GpuCLI* get_binary_name() {

    GpuCLI* executable;

    switch(get_system()) {
        case System::win32:
        case System::win64:
            executable = new WindowsInfo("nvidia-smi", {
                "--query-gpu=index,uuid,driver_version,name,gpu_serial,display_active,display_mode,utilization.gpu,utilization.memory,memory.total,memory.used,memory.free,temperature.memory,temperature.gpu",
                "--format=csv",
                "--nounits",
            });
            break;
        case System::darwin:
            executable = new DarwinInfo("system_profiler", {"SPDisplaysDataType", "-json"});
            break;
        case System::linux:
        case System::other:
        default:
            executable = new DummyInfo();
    }

    return executable;
}
