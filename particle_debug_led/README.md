Debug Led lib
---

#### Usage

```
#include "debug_led.h";

DebugLed debugLed = DebugLed(D7);

void setup() {
  debugLed.registerMsg("starting_loop", 1);
  debugLed.registerMsg("finished_loop", 2);
}

void loop() {
  debugLed.triggerMsg("starting_loop");
  // Do Stuff ...
  debugLed.triggerMsg("finished_loop");
}
```
