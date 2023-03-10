# Lap timing

## Table of contents
- [Lap timing](#lap-timing)
  - [Table of contents](#table-of-contents)
  - [Project description](#project-description)
  - [Requirements](#requirements)
    - [Hardware](#hardware)
    - [Software](#software)
  - [Hardware configuration](#hardware-configuration)
    - [Pit entry/exit](#pit-entryexit)
    - [Start](#start)
    - [Car](#car)
  - [Project layout](#project-layout)
  - [Run instructions](#run-instructions)
  - [Links](#links)
  - [Team members and roles](#team-members-and-roles)

## Project description
This project aims to detect cars passing through a road acquiring the laptime, and also to signal when a car enters or exits the pit. The signaling between the cars and the "ground boards" is done using infrared light. The data is collected by a main board and sent to a backend that will save it on Firebase, as well as send a notification on Telegram.

## Requirements

### Hardware
> Assume *n* as the number of cars, for the demonstration project *n*=2

Boards required:
- 2 Arduino of any kind
- 1 Esp8266 (NodeMCU) or any other wifi module
- 1 Arduino for each car (*n*)
  
Boards used:
- 1 Arduino Leonardo    *[Pit entry]*
- 2 Arduino promicro    *[Pit exit, car]*
- 1 Arduino Nano        *[car]*
- 1 Esp8266 (NodeMCU)   *[Start]*


Passive components:
- 2 1k Ω resistor (could be avoided by using a software serial)
- 2 2k Ω resistor (could be avoided by using a software serial)
- 3+(*n*) IR transmitter
- 3+(*n*) IR receiver
- *n* rgb leds


### Software

- PlatformIO (for the embedded boards)
- Go (for the backend)
- Firebase account (for the database)
- Protobuf compiler (for the protobuf files) (not required if you use the precompiled files)
- Python to compile protobuf files into nanopb files

## Hardware configuration
### Pit entry/exit
The pit entry and the pit exit boards recive the signal transmitted from the car and send it to the start board (via UART 5v dropped to 3.3v). The pit entry board also sends an IR signal that should be recived by the car to change it's status.
![Scheme](img/pit_entry_scheme.png)

### Start

The start board recives the signal from the pit entry and the pit exit boards (UART 3.3v). It receives the IR signal with the `car_id_code` from the car and sends an univoque code to the car.

![Scheme](img/start_scheme.png)

### Car
The car constantly transmits his `car_id_code` (univoque for each car) via IR to the various boards, when it reads the codes sent by the different "ground boards" it changes it's status and changes the led color.
![Scheme](img/car_scheme.png)


## Project layout

```bash
.
├── backend                     #Source code of the backend, written in Go
│   ├── firebase_interface
│   ├── laptiming               #Proto and laptiming routes
│   ├── main.go
│   ├── secrets                 #Backend secrets
│   └── telegram_interface      #Telegram bot interface
├── embedded_code               #Source code of the embedded boards, written in C++
│   ├── car                     #Code for the car boards
│   │   ├── include
│   │   ├── lib
│   │   ├── platformio.ini
│   │   └── src
│   ├── pit                     #Code for the pit boards  
│   │   ├── include
│   │   ├── lib
│   │   ├── platformio.ini
│   │   └── src
│   └── start                   #Code for the start board
│       ├── include
│       │   ├── http      
│       │   ├── nanopb
│       │   ├── proto
│       ├── lib
│       ├── platformio.ini
│       └── src
│           ├── http
│           ├── nanopb
│           └── proto
│── img                         #README images
├── powerpoint                  #presentation
└── proto                       #Protobuf files
    ├── compile.sh
    ├── messages.proto
    └── nanopb
```

## Run instructions

Start the backend:
- Insert the secrets in the `backend/secrets` folder (such as the json file for the firebase account)
  - The secrets must be named `firebase.json` and `telegram.env`
- Run `go run backend/main.go`

Flash the boards:
- Go to the desired folder inside `embedded_code` (e.g. `embedded_code/car`)
- `pio run -e <environment> -t build` (where `<environment>` is the name of the environment in the `platformio.ini` file)
- `pio run -e <environment> -t upload` 

Chage the protobuf definition:
- Write the messages on  `proto/messages.proto`
- Run the `proto/compile.sh` command
- Go in the `proto` folder and run `protoc -o messages.pb messages.proto`
- Then `python nanopb/generator/nanopb_generator.py message.pb`
- Copy the `messages.pb.h` and `messages.pb.cc` inside the `embedded_code/start/` folders (`include` and `src` accordingly)

## Links

- [Presentation](https://docs.google.com/presentation/d/1icwILWqBtrO3eo8dQbCCBscmbWkpSkbpfpHDUgP1MSg/edit?usp=sharing)
- [Video](https://youtu.be/P7rQA0mt8G8)

## Team members and roles

Team members:
- Piccinin Diego [mat. 217776]
- Pivetta Leonardo [mat. 217395]

Roles:
- **Wiring**: Piccinin Diego;
- **Backend**: Leonardo Pivetta;
- **Embedded code**: both of us, we live together so we built almost everything together.
  - Serial communication: mainly Leonardo Pivetta;
  - IR communication: mainly Diego Piccinin;