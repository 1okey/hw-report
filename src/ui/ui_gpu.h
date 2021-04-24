/********************************************************************************
** Form generated from reading UI file 'gpu_layout.ui'
**
** Created by: Qt User Interface Compiler version 5.11.1
**
** WARNING! All changes made in this file will be lost when recompiling UI file!
********************************************************************************/

#ifndef UI_GPU_H
#define UI_GPU_H

#include <QtCore/QVariant>
#include <QtWidgets/QApplication>
#include <QtWidgets/QHBoxLayout>
#include <QtWidgets/QLabel>
#include <QtWidgets/QLineEdit>
#include <QtWidgets/QVBoxLayout>
#include <QtWidgets/QWidget>

QT_BEGIN_NAMESPACE

class Ui_Form
{
public:
    QWidget *verticalWidget_2;
    QVBoxLayout *gpu_layout;
    QLabel *name_label;
    QLineEdit *name_field;
    QLabel *driver_label;
    QLineEdit *driver_field;
    QHBoxLayout *horizontalLayout;
    QVBoxLayout *memory_util_box;
    QLabel *memory_util_label;
    QLineEdit *memory_util_field;
    QVBoxLayout *gpu_util_box;
    QLabel *gpu_util_label;
    QLineEdit *gpu_util_field;
    QHBoxLayout *horizontalLayout_2;
    QVBoxLayout *memory_total_box;
    QLabel *memory_total_label;
    QLineEdit *memory_total_field;
    QVBoxLayout *memory_used_box;
    QLabel *memory_used_label;
    QLineEdit *memory_used_field;
    QHBoxLayout *horizontalLayout_3;
    QVBoxLayout *temp_gpu_box;
    QLabel *temp_gpu_label;
    QLineEdit *temp_gpu_field;
    QVBoxLayout *temp_mem_box;
    QLabel *temp_mem_label;
    QLineEdit *temp_mem_field;

    void setupUi(QWidget *Form)
    {
        if (Form->objectName().isEmpty())
            Form->setObjectName(QStringLiteral("Form"));
        Form->resize(475, 417);
        verticalWidget_2 = new QWidget(Form);
        verticalWidget_2->setObjectName(QStringLiteral("verticalWidget_2"));
        verticalWidget_2->setGeometry(QRect(10, 10, 451, 391));
        gpu_layout = new QVBoxLayout(verticalWidget_2);
        gpu_layout->setObjectName(QStringLiteral("gpu_layout"));
        name_label = new QLabel(verticalWidget_2);
        name_label->setObjectName(QStringLiteral("name_label"));
        name_label->setMaximumSize(QSize(16777215, 20));

        gpu_layout->addWidget(name_label);

        name_field = new QLineEdit(verticalWidget_2);
        name_field->setObjectName(QStringLiteral("name_field"));
        name_field->setFrame(true);
        name_field->setReadOnly(true);

        gpu_layout->addWidget(name_field);

        driver_label = new QLabel(verticalWidget_2);
        driver_label->setObjectName(QStringLiteral("driver_label"));
        driver_label->setMaximumSize(QSize(16777215, 20));

        gpu_layout->addWidget(driver_label);

        driver_field = new QLineEdit(verticalWidget_2);
        driver_field->setObjectName(QStringLiteral("driver_field"));
        driver_field->setFrame(true);
        driver_field->setReadOnly(true);

        gpu_layout->addWidget(driver_field);

        horizontalLayout = new QHBoxLayout();
        horizontalLayout->setObjectName(QStringLiteral("horizontalLayout"));
        memory_util_box = new QVBoxLayout();
        memory_util_box->setObjectName(QStringLiteral("memory_util_box"));
        memory_util_label = new QLabel(verticalWidget_2);
        memory_util_label->setObjectName(QStringLiteral("memory_util_label"));
        memory_util_label->setMaximumSize(QSize(16777215, 20));

        memory_util_box->addWidget(memory_util_label);

        memory_util_field = new QLineEdit(verticalWidget_2);
        memory_util_field->setObjectName(QStringLiteral("memory_util_field"));
        memory_util_field->setReadOnly(true);

        memory_util_box->addWidget(memory_util_field);


        horizontalLayout->addLayout(memory_util_box);

        gpu_util_box = new QVBoxLayout();
        gpu_util_box->setObjectName(QStringLiteral("gpu_util_box"));
        gpu_util_label = new QLabel(verticalWidget_2);
        gpu_util_label->setObjectName(QStringLiteral("gpu_util_label"));
        gpu_util_label->setMaximumSize(QSize(16777215, 20));

        gpu_util_box->addWidget(gpu_util_label);

        gpu_util_field = new QLineEdit(verticalWidget_2);
        gpu_util_field->setObjectName(QStringLiteral("gpu_util_field"));
        gpu_util_field->setReadOnly(true);

        gpu_util_box->addWidget(gpu_util_field);


        horizontalLayout->addLayout(gpu_util_box);


        gpu_layout->addLayout(horizontalLayout);

        horizontalLayout_2 = new QHBoxLayout();
        horizontalLayout_2->setObjectName(QStringLiteral("horizontalLayout_2"));
        memory_total_box = new QVBoxLayout();
        memory_total_box->setObjectName(QStringLiteral("memory_total_box"));
        memory_total_label = new QLabel(verticalWidget_2);
        memory_total_label->setObjectName(QStringLiteral("memory_total_label"));
        memory_total_label->setMaximumSize(QSize(16777215, 20));

        memory_total_box->addWidget(memory_total_label);

        memory_total_field = new QLineEdit(verticalWidget_2);
        memory_total_field->setObjectName(QStringLiteral("memory_total_field"));
        memory_total_field->setReadOnly(true);

        memory_total_box->addWidget(memory_total_field);


        horizontalLayout_2->addLayout(memory_total_box);

        memory_used_box = new QVBoxLayout();
        memory_used_box->setObjectName(QStringLiteral("memory_used_box"));
        memory_used_label = new QLabel(verticalWidget_2);
        memory_used_label->setObjectName(QStringLiteral("memory_used_label"));
        memory_used_label->setMaximumSize(QSize(16777215, 20));

        memory_used_box->addWidget(memory_used_label);

        memory_used_field = new QLineEdit(verticalWidget_2);
        memory_used_field->setObjectName(QStringLiteral("memory_used_field"));
        memory_used_field->setReadOnly(true);

        memory_used_box->addWidget(memory_used_field);


        horizontalLayout_2->addLayout(memory_used_box);


        gpu_layout->addLayout(horizontalLayout_2);

        horizontalLayout_3 = new QHBoxLayout();
        horizontalLayout_3->setObjectName(QStringLiteral("horizontalLayout_3"));
        temp_gpu_box = new QVBoxLayout();
        temp_gpu_box->setObjectName(QStringLiteral("temp_gpu_box"));
        temp_gpu_label = new QLabel(verticalWidget_2);
        temp_gpu_label->setObjectName(QStringLiteral("temp_gpu_label"));
        temp_gpu_label->setMaximumSize(QSize(16777215, 20));

        temp_gpu_box->addWidget(temp_gpu_label);

        temp_gpu_field = new QLineEdit(verticalWidget_2);
        temp_gpu_field->setObjectName(QStringLiteral("temp_gpu_field"));
        temp_gpu_field->setReadOnly(true);

        temp_gpu_box->addWidget(temp_gpu_field);


        horizontalLayout_3->addLayout(temp_gpu_box);

        temp_mem_box = new QVBoxLayout();
        temp_mem_box->setObjectName(QStringLiteral("temp_mem_box"));
        temp_mem_label = new QLabel(verticalWidget_2);
        temp_mem_label->setObjectName(QStringLiteral("temp_mem_label"));
        temp_mem_label->setMaximumSize(QSize(16777215, 20));

        temp_mem_box->addWidget(temp_mem_label);

        temp_mem_field = new QLineEdit(verticalWidget_2);
        temp_mem_field->setObjectName(QStringLiteral("temp_mem_field"));
        temp_mem_field->setReadOnly(true);

        temp_mem_box->addWidget(temp_mem_field);


        horizontalLayout_3->addLayout(temp_mem_box);


        gpu_layout->addLayout(horizontalLayout_3);


        retranslateUi(Form);

        QMetaObject::connectSlotsByName(Form);
    } // setupUi

    void retranslateUi(QWidget *Form)
    {
        Form->setWindowTitle(QApplication::translate("Form", "Form", nullptr));
        name_label->setText(QApplication::translate("Form", "GPU Name", nullptr));
        driver_label->setText(QApplication::translate("Form", "Driver Version", nullptr));
        memory_util_label->setText(QApplication::translate("Form", "Memory Utilization, %", nullptr));
        gpu_util_label->setText(QApplication::translate("Form", "GPU Utilization, %", nullptr));
        memory_total_label->setText(QApplication::translate("Form", "Total Memory, MiB", nullptr));
        memory_used_label->setText(QApplication::translate("Form", "Used Memory, MiB", nullptr));
        temp_gpu_label->setText(QApplication::translate("Form", "GPU Temperature, C", nullptr));
        temp_mem_label->setText(QApplication::translate("Form", "Memory Temperature, C", nullptr));
    } // retranslateUi

};

namespace Ui {
    class Form: public Ui_Form {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_GPU_H
