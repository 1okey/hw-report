from PyQt5 import QtWidgets

from app.gui.components import main
from app.lib.helpers import get_main_gpu
from app.components.gpu import GPUComponent

class Application(QtWidgets.QMainWindow, main.Ui_MainWindow):
    def __init__(self):
        super().__init__()
        self.setupUi(self)
        self.setWindowTitle('hw-report')
        # self.tabWidget.setTabVisible(0, True)

        self.create_tabs()

    def create_tabs(self):
        self.tabs = {
            # 'cpu': gui_layout(self.gpu),
            'gpu': GPUComponent(self.gpu, get_main_gpu()),
            # 'storage': gui_layout(self.gpu),
        }

