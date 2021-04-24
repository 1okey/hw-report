#ifndef HWREPORT_H
#define HWREPORT_H

#include <QMainWindow>
#include "../ui/ui_hwreport.h"

class HWReport : public QMainWindow
{
    Q_OBJECT

public:
    explicit HWReport(QWidget *parent = nullptr);
    ~HWReport();
private:
    Ui::HWReport *ui;
    QString currentFile;
};

#endif