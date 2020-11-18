from PyQt5 import QtWidgets, uic
import sys
 
class Application:
    def __init__(self):
        self.app = QtWidgets.QApplication([])
        self.main_window = uic.loadUi("./app/gui/main_window.ui")
    
    def run(self):
        self.main_window.show()
        sys.exit(self.app.exec())
 
