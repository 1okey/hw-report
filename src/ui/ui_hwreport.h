/********************************************************************************
** Form generated from reading UI file 'window.ui'
**
** Created by: Qt User Interface Compiler version 5.11.1
**
** WARNING! All changes made in this file will be lost when recompiling UI file!
********************************************************************************/

#ifndef UI_HWREPORT_H
#define UI_HWREPORT_H

#include <QtCore/QVariant>
#include <QtWidgets/QApplication>
#include <QtWidgets/QMainWindow>
#include <QtWidgets/QTabWidget>
#include <QtWidgets/QWidget>
#include <QMainWindow>

QT_BEGIN_NAMESPACE

class Ui_HWReport
{
public:
    QWidget *centralwidget;
    QTabWidget *tabWidget;
    QWidget *cpu;
    QWidget *gpu;
    QWidget *storage;

    void setupUi(QMainWindow *HWReport)
    {
        if (HWReport->objectName().isEmpty())
            HWReport->setObjectName(QStringLiteral("HWReport"));
        HWReport->resize(505, 463);
        HWReport->setToolButtonStyle(Qt::ToolButtonIconOnly);
        centralwidget = new QWidget(HWReport);
        centralwidget->setObjectName(QStringLiteral("centralwidget"));
        tabWidget = new QTabWidget(centralwidget);
        tabWidget->setObjectName(QStringLiteral("tabWidget"));
        tabWidget->setGeometry(QRect(10, 10, 481, 441));
        tabWidget->setIconSize(QSize(22, 22));
        tabWidget->setMovable(false);
        cpu = new QWidget();
        cpu->setObjectName(QStringLiteral("cpu"));
        tabWidget->addTab(cpu, QString());
        gpu = new QWidget();
        gpu->setObjectName(QStringLiteral("gpu"));
        tabWidget->addTab(gpu, QString());
        storage = new QWidget();
        storage->setObjectName(QStringLiteral("storage"));
        tabWidget->addTab(storage, QString());
        HWReport->setCentralWidget(centralwidget);

        retranslateUi(HWReport);

        tabWidget->setCurrentIndex(0);


        QMetaObject::connectSlotsByName(HWReport);
    } // setupUi

    void retranslateUi(QMainWindow *HWReport)
    {
        HWReport->setWindowTitle(QApplication::translate("HWReport", "MainWindow", nullptr));
        tabWidget->setTabText(tabWidget->indexOf(cpu), QApplication::translate("HWReport", "CPU", nullptr));
        tabWidget->setTabText(tabWidget->indexOf(gpu), QApplication::translate("HWReport", "GPU", nullptr));
        tabWidget->setTabText(tabWidget->indexOf(storage), QApplication::translate("HWReport", "Storage", nullptr));
    } // retranslateUi

};

namespace Ui {
    class HWReport: public Ui_HWReport {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_HWREPORT_H
