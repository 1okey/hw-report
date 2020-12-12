import os
import json

from subprocess import Popen, PIPE

from app.lib.gpu import GPU

class gpu_cli:
    def __init__(self, binary_path, params):
        self.binary_path = binary_path
        self.cli_params = params

    def get_output(self):
        try:
            proc = Popen([self.binary_path, *self.cli_params], stdout=PIPE)
            stdout, stderr = proc.communicate()
        except:
            print(stderr)

        return stdout.decode('UTF-8') if stdout != None else ''

class WindowsNVIDIAInfo(gpu_cli):
    REQUIRED_FIELDS = [
        'name',
        'driver_version', 
        'utilization.gpu', 
        'utilization.memory',
        'memory.total',
        'memory.used',
        'temperature.gpu', 
        'temperature.memory',
    ]

    def __init__(self, bin_path):
        self.binary_path = bin_path

        self.memory_flags = 'memory.total,memory.used,memory.free'
        self.util_flags = 'utilization.gpu,utilization.memory'
        self.temp_flags = 'temperature.memory,temperature.gpu'

        self.cli_params = [
            f'--query-gpu=index,uuid,driver_version,name,gpu_serial,display_active,display_mode,{self.util_flags},{self.memory_flags},{self.temp_flags}'
            '--format=csv',
            '--nounits',
        ]

        super().__init__(self.binary_path, self.cli_params)
    
    def run(self):
        output = self.get_output()
        gpu_rows = output.split(os.linesep)
        headers = [header[0: header.find(' ')] if ' ' in header else header for header in gpu_rows[0].split(', ') ]
        gpus = list()

        for line in gpu_rows[1:]:
            if len(line) > 0:
                gpus.append(self.__parse_gpu(headers, line.split(', ')))

        return gpus

    def __parse_gpu(self, headers, props_list):
        data = dict()
        for field in WindowsNVIDIAInfo.REQUIRED_FIELDS:
            if field in headers:
                index = headers.index(field)
                data[field] = props_list[index]
            else:
                data[field] = 'N/A'

        return GPU(
            name=data['name'],
            driver_version=data['driver_version'],
            total_memory=data['memory.total'],
        )


class DarwinInfo(gpu_cli):
    def __init__(self, binary_path):
        self.binary_path = binary_path
        self.cli_params = [
            'SPDisplaysDataType',
            '-json'
        ]

        super().__init__(self.binary_path, self.cli_params)

    def run(self):
        output_json = self.get_output()
        gpu_data_list = json.loads(output_json)["SPDisplaysDataType"]
        gpus = list()
        for gpu_data in gpu_data_list:
            gpus.append(self.__parse_gpu(gpu_data))
            
        return gpus

    def __parse_gpu(self, gpu_data: dict):
        print(gpu_data)

        total_memory = gpu_data['spdisplays_vram_shared'] if 'spdisplays_vram_shared' in gpu_data else gpu_data['spdisplays_vram']

        return GPU(
            name=gpu_data['sppci_model'],
            driver_version='N/A',
            total_memory=total_memory
        )