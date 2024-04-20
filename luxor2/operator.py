#!/usr/bin/env python3

import time
from kubernetes import client, config, watch

# Load Kubernetes configuration
config.load_kube_config()

# Custom resource definition
CRD_GROUP = 'example.com'
CRD_VERSION = 'v1'
CRD_PLURAL = 'httpserverconfigs'

# Kubernetes API client
custom_api = client.CustomObjectsApi()

def handle_config_event(event):
    config = event['object']
    if event['type'] == 'ADDED' or event['type'] == 'MODIFIED':
        # Handle configuration change
        print("Configuration updated:", config)
    elif event['type'] == 'DELETED':
        # Handle configuration deletion
        print("Configuration deleted:", config)

def watch_http_server_configs():
    resource_version = ''
    while True:
        stream = watch.Watch().stream(
            custom_api.list_namespaced_custom_object,
            CRD_GROUP,
            CRD_VERSION,
            namespace='luxor2',
            plural=CRD_PLURAL,
            resource_version=resource_version
        )
        for event in stream:
            resource_version = event['object']['metadata']['resourceVersion']
            handle_config_event(event)

if __name__ == '__main__':
    watch_http_server_configs()
