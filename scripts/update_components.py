from os import listdir, system, path

def update_components():
    dest_path = 'app/gui/components'
    src_path ='app/gui/layouts'
    for layout in listdir(src_path):
        component_path = f'{dest_path}/{layout[:-3]}.py'
        if path.exists(component_path):
            print(f'Updating {component_path}')
        else:
            print(f'Creating {component_path}')
        
        system(f'pyuic5 {src_path}/{layout} -o {component_path}')
        
update_components()