#ifndef SCHEDULER_H
#define SCHEDULER_H

#include <vector>

typedef void (*Task)();

class Scheduler {
public:
  Scheduler();
  void add(int d, Task t);
  void run();

private:
  std::vector<int> runEveryDurations;
  std::vector<int> lastRunTimes;
  std::vector<Task> tasks;
};

#endif
