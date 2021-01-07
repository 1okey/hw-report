class GPU:

    NAME = 'name'
    VENDOR = 'vendor'
    DRIVER_VERSION = 'driver_version'
    DRIVER_TOTAL_MEMORY = 'total_memory'

    def __init__(self, name, driver_version, total_memory):
        self.name = name
        self.driver_version = driver_version

        self.utilization = {
            'gpu': 'N/A', # data['utilization.gpu'],
            'memory': 'N/A' # data['utilization.memory']
        } 

        self.memory = {
            'total': total_memory, # data['memory.total'],
            'used': 'N/A' # data['memory.used']
        }

        self.temperature = {
            'gpu': 'N/A', # data['temperature.gpu'],
            'memory': 'N/A' # data['temperature.memory']
        }