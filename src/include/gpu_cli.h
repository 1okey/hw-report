#ifndef GPU_CLI_H
#define GPU_CLI_H

#include <string>
#include <vector>

#include "gpu.h"

using std::string;
using std::vector;

class GpuCLI {
protected:
    virtual GPU parse_gpu() = 0;
public:
    explicit GpuCLI();
    virtual string get_output();
    virtual vector<GPU> run() = 0;
};

class WindowsNVIDIAInfo: public GpuCLI
{
protected:
    GPU parse_gpu() override;
public:
    vector<GPU> run() override;
};

class DarwinInfo: public GpuCLI
{
protected:
    GPU parse_gpu() override;
public:
    vector<GPU> run() override;
};

#endif // GPU_CLI_H