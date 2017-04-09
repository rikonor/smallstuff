#include "application.h"
#include "debug_led.h"

DebugLed::DebugLed(int lp) {
  ledPin = lp;
  pinMode(ledPin, OUTPUT);
}

void DebugLed::registerMsg(String msg, int times) {
  msgCounts[msg] = times;
}

void DebugLed::triggerMsg(String msg) {
  std::map<String, int>::iterator it;

  it = msgCounts.find(msg);
  if (it != msgCounts.end()) {
    this->blink(it->second);
  }
}

void DebugLed::blink(int times) {
  for (int i = 0; i < times; i++) {
    digitalWrite(ledPin, HIGH);
    delay(200);
    digitalWrite(ledPin, LOW);
    delay(200);
  }
}
