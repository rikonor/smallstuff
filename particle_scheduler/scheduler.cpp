#include "application.h"
#include "scheduler.h"

Scheduler::Scheduler() {}

void Scheduler::add(int d, Task t) {
  tasks.push_back(t);
  runEveryDurations.push_back(d);
  lastRunTimes.push_back(0);
}

void Scheduler::run() {
  int numOfTasks = tasks.size();

  for (int i = 0; i < numOfTasks; i++) {
    int lastRun = lastRunTimes[i];
    int runEvery = runEveryDurations[i];

    // Check if task should run
    if (millis() - lastRun > runEvery) {
      // Run task
      Task task = tasks[i];
      (*task)();

      // Update last run time for task
      lastRunTimes[i] = millis();
    }
  }
}
