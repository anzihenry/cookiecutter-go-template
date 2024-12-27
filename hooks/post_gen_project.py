import os
import shutil

def remove_unnecessary_dirs():
    project_type = "{{ cookiecutter.project_type }}"
    if project_type == "library":
        shutil.rmtree("cmd")
        shutil.rmtree("proto")
    elif project_type == "console":
        shutil.rmtree("proto")
        shutil.rmtree("examples")
    elif project_type == "web":
        shutil.rmtree("proto")
        shutil.rmtree("examples")
    elif project_type == "grpc":
        shutil.rmtree("examples")
    else:
        print(f"Unknown project type: {project_type}")

def rename_main_file():
    project_type = "{{ cookiecutter.project_type }}"
    main_file_map = {
        "console": "main_console.go",
        "web": "main_web.go",
        "grpc": "main_grpc.go",
        "library": "main_library.go",
    }

    selected_main_file = main_file_map.get(project_type)
    if selected_main_file:
        src = f"cmd/{{cookiecutter.main_package}}/{selected_main_file}"
        dest = f"cmd/{{cookiecutter.main_package}}/main.go"
        os.rename(src, dest)

    # 删除未选中的主文件模板
    for file in main_file_map.values():
        if file != selected_main_file:
            file_path = f"cmd/{{cookiecutter.main_package}}/{file}"
            if os.path.exists(file_path):
                os.remove(file_path)def rename_main_file():
    project_type = "{{ cookiecutter.project_type }}"
    main_file_map = {
        "console": "main_console.go",
        "web": "main_web.go",
        "grpc": "main_grpc.go",
        "library": "main_library.go",
    }

    selected_main_file = main_file_map.get(project_type)
    if selected_main_file:
        src = f"cmd/{{cookiecutter.main_package}}/{selected_main_file}"
        dest = f"cmd/{{cookiecutter.main_package}}/main.go"
        os.rename(src, dest)

    # 删除未选中的主文件模板
    for file in main_file_map.values():
        if file != selected_main_file:
            file_path = f"cmd/{{cookiecutter.main_package}}/{file}"
            if os.path.exists(file_path):
                os.remove(file_path)

def set_license():
    license_map = {
        "MIT": "LICENSE_MIT",
        "Apache-2.0": "LICENSE_Apache-2.0",
        "GPL-3.0": "LICENSE_GPL-3.0",
        "None": "LICENSE_None",
    }
    selected_license = "{{ cookiecutter.license }}"
    license_file = license_map.get(selected_license)

    if license_file != "LICENSE_None" and license_file:
        os.rename(license_file, "LICENSE")

    # 删除未使用的许可证模板
    for file in license_map.values():
        if file != license_file and os.path.exists(file):
            os.remove(file)

def set_go_mod():
    project_type = "{{ cookiecutter.project_type }}"
    go_mod_map = {
        "library": "go.mod_library",
        "console": "go.mod_console",
        "web": "go.mod_web",
        "grpc": "go.mod_grpc",
    }

    selected_go_mod = go_mod_map.get(project_type)
    if selected_go_mod:
        os.rename(selected_go_mod, "go.mod")

    # 删除未使用的 go.mod 模板
    for file in go_mod_map.values():
        if file != selected_go_mod and os.path.exists(file):
            os.remove(file)

def set_makefile():
    project_type = "{{ cookiecutter.project_type }}"
    makefile_map = {
        "library": "Makefile_library",
        "console": "Makefile_console",
        "web": "Makefile_web",
        "grpc": "Makefile_grpc",
    }

    selected_makefile = makefile_map.get(project_type)
    if selected_makefile:
        os.rename(selected_makefile, "Makefile")

    # 删除未使用的 Makefile 模板
    for file in makefile_map.values():
        if file != selected_makefile and os.path.exists(file):
            os.remove(file)

set_makefile()


# 执行钩子
rename_main_file()
set_license()
set_go_mod()
set_makefile()
remove_unnecessary_dirs()
