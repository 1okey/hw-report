import sys
from PyQt5 import QtWidgets

from app.app import Application

qt_app = QtWidgets.QApplication(sys.argv)
app = Application()
app.show()
qt_app.exec_()