Simple scheduler for Particle devices
---

#### Usage

```
#include "scheduling.h";

void fOne();
void fTwo();
void fThree();
void fFour();

Scheduler sch = Scheduler();

void setup() {
  sch.add(1200, &fOne);
  sch.add(3000, &fTwo);
  sch.add(4500, &fThree);
  sch.add(500, &fFour);
}

void loop() {
  sch.run();
}

void fOne() {
  Serial.println("1");
}

void fTwo() {
  Serial.println("2");
}

void fThree() {
  Serial.println("3");
}

void fFour() {
  Serial.println("4");
}
```
