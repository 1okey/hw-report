from PyQt5 import QtWidgets

from app.gui.components import main

class Application(QtWidgets.QMainWindow, main.Ui_MainWindow):
    def __init__(self):
        super().__init__()
        self.setupUi(self) 