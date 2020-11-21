from PyQt5 import QtWidgets, uic
import sys
from app.lib.helpers import build_gpus
 

class Application:
    def __init__(self):
        self.app = QtWidgets.QApplication([])
        self.main_window = uic.loadUi("./app/gui/main_window.ui")
        self.gpus = list()
    
    def run(self):
        self.main_window.show()
        self.gpus = build_gpus()
        sys.exit(self.app.exec())
 
