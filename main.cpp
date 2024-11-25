#include "tstring.h"
#include <iostream>

const TagLib::String a("hello");
const std::string b = "hello";

__attribute__((export_name("do_debug"))) void do_debug() {
  std::cout << "a: " << a << std::endl;
  std::cout << "b: " << b << std::endl;

  std::cout << std::endl;

  std::cout << "a size: " << a.size() << std::endl;

  const auto &data = a.data(TagLib::String::UTF8);
  for (size_t i = 0; i < data.size(); i++)
    std::cout << "a byte " << i << ": " << (int)data[i] << std::endl;
}

int main() { return 0; };
