#!/bin/bash
set -eux

function not() {
    if ! $@; then
        return 0
    fi
    return 1
}


function retry() {
    local TIMEOUT=60
    local SLEEP=10
    while :
    do
        case "${1}" in
            TIMEOUT=*|SLEEP=*)
                local "${1}"
                shift
                ;;
            *)
                break
                ;;
        esac
    done

    local start_time="$(date +"%s")"
    local end_time="$(($start_time + ${TIMEOUT}))"
    while true; do
        if $@; then
            return 0
        fi
        local current_time="$(date +"%s")"
        local remaining_time="$((end_time - current_time))"
        if [[ "${remaining_time}" -lt 0 ]]; then
            return 1
        fi
        local sleep_time="${SLEEP}"
        if [[ "${remaining_time}" -lt "${SLEEP}" ]]; then
            sleep_time="${remaining_time}"
        fi
        sleep "${sleep_time}"
    done
    return 1
}

function kube_delete_namespace_and_wait() {
    local namespace=$1
    # Delete ESCs and C* clusters in the namespace
    if ! retry kubectl --namespace "${namespace}" \
         delete \
         elasticsearchclusters,cassandraclusters \
         --now \
         --all
    then
        # If multiple attempts to delete resources fails, display the remaining
        # resources.
        return 1
    fi
    # This is a work around for Kubernetes 1.7 which doesn't support garbage
    # collection of resources owned by third party resources.
    # See https://github.com/kubernetes/kubernetes/issues/44507
    if ! retry kubectl --namespace "${namespace}" \
         delete \
         deployments,replicasets,statefulsets,pods \
         --now \
         --all
    then
        # If multiple attempts to delete resources fails, display the remaining
        # resources.
        return 1
    fi
    if ! wait_for_namespace_empty "${namespace}"; \
    then
        return 1
    fi
    return 0
}

# waits for a namespace to contain 0 pods
function wait_for_namespace_empty() {
    local namespace=$1
    if retry TIMEOUT=300 namespace_empty "${namespace}"; then
        return 0
    fi
    return 1
}

function namespace_empty() {
    local namespace=$1
    if stdout_equals "0" kubectl \
        --namespace "${namespace}" \
        get pods \
        --output='go-template={{len .items}}'; then
        return 0
    fi
    return 1
}

function kube_event_exists() {
    local namespace="${1}"
    local event="${2}"
    local go_template='{{range .items}}{{.source.component}}:{{.involvedObject.kind}}:{{.type}}:{{.reason}}{{"\n"}}{{end}}'
    if kubectl get \
               --namespace "${namespace}" \
               events \
               --output=go-template="${go_template}" \
            | grep "^${event}$"; then
        return 0
    fi
    return 1
}

function simulate_unresponsive_cassandra_process() {
    local namespace=$1
    local pod=$2
    local container=$3
    # Decommission causes cassandra to stop accepting CQL connections.
    kubectl \
        --namespace="${namespace}" \
        exec "${pod}" --container="${container}" -- \
        nodetool decommission
}

function stdout_equals() {
    local expected="${1}"
    shift
    local actual=$("${@}")
    if [[ "${expected}" == "${actual}" ]]; then
        return 0
    fi
    return 1
}

function stdout_gt() {
    local expected="${1}"
    shift
    local actual=$("${@}")
    re='^[0-9]+$'
    if ! [[ "${actual}" =~ $re ]]; then
        echo "${actual} is not a number"
        return 1
    fi
    if [[ "${actual}" -gt "${expected}" ]]; then
        return 0
    fi
    return 1
}

function dump_debug_logs() {
    local namespace="${1}"
    local output_dir="$(pwd)/_artifacts/${namespace}"
    echo "Dumping cluster state to ${output_dir}"
    mkdir -p "${output_dir}"
    kubectl cluster-info dump --namespaces "${namespace}" > "${output_dir}/dump.txt" || true
}

function fail_and_exit() {
    local namespace="${1}"

    kubectl api-versions
    kubectl get apiservice -o yaml

    dump_debug_logs "${namespace}"

    exit 1
}
