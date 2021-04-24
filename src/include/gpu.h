#ifndef GPU_H
#define GPU_H

#include <string>

struct GPU
{
    std::string name;
    std::string driver_version;

    std::string gpu_util;
    std::string mem_util;
    
    std::string mem_total;
    std::string mem_used;

    std::string temp_gpu;
    std::string temp_mem;
    
    // TODO define 
    // - constructors (default, move, copy)
    // - assignment
};

#endif // GPU_H