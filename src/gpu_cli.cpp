#include "include/gpu_cli.h"
#include "include/utils.h"

using std::string;
using std::vector;

explicit GpuCLI::GpuCLI(string binary_name, vector<string> cli_params)
: bin_name(binary_name), cli_params(cli_params)
{}

GpuCLI::~GpuCLI(){}

string GpuCLI::get_output()
{
    char buffer[128];
    string result = "";

    string command = command;
    if (cli_params.size() > 0) {
        for (string& param : cli_params) {
            command += " " + param;
        }
    }

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

bool GpuCLI::can_be_run()
{
    FILE* pipe = run(bin_name, "r");
    if (!pipe) {
        return false;
    }
    
    close(pipe);
    return true;
}

// ================ WindowsInfo ================
GPU WindowsInfo::parse_gpu() 
{

}

vector<GPU> WindowsInfo::get_gpus()
{
    QString cli_output = QString::fromStdString(this->get_output());
}
// ============== WindowsInfo END ==============

//  =============== DarwinInfo =================
GPU DarwinInfo::parse_gpu() 
{

}

vector<GPU> DarwinInfo::get_gpus()
{
    QString cli_output = QString::fromStdString(this->get_output());
}
// ============== DarwinInfo END ==============

//  ================ DummyInfo ================
GPU DummyInfo::parse_gpu()
{
    return GPU();
}

vector<GPU> DummyInfo::get_gpus() 
{
    return { parse_gpu() };
}
//  ============== DummyInfo END ==============
