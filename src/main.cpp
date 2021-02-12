#include "include/hwreport.h"
#include <QApplication>

int main(int argc, char *argv[])
{
    QApplication hwreport(argc, argv);

    HWReport app;
    app.show();

    return hwreport.exec();
}