#!/bin/bash

# set your own paths here
FB="/usr/local/fb46/fb"
FBPROC="/usr/local/fbproc/fbproc"

$FBPROC -f hello.bas -d defines.def

$FB hello.bas.fbp

rm -rf *asm *fbp *lbl *lst *o *xex
