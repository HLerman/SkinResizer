#!/usr/bin/env python3
import subprocess
import os

def build(goos, goarch, filename):
    env = os.environ.copy()
    env["GOOS"] = goos
    env["GOARCH"] = goarch

    subprocess.check_call(["go", "build", "-o", filename], env = env)

build("windows","amd64","SkinResizer.exe")
build("linux","amd64","SkinResizer")