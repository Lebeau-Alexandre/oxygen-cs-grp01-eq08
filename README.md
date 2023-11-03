# LOG-680 : Template for Oxygen-CS

![image](./doc/wheel.png)

This Python application continuously monitors a sensor hub and manages HVAC (Heating, Ventilation, and Air Conditioning) system actions based on received sensor data.

It leverages `signalrcore` to maintain a real-time connection to the sensor hub and utilizes `requests` to send GET requests to a remote HVAC control endpoint.

This application uses `pipenv`, a tool that aims to bring the best of all packaging worlds to the Python world.

## Requirements

- Golang 1.18+

## Getting Started

Cloning the repository :

```
git clone https://github.com/Lebeau-Alexandre/oxygen-cs-grp01-eq08.git
cd log680
```

## Setup

You need to setup the following variables inside the Main class:

- HOST: The host of the sensor hub and HVAC system.
- TOKEN: The token for authenticating requests.
- T_MAX: The maximum allowed temperature.
- T_MIN: The minimum allowed temperature.
- DATABASE: The database connection details.

## Running the Program

After setup, you can start the program with the following command:

```
go run .
```

## Logging

The application logs important events such as connection open/close and error events to help in troubleshooting.

## To Implement

There are placeholders in the code for sending events to a database and handling request exceptions. These sections should be completed as per the requirements of your specific application.

## Intégration continue

### Pre-commit git hook

Les pre-commit hooks ont été implémenté, dans notre cas, grâce à [Husky](https://github.com/automation-co/husky). Cet outil nous a permis de rouler nos tests et analyser le formatage _(linting)_ du code.

### CI/CD

Un workflow Git a été créée pour être exécuté sur la branche `main` pour chaque _push_ ou _pull request_. Celui-ci démarre les tests et le _lint_ simultanément pour ensuite créer et déployer [l'image Docker](https://hub.docker.com/repository/docker/lebeauets/oxygen/general).  


## License

MIT

## Contact

For more information, please feel free to contact the repository owner.
