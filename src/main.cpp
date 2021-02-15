#include <QApplication>

#include "include/hwreport.h"
#include "include/utils.h"

int main(int argc, char *argv[])
{
    QApplication hwreport(argc, argv);

    HWReport app;
    app.show();

    return hwreport.exec();
}