class GPU:

    REQUIRED_FIELDS = (
        'name',
        'driver_version', 
        'utilization.gpu', 
        'utilization.memory',
        'memory.total',
        'memory.used',
        'temperature.gpu', 
        'temperature.memory',
    )

    def __init__(self, data):
        self.name = data['name'] 

        self.driver_version = data['driver_version'] 

        self.utilization = {
            'gpu': data['utilization.gpu'],
            'memory': data['utilization.memory']
        } 

        self.memory = {
            'total': data['memory.total'],
            'used': data['memory.used']
        }

        self.temperature = {
            'gpu': data['temperature.gpu'],
            'memory': data['temperature.memory']
        }