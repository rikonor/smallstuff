#ifndef DEBUG_LED_H
#define DEBUG_LED_H

#include <map>;

class DebugLed {
public:
  DebugLed(int ledPin);
  void blink(int times);
  void registerMsg(String msg, int times);
  void triggerMsg(String msg);

private:
  int ledPin;
  std::map<String, int> msgCounts;
};

#endif
