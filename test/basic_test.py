from app.app import Application
import pytest

@pytest.fixture
def application():
    return Application()

def test_app(application):
    assert application.main_window != None