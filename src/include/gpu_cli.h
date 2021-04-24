#ifndef GPU_CLI_H
#define GPU_CLI_H

#include <QString>

#include <string>
#include <vector>
#include <sstream>

#include "gpu.h"

using std::string;
using std::vector;

class GpuCLI {
protected:
    string bin_name;
    vector<string> cli_params;
protected:
    virtual GPU parse_gpu() = 0;
    string get_output();
public:
    GpuCLI();
    explicit GpuCLI(string binary_name, vector<string> cli_params);
    virtual ~GpuCLI();
    virtual vector<GPU> get_gpus() = 0;
    bool can_be_run();
};


class WindowsInfo: public GpuCLI
{
protected:
    GPU parse_gpu() override;
public:
    WindowsInfo(const string& bin_name, vector<string> cli_params): GpuCLI(bin_name, cli_params){}
    vector<GPU> get_gpus() override;
};


class DarwinInfo: public GpuCLI
{
protected:
    GPU parse_gpu() override;
public:
    DarwinInfo(const string& bin_name, vector<string> cli_params): GpuCLI(bin_name, cli_params) {}
    vector<GPU> get_gpus() override;
};


class DummyInfo: public GpuCLI
{
protected:
    GPU parse_gpu() override;
public:
    DummyInfo(): GpuCLI() {}
    vector<GPU> get_gpus() override;
};

#endif // GPU_CLI_H