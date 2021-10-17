#!/usr/bin/python
# -*- coding: utf-8 -*-

# モジュールのimport
import sys
import os

# モジュール os.pathのクラスsplitextを初期化
from os.path import splitext
from objc import YES, NO
from Foundation import NSData
from AppKit import *

NSApp = NSApplication.sharedApplication()

pdf_file = sys.argv[1].decode('utf-8')
fileData = NSData.dataWithContentsOfFile_(pdf_file)
refer = NSPDFImageRep.imageRepWithData_(fileData)
numOfPages = refer.pageCount()


print pdf_file
print numOfPages


#if __name__ == '__main__':
#       main()
