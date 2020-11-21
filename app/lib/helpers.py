import os
import platform
from subprocess import Popen, PIPE
from distutils import spawn


def get_gpus_info():

    executable = get_executable()

    memory_flags = 'memory.total,memory.used,memory.free'
    util_flags = 'utilization.gpu,utilization.memory'
    temp_flags = 'temperature.memory,temperature.gpu'
    query_flags = f'--query-gpu=index,uuid,driver_version,name,gpu_serial,display_active,display_mode,{util_flags},{memory_flags},{temp_flags}'
    format_flags = "--format=csv" 

    try:
        proc = Popen([executable, f'{query_flags}', format_flags], stdout=PIPE)
        stdout, stderr = proc.communicate()
    except:
        print(stderr)

    return stdout.decode('UTF-8') if stdout != None else ''


def get_executable():
    executable = spawn.find_executable('nvidia-smi')
    if platform.system == 'Windows':
        pass
    elif platform.system == 'Darwin' or len(executable) == 0:
        # not supported or missing executable
        return []

    return executable


def parse_gpu(csv_line):
    print(csv_line)


def build_gpus():
    output = get_gpus_info()
    gpu_rows = output.split(os.linesep)
    headers = gpu_rows[0].split(', ')
    gpus = list()
    for line in gpu_rows[1:]:
        gpus.append(parse_gpu(line))

    

