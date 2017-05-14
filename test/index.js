
const assert = require('assert');
const equal = assert.deepStrictEqual;
const sinon = require('sinon');

describe('Interpreter Unit Tests', () => {
	const Interpreter = require('..');
	let ns;

	beforeEach(() => ns = new Interpreter);

	it('should be a class constructor', () => {
		equal(typeof ns, 'object');
	});

	it('should have a getStack method', () => {
		equal(typeof ns.getStack, 'function');
	});

	it('should initialize the stack to an INT array', () => {
		equal(ns.getStack(), [0]);
	});

	it('should have a getPointer method', () => {
		equal(typeof ns.getPointer, 'function');
	});

	it('should initialize the pointer to 0', () => {
		equal(ns.getPointer(), 0);
	});

	it('should have a getValue method', () => {
		equal(typeof ns.getValue, 'function');
	});

	it('should initialize the value to 0', () => {
		equal(ns.getValue(), 0);
	});

	it('should have a movePointer method', () => {
		equal(typeof ns.movePointer, 'function');
	});

	it('should move the pointer to the right with a falsy argument', () => {
		equal(ns.movePointer(), 1);
		equal(ns.getPointer(), 1);
	});

	it('should move the pointer to the left with a truthy argument', () => {
		equal(ns.movePointer('left'), 0);
		equal(ns.getPointer(), 0);
	});

	it('should not move the pointer to the left of 0', () => {
		equal(ns.movePointer('left'), 0);
		equal(ns.getPointer(), 0);
	});

	it('should initialize new items in the stack to 0', () => {
		ns.movePointer();
		equal(ns.getValue(), 0);
		ns.movePointer('left');
	});

	it('should have an incrementValue method', () => {
		equal(typeof ns.incrementValue, 'function');
	});

	it('should increment the value by the integer argument', () => {
		equal(ns.incrementValue(10), 10);
		equal(ns.getValue(), 10);
		equal(ns.incrementValue(-10), 0);
		equal(ns.getValue(), 0);
	});

	it('should throw an error when incrementing by a non-number', () => {
		assert.throws(() => ns.incrementValue());
	});

	it('should throw an error when incrementing by a non-integer', () => {
		assert.throws(() => ns.incrementValue(3.14));
	});

	it('should have an eval method', () => {
		equal(typeof ns.eval, 'function');
	});

	describe('Numberscript Eval Unit Tests', () => {
		const logSpy = sinon.spy(process.stdout, 'write');
		let ns;

		beforeEach(() => {
			ns = new Interpreter;
			logSpy.reset();
		});

		function concatArgs(spy) {
			let args = '';
			let index = 0;
			let current;

			while (current = spy.args[index++]) {
				args += current.join``;
			}

			return args;
		}

		it('should print "0"', () => {
			ns.eval('48π');
			equal(logSpy.calledOnce, true);
			equal(logSpy.alwaysCalledWith('0'), true);
			equal(concatArgs(logSpy), '0');
		});

		it('should print "A"', () => {
			ns.eval('65π');
			equal(logSpy.calledOnce, true);
			equal(logSpy.alwaysCalledWith('A'), true);
			equal(concatArgs(logSpy), 'A');
		});

		it('should print "a"', () => {
			ns.eval('97π');
			equal(logSpy.calledOnce, true);
			equal(logSpy.alwaysCalledWith('a'), true);
			equal(concatArgs(logSpy), 'a');
		});

		it('should print "0Aa"', () => {
			ns.eval('48π17π32π');
			equal(logSpy.calledThrice, true);
			equal(logSpy.firstCall.calledWith('0'), true);
			equal(logSpy.secondCall.calledWith('A'), true);
			equal(logSpy.thirdCall.calledWith('a'), true);
			equal(concatArgs(logSpy), '0Aa');
		});

		it('should print "ZA"', () => {
			ns.eval('90π-25π');
			equal(logSpy.calledTwice, true);
			equal(concatArgs(logSpy), 'ZA');
		});

		it('should print "ZYZ"', () => {
			ns.eval('90π-π+π');
			equal(logSpy.calledThrice, true);
			equal(concatArgs(logSpy), 'ZYZ');
		});

		it('should move the pointer', () => {
			equal(ns.getPointer(), 0);
			ns.eval('>');
			equal(ns.getPointer(), 1);
			ns.eval('<');
			equal(ns.getPointer(), 0);
		});

		it('should print "0Aa"', () => {
			ns.eval('97>65>48π<π<π');
			equal(logSpy.calledThrice, true);
			equal(concatArgs(logSpy), '0Aa');
		});

		it('should skip loops when the value is 0', () => {
			ns.eval('(48π)49π');
			equal(logSpy.calledOnce, true);
			equal(concatArgs(logSpy), '1');
		});

		it('should looop until the value is 0', () => {
			ns.eval('48>3(<π>-1)');
			equal(logSpy.calledThrice, true);
			equal(concatArgs(logSpy), '000');
		});

		it('should looop until the value is 0', () => {
			ns.eval('48>10(<π+>-)<7>26(<π+>-)<6>26(<π+>-)');
			equal(logSpy.callCount, 62);
			equal(concatArgs(logSpy), '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz');
		});

		it('should print Hello World!', () => {
			ns.eval('#!dlroW olleH(π<)');
			equal(concatArgs(logSpy), 'Hello World!');
		});

	});

});
