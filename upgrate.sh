#!/bin/bash

basepath=$(cd `dirname $0`; pwd)

SHELL_DIR="${basepath}/shell"
SCRIPTS_DIR="${basepath}/scripts"
PACKAGE_DIR="${basepath}/package"
INIT_DIR="${basepath}/init.d"


source /etc/svxnetworks/conf.d/base.conf
source $versionconf

[[ "$softversion" =~ ^6.2.[01] ]] || exit 1

touch /etc/forge/activation.conf
[ -e /usr/lib/python2.7/site-packages/kafka ] || tar -zxf $basepath/kafka-python-2.0.2.tar.gz -C /usr/lib/python2.7/site-packages/

#7xcli shell目录
CLI_ROOT='/usr/local/sv'


