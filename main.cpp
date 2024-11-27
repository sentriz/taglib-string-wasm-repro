#include <iostream>

const auto EXAMPLE = std::make_shared<std::wstring>(L"hello");

std::string convertString(std::wstring &input) {
  std::string result;
  result.resize(input.size() * 4);

  size_t pos = 0;
  const wchar_t *data_start = input.data();
  const wchar_t *data_end = data_start + input.size();

  while (data_start != data_end) {
    result[pos] = static_cast<char>(*data_start++);
    pos++;
  }

  result.resize(pos);
  return result;
}

__attribute__((export_name("do_debug"))) void do_debug() {
  auto s = convertString(*EXAMPLE);
  std::cout << "b: " << s << std::endl;
  std::cout << "b size: " << s.size() << std::endl;
}

int main() {
  auto s = convertString(*EXAMPLE);
  std::cout << "b: " << s << std::endl;
  std::cout << "b size: " << s.size() << std::endl;
  return 0;
};
