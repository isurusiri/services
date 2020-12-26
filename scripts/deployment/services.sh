#!/bin/bash

# services <ACTION> <COMPONENT> <NAME>
# services install namespace names
ACTION=$1
COMPONENT=$2
COMPONENT_NAME=$3 

# ---------------------------------------------------------
# Section: Namespaces
# ---------------------------------------------------------

function handle_namespaces() {
    kubectl $ACTION -f workloads/namespaces.yaml
}


# ---------------------------------------------------------
# Section: Deployments
# ---------------------------------------------------------

function handle_bar_service() {
    kubectl $ACTION -f workloads/bar/service-account.yaml
    kubectl $ACTION -f workloads/bar/deployment.yaml
}

function handle_foo_service() {
    kubectl $ACTION -f workloads/foo/service-account.yaml
    kubectl $ACTION -f workloads/foo/deployment.yaml
}

# ---------------------------------------------------------
# Section: Services
# ---------------------------------------------------------

function handle_bar_service_svc() {
    kubectl $ACTION -f workloads/bar/service.yaml
}

function handle_foo_service_svc() {
    kubectl $ACTION -f workloads/foo/service.yaml
}

# ---------------------------------------------------------
# Section: Main
# ---------------------------------------------------------

echo "Executing commands: for ${ACTION} ${COMPONENT} ${COMPONENT_NAME}"

function help() {
    # Display Help
    echo "Functionalities made available through this script."
    echo
    echo "Syntax: services <ACTION> <COMPONENT> <NAME>"
    echo "parameters:"
    echo "ACTION              Kubernetes action to be performed."
    echo "                      Applicable values:"
    echo "                      -  create: Creates the compoent specifed by COMPONENT parameter."
    echo "                      -  apply:  Create and override if already exist the compoent specifed by COMPONENT parameter."
    echo "                      -  delete: Deletes the compoent specifed by COMPONENT parameter."
    echo "COMPONENT           Kubernetes resource type."
    echo "                      Applicable values:"
    echo "                      -  namespace:  Specifies the Kubernetes namespaces."
    echo "                      -  deployment: Specifies the Kubernetes deployments."
    echo "                      -  service:    Specifies the Kubernetes services."
    echo "COMPONENT_NAME      Name of the component to be handled."
    echo "                      Applicable values:"
    echo "                      -  bar:            Specifies the bar service (demo microservice)."
    echo "                      -  foo:            Specifies the foo service (demo microservice)."
    echo "                      -  demo-services:  Specifies the demo-services namespace."
    echo
}

case $COMPONENT_NAME in
    bar)
        handle_bar_service
        handle_bar_service_svc
        break
        ;;
    foo)
        handle_foo_service
        handle_foo_service_svc
        break
        ;;
    demo-services)
        handle_namespaces
        break
        ;;
    *)
        help
        break
        ;;
esac