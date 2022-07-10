# Tion Breazer 3S Go API and Clients
There are thee backend available:
  - https://github.com/m-pavel/go-gattlib
  - https://github.com/paypal/gatt
  - https://github.com/muka/go-bluetooth <- Preferable
## Pairing
Device must be paired. E.g. using bluetoothctrl
## Reading state
```
    t = tionm.New("MAC")
    ts.t.Connect(timeout)
    defer ts.t.Disconnect()
    s, err := ts.t.ReadState(timeout)
    fmt.Println(s)
```
# Tion Breazer Home Assistant MQTT integration
## Sensors
```
- platform: mqtt
  name: "Temperature Inside (Tion)"
  state_topic: "nn/tion"
  value_template: "{{ value_json.temp_out }}"
  availability_topic: "nn/tion-aval"
  icon: "mdi:thermometer"
  unit_of_measurement: '°C'

- platform: mqtt
  name: "Temperature Outside (Tion)"
  state_topic: "nn/tion"
  value_template: "{{ value_json.temp_in }}"
  availability_topic: "nn/tion-aval"
  icon: "mdi:thermometer"
  unit_of_measurement: '°C'

- platform: mqtt
  name: "Temperature Target (Tion)"
  state_topic: "nn/tion"
  value_template: "{{ value_json.temp_target }}"
  availability_topic: "nn/tion-aval"
  icon: "mdi:thermometer"
  unit_of_measurement: '°C'

- platform: mqtt
  name: "Speed (Tion)"
  state_topic: "nn/tion"
  value_template: "{{ value_json.speed }}"
  availability_topic: "nn/tion-aval"
  icon: "mdi:fan"

```
## Control channel
Turn on/off
```
    {
      "payload_template": "{% if is_state('binary_sensor.tion' , 'off') %} \n  { \"on\": true }\n{% else %}\n  { \"on\": false }\n{% endif %}\n",
      "qos": 1,
      "topic": "nn/tion-control"
    }
```
## Building
```
    make GOARCH=<target arch> IMPL=<backend implementation>
```
GOARCH - https://golang.org/doc/install/source#environment

IMPL - one of below
  - muka https://github.com/muka/go-bluetooth backend implementation, default
  - ppal
  - gatt
  - fake
  - mqtt
## Magic
 When getting 'Unexpected response length' with muka driver, run once with ppal driver
 In fact it performs adaptor reset like below
 ```
 sudo hciconfig hci0 up
 sudo hciconfig hci0 reset
 sudo hciconfig hci0 down
 sudo hciconfig hci0 up
```

## Installation
 1. Install dependencies.
    - Go toolchain required if you are going to compile from sources.
    - [Gattlib](https://github.com/m-pavel/go-gattlib) required if you are using gatt backend.
    - Install [Bluez](http://www.bluez.org/) to scan and pair device.
 2. [Compile](INSTALL.md) or download [released](https://github.com/m-pavel/go-tion/releases) binaries.
 3. Pair device with bluetoothctrl
    First, press 'On' button on device for 5 seconds while blue ligth is turned on.
    Then pair with bluetoothctrl
     ```
        [bluetooth]# scan on
        ...
        [NEW] Device E1:AA:BB:CC:DD:EE Tion Breezer 3S
        ...
        [bluetooth]# scan off
        [bluetooth]# pair E1:AA:BB:CC:DD:EE

    ```
4. Now you can use tion tools
   - cli :
   ```
    $ ./tion-cli --device E1:AA:BB:CC:DD:EE
    Using implementation github.com/muka/go-bluetooth/api
    Status: on, Heater: on, Sound: on
    Target: 15 C, In: 3 C, Out: 14 C
    Speed 2, Gate: outdoor, Error: 0, FW: 32
    Filters remain: 42 days, Uptime 280 days 17:19

    $ ./tion-cli --device E1:AA:BB:CC:DD:EE -off
    Turned off

    $ ./tion-cli --device E1:AA:BB:CC:DD:EE -on
    Turned on
    ```
   - mqtt can be controlled by systemd, below example configuration
   ```go
   [Unit]
   Description=Tion MQTT

   [Service]
   User=pi
   ExecStart=/pathto/tion-mqtt -device  E1:AA:BB:CC:DD:EE -mqtt ssl://mqtt -mqtt-user user -mqtt-pass password -mqtt-ca /path-to-ca -n -d -keepbt -interval 30
   Restart=always
   RestartSec=20
   StartLimitInterval=0

   [Install]
   WantedBy=multi-user.target
   ```
   MQTT client publish tion status to topic periodically and can be managed by control messages from control topic.
   - influx can send tion parameters to [InfluxDB](https://www.influxdata.com/)
   - schedule can configure and run actions by timer

