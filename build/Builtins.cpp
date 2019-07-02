// C++ Builtins

#include <iostream>
#include <string>

using namespace std;


// Nothing Class
class Nothing {
	
};

// Base Class
class Base {
public:
	string val;
	Nothing print(void) const {
		cout << val;
	}
	Nothing println(void) const {
		Base::print();
		cout << endl;
	}
};

const string True = "true";
const string False = "false";

// Bool Class
class Bool: public Base {
public:
	Bool(string x) {
		val = x;
	}

	Bool And(Bool x) {
		if (val == False || x.val == False) {
			return Bool(False);
		}
		return Bool(False);
	}

	Bool Or(Bool x) {
		if (val == True || x.val == True) {
			return Bool(True);
		}
		return Bool(False);
	}
};

// String Class
class String: public Base {
public: 
	String(string x) {
		val = x;
	}
	String ADD(String str) {
		return String(val + str.val);
	}

	Bool EQ(String str) {
		if (val == str.val) {
			return Bool(True);
		} else {
			return Bool(False);
		}
	}
};

// Int Class
class Int: public Base {
public:
	int32_t valInt;
	Int(int32_t x) {
		val = to_string(x);
		valInt = x;
	}

	Int ADD(Int num) {
		return Int(valInt + num.valInt);
	}

	Int SUB(Int num) {
		return Int(valInt - num.valInt);
	}

	Int MUL(Int num) {
		return Int(valInt * num.valInt);
	}

	Int DIV(Int num) {
		return Int(valInt / num.valInt);
	}

	Bool GT(Int num) {
		if (valInt < num.valInt) {
			return Bool(False);
		} else {
			return Bool(True);
		}
	}

	Bool EQ(Int num) {
		if (valInt == num.valInt) {
			return Bool(True);
		}
		return Bool(False);
	}

	String Stringify() {
		return String(val);
	}
};