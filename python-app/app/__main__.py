import sys
from PyQt5 import QtWidgets

from app.app import Application

qt_app = QtWidgets.QApplication(sys.argv)

app = Application()
app.show()

sys.exit(qt_app.exec())