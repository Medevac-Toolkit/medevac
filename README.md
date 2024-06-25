# Medevac
![image](https://github.com/Medevac-Toolkit/medevac/assets/1908752/4d0ec5bd-d9fd-46b9-b0b3-15d05c48368f)
## Overview

Medevac is a tool designed to automate the detection, analysis, and resolution of common issues in Kubernetes clusters. This tool aims to swiftly diagnose and fix problems in your containerized environments, ensuring optimal performance and uptime.

## Features

- **Command-Line Interface (CLI)**: Easily interact with the tool to diagnose and fix issues from your terminal.
- **Web UI**: Define guides and map problems to text-based solutions via an intuitive web interface.
- **Plugins**: Extensible plugins to handle common technical problems like memory leaks and cache cleaning.
- **Kubernetes Agent**: A dedicated agent running within your Kubernetes cluster to gather data and apply fixes in real-time.
- **Integration**: Seamless integration with other tools like ArgoCD and Grafana for deployment and monitoring.
- **Solution Database**: A customizable database of organization-specific solutions to streamline troubleshooting.

## Architecture

Medevac consists of several key components:

- **CLI**: Provides a command-line interface for interacting with the tool.
- **Web UI**: Allows users to define guides and map problems to solutions through a web interface.
- **Plugins**: Extensible modules that address specific technical problems.
- **Kubernetes Agent**: Runs within the cluster to monitor and fix issues in real-time.
- **Integration Modules**: Integrate with external tools like ArgoCD and Grafana.
- **Solution Database**: Stores custom solutions tailored to organizational needs.

## Status

**Work in Progress**: Medevac is currently under development. It is not recommended for production use at this stage.

## License

Medevac is released under the MIT License. See the [LICENSE](LICENSE) file for more details.
