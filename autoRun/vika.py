#!/usr/bin/python3.7
# -*- coding: utf-8 -*-
#
# Copyright (C) 2021 桜火, Inc. All Rights Reserved 
#
# @Time    : 2021/4/26 21:35
# @Author  : 桜火
# @Email   : xie@loli.fit
# @File    : aba.py
# @Software: PyCharm
from vika import Vika
vika = Vika("uskK7MnF6tOf89ammoV851g")
dst = vika.datasheet("https://vika.cn/space/spcNuHrHJgQVZ/workbench/dstXHYNcifpEsJuMjZ/viwvYeltzpZ5Z")
record = dst.records.create({"提交ID":"new record from Python SDK"})
print(GITHUB_EVENT_PATH)