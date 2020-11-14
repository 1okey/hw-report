from PyQt5 import QtWidgets, uic
import sys
 
app = QtWidgets.QApplication([])
window = uic.loadUi("./gui/main_window.ui")
 
window.show()
sys.exit(app.exec())