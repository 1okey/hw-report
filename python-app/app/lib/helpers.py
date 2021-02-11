import platform
from distutils import spawn

from app.lib.gpus_cli import WindowsNVIDIAInfo, DarwinInfo

def get_executable():
    if platform.system() == 'Windows':
        # TODO consider using dxdiag to support
        executable = spawn.find_executable('nvidia-smi')
        return WindowsNVIDIAInfo(executable)
        
    elif platform.system() == 'Darwin':
        executable = spawn.find_executable('system_profiler')
        return DarwinInfo(executable)

    return None

def get_main_gpu():
    gpu_cli = get_executable()
    gpus = gpu_cli.run()
    
    if len(gpus) > 1:
        return gpus[-1]
    
    return gpus[0]

