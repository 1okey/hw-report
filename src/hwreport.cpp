#include "include/hwreport.h"

HWReport::HWReport(QWidget *parent):
    QMainWindow(parent),
    ui(new Ui::HWReport)
{
    ui->setupUi(this);

}

HWReport::~HWReport()
{
    delete ui;
}
