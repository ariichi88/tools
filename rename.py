#! /usr/bin/env python3
# coding: UTF-8

import os
import sys


def rename(newName):
    files = [f for f in os.listdir(os.getcwd()) if os.path.isfile(f)]
    files.sort()
    cnt = 1
    for file in files:
        _, ext = os.path.splitext(file)
        os.rename(file, newName + '-' + format(cnt, '02') + ext)
        cnt = cnt + 1


if __name__ == '__main__':
    args = sys.argv
    if len(args) < 2:
        print('引数が足りません')
        sys.exit(1)
    rename(args[1])
