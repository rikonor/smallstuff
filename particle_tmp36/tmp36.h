#ifndef TMP36_H
#define TMP36_H

const float tempSensorVoltageDivider = 4095.0;
const float tempSensorReferenceVoltage = 3.3;
const float tempSensorOffsetVoltage = 0.5;

class TMP36TemperatureSensor {
public:
  TMP36TemperatureSensor(int sensorPin);
  float temperatureCelcius();

private:
  int tempSensorPin;
};

TMP36TemperatureSensor::TMP36TemperatureSensor(int pin) {
  tempSensorPin = pin;

  pinMode(tempSensorPin, INPUT);
}

float TMP36TemperatureSensor::temperatureCelcius() {
  float tempSensorVoltage = (tempSensorReferenceVoltage * analogRead(tempSensorPin)) / tempSensorVoltageDivider;
  float tempSensorTemperature = (tempSensorVoltage - tempSensorOffsetVoltage) * 100;

  return tempSensorTemperature;
}

#endif
