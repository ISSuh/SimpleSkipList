/**
 *
 *  Copyright:  Copyright (c) 2023, ISSuh
 *
 */

#include <iostream>
#include <string>
#include <chrono>

#include "sdb/skip_list.h"

int32_t main() {
  sdb::SkipList<int32_t, int32_t> skip_list(8);

  auto start = std::chrono::high_resolution_clock::now();
  for (auto i = 0 ; i < 10000 ; ++i) {
    skip_list.Update(i, i);
  }
  auto stop = std::chrono::high_resolution_clock::now();

  auto duration =
    std::chrono::duration_cast<std::chrono::milliseconds>(stop - start);

  std::cout << "Update : "
            << duration.count() << " millisecond" << std::endl;

  start = std::chrono::high_resolution_clock::now();
  int32_t value;
  for (auto i = 0 ; i < 10000 ; ++i) {
    skip_list.Find(i, &value);
  }
  stop = std::chrono::high_resolution_clock::now();

  duration =
    std::chrono::duration_cast<std::chrono::milliseconds>(stop - start);

  std::cout << "Find : "
            << duration.count() << " millisecond" << std::endl;

  return 0;
}
