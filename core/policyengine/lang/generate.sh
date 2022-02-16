#!/bin/bash

antlr4='java -Xmx500M -cp ".:/usr/local/lib/antlr-4.9.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
grun='java -Xmx500M -cp ".:/usr/local/lib/antlr-4.9.2-complete.jar:$CLASSPATH" org.antlr.v4.gui.TestRig'

$antlr4 -Dlanguage=Go -o parser -package parser -visitor Sfpl.g4
