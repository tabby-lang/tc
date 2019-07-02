#include <string>
#include <iostream>
#include "Builtins.cpp"

Int add(Int b,Int a) {
Int tmp_1 = a.ADD(b);
return tmp_1;
}

int main() {
Int tmp_2 = Int(2);
Int tmp_3 = Int(-2147483648);
Int tmp_4 = add(tmp_2,tmp_3);Int num = tmp_4;
Nothing tmp_5 = num.println();tmp_5;
return 0;
}