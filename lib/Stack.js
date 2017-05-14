
class Stack {
	constructor() {
		this.initialize();
	}
	initialize() {
		this.stack = [0];
		this.pointer = 0;
	}
	getPointer() {
		return this.pointer;
	}
	getStack() {
		return this.stack;
	}
	getValue() {
		return this.getStack()[this.getPointer()];
	}
	movePointer(left) {
		if (left) {
			this.pointer -= 1
		} else {
			this.pointer += 1
		}

		if (this.pointer < 0) {
			this.initialize();
		}

		this.checkStackForInitialValue();

		return this.pointer;
	}
	checkStackForInitialValue() {
		const value = this.getValue();

		if (typeof value === 'object' && value.ns) {
			return void 0;
		}
		if (isNaN(value)) {
			this.stack[this.getPointer()] = 0;
		}
	}
	incrementValue(value) {
		if (isNaN(value)) {
			throw new Error('Cannot increment the value by a non-number');
		}
		if (!Number.isInteger(value)) {
			throw new Error('Cannot increment the value by a non-integer');
		}
		return this.stack[this.getPointer()] += value;
	}
	storeLiteral(literal) {
		if (typeof literal !== 'string') {
			throw new Error('Expected a literal string');
		}

		let string = literal;

		while (string) {
			this.incrementValue(string[0].charCodeAt(0));
			this.movePointer();
			string = string.slice(1);
		}
		if  (literal.length) { // reset pointer to last literal char
			this.movePointer('LEFT');
		}

	}
	storeReference(string) {
		if (typeof string !== 'string') {
			throw new Error('Expected a string of code to reference');
		}
		const NS = require('..');

		if (string) {
			this.stack[this.getPointer()] = {
				code: string,
				ns: new NS
			};
		}
	}
	resetValue() {
		return this.stack[this.getPointer()] = 0;
	}
};

module.exports = Stack;
