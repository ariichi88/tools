#! /usr/bin/env python3
# coding: UTF-8

import os
import sys
import datetime
import shutil

FromDir = '/home/yuuichi/Dropbox/カメラアップロード/'
ToDir = '/home/yuuichi/Dropbox/Photo/'


def ChackAndMakeDir(date):
    dt = datetime.datetime.strptime(date, '%Y/%m/%d')
    datedir = dt.strftime('%Y/%m/%d')
    if os.path.exists(ToDir + datedir):
        print('すでにあるディレクトリは選択できません')
        sys.exit(1)
    else:
        os.makedirs(ToDir + datedir)
        return datedir


def CopyFiles(datedir, newname):
    datestr = datedir.replace('/', '-')
    fromfiles = [f for f in os.listdir(FromDir) if datestr in f]
    fromfiles.sort()
    for count, fromfile in enumerate(fromfiles):
        _, ext = os.path.splitext(fromfile)
        tofile = newname + '-' + format(count + 1, '02') + ext
        shutil.copy2(FromDir + fromfile, ToDir + datedir + '/' + tofile)


if __name__ == '__main__':
    arg = sys.argv
    if len(arg) == 1:
        print('引数がありません')
        sys.exit(1)
    elif len(arg) == 2:
        dt = datetime.datetime.now()
        date = dt.strftime('%Y/%m/%d')
        name = arg[1]
    elif len(arg) == 3:
        date = arg[1]
        name = arg[2]
    else:
        print('引数が多すぎます')
        sys.exit(1)
    CopyFiles(ChackAndMakeDir(date), name)
