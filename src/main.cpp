#include <QApplication>

#include "include/hwreport.h"
#include "include/utils.h"
#include "include/gpu_cli.h"

int main(int argc, char *argv[])
{
    QApplication hwreport(argc, argv);

    HWReport app;
    app.show();

    GpuCLI* cli = get_cli_binary();
    vector<GPU> gpus = cli->get_gpus();

    return hwreport.exec();
}