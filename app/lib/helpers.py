import os
import platform
from subprocess import Popen, PIPE
from distutils import spawn

from app.lib.gpu import GPU


def get_gpus_info():

    executable = get_executable()

    memory_flags = 'memory.total,memory.used,memory.free'
    util_flags = 'utilization.gpu,utilization.memory'
    temp_flags = 'temperature.memory,temperature.gpu'
    query_flags = f'--query-gpu=index,uuid,driver_version,name,gpu_serial,display_active,display_mode,{util_flags},{memory_flags},{temp_flags}'
    format_flags = "--format=csv" 

    try:
        proc = Popen([executable, f'{query_flags}', format_flags, '--nounits'], stdout=PIPE)
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


def parse_gpu(headers, props_list):
    data = dict()
    for field in GPU.REQUIRED_FIELDS:
        if field in headers:
            index = headers.index(field)
            data[field] = props_list[index]
        else:
            data[field] = 'N/A'

    return GPU(data)


def build_gpus():
    output = get_gpus_info()
    gpu_rows = output.split(os.linesep)
    headers = [header[0: header.find(' ')] if ' ' in header else header for header in gpu_rows[0].split(', ') ]
    gpus = list()

    for line in gpu_rows[1:]:
        if len(line) > 0:
            gpus.append(parse_gpu(headers, line.split(', ')))

    return gpus

    

