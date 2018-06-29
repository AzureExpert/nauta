#!/usr/bin/env sh

set -e

CURDIR=$(dirname $(realpath "$0"))
LIBDIR=${CURDIR}/libs
BINDIR=${CURDIR}/bin

. ${LIBDIR}/os-detect.sh
. ${LIBDIR}/bins-detect.sh
. ${LIBDIR}/ansible.sh

inventory() {
    if [ X"${ENV_INVENTORY}" = X"" ]; then
        echo "-c local -i localhost"
    else
        echo "--inventory-file=$(realpath ${ENV_INVENTORY})"
    fi
}

config() {
    if [ X"${ENV_CONFIG}" = X"" ]; then
        echo ""
    else
        echo "-e @$(realpath ${ENV_CONFIG})"
    fi
}

ansible_run() {
    export ANSIBLE_CONFIG=${CURDIR}/ansible.cfg
    ansible $(inventory) $(config) \
    -e "upgrade=${UPGRADE:-False}" -e "kubectl=${KUBECTL}" -e "helm=${HELM}" \
    -e "kubeconfig=$(realpath ${KUBECONFIG:-${HOME}/.kube/config})" -e "loader=${LOADER}" $@
    return $?
}

platform_ansible_run() {
    ansible_run ${CURDIR}/platform/dls.yml
    return $?
}

dls4e_ansible_run() {
    ansible_run ${CURDIR}/dls4e/install.yml
    return $?
}

COMMAND=$1

if [ X"${COMMAND}" = X"" ]; then
    >&2 echo "Command is not provided"
    exit 1
fi

case "${COMMAND}" in
  install) platform_ansible_run && dls4e_ansible_run
     ;;
  platform-install) platform_ansible_run
     ;;
  dls4e-install) dls4e_ansible_run
     ;;
  dls4e-upgrade) UPGRADE=True dls4e_ansible_run
     ;;
  *) echo "Unknown command"
     ;;
esac