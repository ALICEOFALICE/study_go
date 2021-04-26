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
import os
PR_TITLE=os.environ.get('PR_TITLE','')

vika = Vika("uskK7MnF6tOf89ammoV851g")
dst = vika.datasheet("https://vika.cn/space/spcNuHrHJgQVZ/workbench/dstXHYNcifpEsJuMjZ/viwvYeltzpZ5Z")
print(PR_TITLE)
record = dst.records.create({"提交ID":PR_TITLE})
