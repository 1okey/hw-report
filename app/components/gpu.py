from PyQt5 import QtWidgets

from app.gui.components import gpu_layout

class GPUComponent(QtWidgets.QWidget, gpu_layout.Ui_Form):
    def __init__(self, parent, gpu_list):
        super().__init__(parent)
        self.setupUi(self)
        self.gpu = gpu_list

        self.set_values()
    
    def set_values(self):
        self.name_field.setText(self.gpu.name)
        self.driver_field.setText(self.gpu.driver_version)

        self.memory_util_field.setText(self.gpu.utilization['memory'])
        self.gpu_util_field.setText(self.gpu.utilization['gpu'])
        
        self.memory_total_field.setText(self.gpu.memory['total'])
        self.memory_used_field.setText(self.gpu.memory['used'])
        
        self.temp_gpu_field.setText(self.gpu.temperature['gpu'])
        self.temp_mem_field.setText(self.gpu.temperature['memory'])


