package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	PAD "github.com/pinpt/go-common/strings"
)

func hexToB64(h string) (string, error) {
	p, err := hex.DecodeString(h)
	if err != nil {
		log.Fatal(err)
	}
	b64 := base64.StdEncoding.EncodeToString([]byte(p))
	return b64, err
}

func randSensorData(sensorType string, i int) string {
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder

	switch sensorType {
	case "temperature":
		sb.WriteString("01")
		min := 0
		max := 400
		v := PAD.PadLeft((strconv.FormatUint(uint64((rand.Intn(max-min+1) + min)), 16)), 4, 0x30)
		sb.WriteString(v)

	case "humidity":
		sb.WriteString("02")
		min := 0
		max := 100
		v := PAD.PadLeft((strconv.FormatUint(uint64((rand.Intn(max-min+1) + min)), 16)), 4, 0x30)
		sb.WriteString(v)

	case "boardVoltage":
		sb.WriteString("0c")
		min := 0
		max := 4200
		v := PAD.PadLeft((strconv.FormatUint(uint64((rand.Intn(max-min+1) + min)), 16)), 4, 0x30)
		sb.WriteString(v)

	case "counter":
		sb.WriteString("0b")
		min := 0
		// max := 16777215
		v := PAD.PadLeft((strconv.FormatUint(uint64(min+i), 16)), 6, 0x30)
		sb.WriteString(v)

	case "counter_0d":
		sb.WriteString("0d")
		min := 0
		max := 4096
		v := PAD.PadLeft((strconv.FormatUint(uint64((rand.Intn(max-min+1) + min)), 16)), 4, 0x30)
		sb.WriteString(v)

	case "distance":
		sb.WriteString("13")
		min := 0
		max := 4096
		v := PAD.PadLeft((strconv.FormatUint(uint64((rand.Intn(max-min+1) + min)), 16)), 4, 0x30)
		sb.WriteString(v)

	case "pressure":
		sb.WriteString("03")
		min := 0
		max := 4096
		v := PAD.PadLeft((strconv.FormatUint(uint64((rand.Intn(max-min+1) + min)), 16)), 4, 0x30)
		sb.WriteString(v)
	}
	return sb.String()
}

func randDeviceData(deviceType string, i int) (string, string) {
	type deviceInfo struct {
		nodeName string
		devEUI   string
	}
	var sb strings.Builder
	var stb strings.Builder

	sb.WriteString(`{"applicationID":"`)

	switch deviceType {
	case "SmartLights":
		var sdb strings.Builder // string data builder
		deviceInfos := []deviceInfo{
			{nodeName: "SmartLight_1", devEUI: "0004a30b00000001"},
			{nodeName: "SmartLight_2", devEUI: "0004a30b00000002"},
			{nodeName: "SmartLight_3", devEUI: "0004a30b00000003"},
			{nodeName: "SmartLight_4", devEUI: "0004a30b00000004"},
			{nodeName: "SmartLight_5", devEUI: "0004a30b00000005"},
			{nodeName: "SmartLight_6", devEUI: "0004a30b00000006"},
			{nodeName: "SmartLight_7", devEUI: "0004a30b00000007"},
		}
		j := float64(i)
		if i >= len(deviceInfos) {
			j = math.Mod(float64(i), float64(len(deviceInfos)))
		}
		k := uint(j)

		sb.WriteString(`17","applicationName":"SmartCampusMaua","nodeName":"`)
		sb.WriteString(deviceInfos[k].nodeName)
		sb.WriteString(`","devEUI":"`)
		sb.WriteString(deviceInfos[k].devEUI)
		sb.WriteString(`","rxInfo":[{"mac":"7276ff000b031df7","time":"2024-03-28T22:33:26.728817Z","rssi":-97,"loRaSNR":11.2,"name":"IMT-kerlink-blocoW2","latitude":-23.64946,"longitude":-46.57367,"altitude":777},{"mac":"7276ff00080801db","time":"2024-03-28T22:33:25Z","rssi":-100,"loRaSNR":11.5,"name":"IMT-kerlink-blocoH","latitude":-23.64891,"longitude":-46.57413,"altitude":758}],"txInfo":{"frequency":916200000,"dataRate":{"modulation":"LORA","bandwidth":125,"spreadFactor":9},"adr":true,"codeRate":"4/5"},"fCnt":135,"fPort":100,"data":"`)
		sdb.WriteString(randSensorData("counter", i))
		sdb.WriteString(randSensorData("counter_0d", i))
		sdb.WriteString(randSensorData("counter_0d", i))
		sdb.WriteString(randSensorData("boardVoltage", i))
		data, _ := hexToB64(sdb.String())
		sb.WriteString(data)
		sb.WriteString(`"}`)
		stb.WriteString(`application/17/node/`)
		stb.WriteString(deviceInfos[k].devEUI)
		stb.WriteString(`/rx`)

	case "WaterTankLevels":
		var sdb strings.Builder // string data builder
		deviceInfos := []deviceInfo{
			{nodeName: "WaterTankLevel_1", devEUI: "0004a30b00001001"},
			{nodeName: "WaterTankLevel_2", devEUI: "0004a30b00001002"},
			{nodeName: "WaterTankLevel_3", devEUI: "0004a30b00001003"},
			{nodeName: "WaterTankLevel_4", devEUI: "0004a30b00001004"},
			{nodeName: "WaterTankLevel_5", devEUI: "0004a30b00001005"},
			{nodeName: "WaterTankLevel_6", devEUI: "0004a30b00001006"},
			{nodeName: "WaterTankLevel_7", devEUI: "0004a30b00001007"},
			{nodeName: "WaterTankLevel_8", devEUI: "0004a30b00001008"},
		}
		j := float64(i)
		if i >= len(deviceInfos) {
			j = math.Mod(float64(i), float64(len(deviceInfos)))
		}
		k := uint(j)

		sb.WriteString(`17","applicationName":"SmartCampusMaua","nodeName":"`)
		sb.WriteString(deviceInfos[k].nodeName)
		sb.WriteString(`","devEUI":"`)
		sb.WriteString(deviceInfos[k].devEUI)
		sb.WriteString(`","rxInfo":[{"mac":"7276ff000b031df7","time":"2024-03-28T22:33:26.728817Z","rssi":-97,"loRaSNR":11.2,"name":"IMT-kerlink-blocoW2","latitude":-23.64946,"longitude":-46.57367,"altitude":777},{"mac":"7276ff00080801db","time":"2024-03-28T22:33:25Z","rssi":-100,"loRaSNR":11.5,"name":"IMT-kerlink-blocoH","latitude":-23.64891,"longitude":-46.57413,"altitude":758}],"txInfo":{"frequency":916200000,"dataRate":{"modulation":"LORA","bandwidth":125,"spreadFactor":9},"adr":true,"codeRate":"4/5"},"fCnt":135,"fPort":100,"data":"`)
		sdb.WriteString(randSensorData("distance", i))
		sdb.WriteString(randSensorData("boardVoltage", i))
		data, _ := hexToB64(sdb.String())
		sb.WriteString(data)
		sb.WriteString(`"}`)
		stb.WriteString(`application/17/node/`)
		stb.WriteString(deviceInfos[k].devEUI)
		stb.WriteString(`/rx`)

	case "Hidrometers":
		var sdb strings.Builder // string data builder
		deviceInfos := []deviceInfo{
			{nodeName: "Hidrometer_1", devEUI: "0004a30b00101001"},
			{nodeName: "Hidrometer_2", devEUI: "0004a30b00101002"},
			{nodeName: "Hidrometer_3", devEUI: "0004a30b00101003"},
			{nodeName: "Hidrometer_4", devEUI: "0004a30b00101004"},
			{nodeName: "Hidrometer_5", devEUI: "0004a30b00101005"},
			{nodeName: "Hidrometer_6", devEUI: "0004a30b00101006"},
			{nodeName: "Hidrometer_7", devEUI: "0004a30b00101007"},
			{nodeName: "Hidrometer_8", devEUI: "0004a30b00101008"},
		}
		j := float64(i)
		if i >= len(deviceInfos) {
			j = math.Mod(float64(i), float64(len(deviceInfos)))
		}
		k := uint(j)

		sb.WriteString(`17","applicationName":"SmartCampusMaua","nodeName":"`)
		sb.WriteString(deviceInfos[k].nodeName)
		sb.WriteString(`","devEUI":"`)
		sb.WriteString(deviceInfos[k].devEUI)
		sb.WriteString(`","rxInfo":[{"mac":"7276ff000b031df7","time":"2024-03-28T22:33:26.728817Z","rssi":-97,"loRaSNR":11.2,"name":"IMT-kerlink-blocoW2","latitude":-23.64946,"longitude":-46.57367,"altitude":777},{"mac":"7276ff00080801db","time":"2024-03-28T22:33:25Z","rssi":-100,"loRaSNR":11.5,"name":"IMT-kerlink-blocoH","latitude":-23.64891,"longitude":-46.57413,"altitude":758}],"txInfo":{"frequency":916200000,"dataRate":{"modulation":"LORA","bandwidth":125,"spreadFactor":9},"adr":true,"codeRate":"4/5"},"fCnt":135,"fPort":100,"data":"`)
		sdb.WriteString(randSensorData("counter", i))
		sdb.WriteString(randSensorData("boardVoltage", i))
		data, _ := hexToB64(sdb.String())
		sb.WriteString(data)
		sb.WriteString(`"}`)
		stb.WriteString(`application/17/node/`)
		stb.WriteString(deviceInfos[k].devEUI)
		stb.WriteString(`/rx`)

	case "ArtesianWells":
		var sdb strings.Builder // string data builder
		deviceInfos := []deviceInfo{
			{nodeName: "ArtesianWell_1", devEUI: "0004a30b0040100c"},
		}
		j := float64(i)
		if i >= len(deviceInfos) {
			j = math.Mod(float64(i), float64(len(deviceInfos)))
		}
		k := uint(j)

		sb.WriteString(`17","applicationName":"SmartCampusMaua","nodeName":"`)
		sb.WriteString(deviceInfos[k].nodeName)
		sb.WriteString(`","devEUI":"`)
		sb.WriteString(deviceInfos[k].devEUI)
		sb.WriteString(`","rxInfo":[{"mac":"7276ff000b031df7","time":"2024-03-28T22:33:26.728817Z","rssi":-97,"loRaSNR":11.2,"name":"IMT-kerlink-blocoW2","latitude":-23.64946,"longitude":-46.57367,"altitude":777},{"mac":"7276ff00080801db","time":"2024-03-28T22:33:25Z","rssi":-100,"loRaSNR":11.5,"name":"IMT-kerlink-blocoH","latitude":-23.64891,"longitude":-46.57413,"altitude":758}],"txInfo":{"frequency":916200000,"dataRate":{"modulation":"LORA","bandwidth":125,"spreadFactor":9},"adr":true,"codeRate":"4/5"},"fCnt":135,"fPort":100,"data":"`)
		sdb.WriteString(randSensorData("pressure", i))
		sdb.WriteString(randSensorData("pressure", i))
		sdb.WriteString(randSensorData("boardVoltage", i))
		data, _ := hexToB64(sdb.String())
		sb.WriteString(data)
		sb.WriteString(`"}`)
		stb.WriteString(`application/17/node/`)
		stb.WriteString(deviceInfos[k].devEUI)
		stb.WriteString(`/rx`)
	}
	return stb.String(), sb.String()

}

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(1000 * time.Second)
		done <- true
	}()

	id := uuid.New().String()
	var sbMqtt strings.Builder
	sbMqtt.WriteString("mqtt-")
	sbMqtt.WriteString(id)
	i := 0

	// MQTT
	broker := "mqtt://localhost:1883"
	mqttClientId := sbMqtt.String()
	qos := 0

	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(mqttClientId)

	choke := make(chan [2]string)

	opts.SetDefaultPublishHandler(func(mqttClient MQTT.Client, msg MQTT.Message) {
		choke <- [2]string{msg.Topic(), string(msg.Payload())}
	})

	mqttClient := MQTT.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		select {
		case <-done:
			fmt.Println("Done!")
			mqttClient.Disconnect(250)
			fmt.Println("Sample Publisher Disconnected")
			return
		case t := <-ticker.C:
			// fmt.Printf("\n%s,\nSent payload: %v\n\n", t, payload[rand.Intn(len(payload))])
			topic, payload := randDeviceData("WaterTankLevels", i)
			mqttClient.Publish(topic, byte(qos), false, payload)
			topic, payload = randDeviceData("Hidrometers", i)
			mqttClient.Publish(topic, byte(qos), false, payload)
			topic, payload = randDeviceData("ArtesianWells", i)
			mqttClient.Publish(topic, byte(qos), false, payload)
			topic, payload = randDeviceData("SmartLights", i)
			mqttClient.Publish(topic, byte(qos), false, payload)
			fmt.Printf("\nt: %v\n\n", t.UnixNano())
			i++
			// token.Wait()

		}
	}
}
