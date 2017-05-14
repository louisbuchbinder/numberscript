
const Parse = require('./Parse');
const Balanced = require('./Balanced');

class Handler {
	constructor(context) {
		this.context = context;
	}
	execPrint(character) {
		if (Parse.isNot.isPrint(character)) {
			return 0;
		}
		process.stdout.write(
			String.fromCharCode(this.context.getValue())
		);
		return 1;
	}
	execZero(character) {
		if (Parse.isNot.isZero(character)) {
			return 0;
		}
		this.context.resetValue();
		return 1;
	}
	execPointer(character) {
		if (Parse.isNot.isPointer(character)) {
			return 0;
		}
		if (Parse.isRightPointer(character)) {
			this.context.movePointer();
		}
		if (Parse.isLeftPointer(character)) {
			this.context.movePointer('LEFT');
		}
		return 1;
	}
	execNumber(character, code) {
		let index = 0;

		if ( Parse.isNot.isZero(character) && (
			Parse.isNumber(character) ||
			Parse.isNumericOperator(character)
		)) {
			let next;
			let current = character;

			if (Parse.isNumber(current)) {
				next = Parse.captureNumber(code).slice(1);
			}
			if (Parse.isNumericOperator(current)) {
				next = Parse.captureNumber(code.slice(1));
			}
			if (next) {
				current += next;
				index = current.length;
			}

			if (Parse.isNumericOperator(current)) {
				current += '1';
			}

			this.context.incrementValue(Number(current));
		}

		return index;
	}
	execLiteral(character, code) {
		let index = 0;
		let literal = '';

		if (Parse.isLiteral(character)) {
			literal = Parse.captureLiteral(code.slice(1));
			index += 1 + literal.length;
		}
		this.context.storeLiteral(literal);
		return index;
	}
	execLoop(character, code) {
		let index = 0;

		if (Parse.isLoop(character)) {
			const balancedIndex = Balanced.findBalancedIndex(code);
			while (this.context.getValue() !== 0) {
				this.context.eval(code.slice(1, balancedIndex));
			}
			index = balancedIndex;
		}
		return index;
	}
	execReference(character, code) {
		let index = 0;
		let ref = '';

		if (Parse.isReference(character)) {
			ref = Parse.captureReference(code.slice(1));
			index += 1 + ref.length + 1; // % ref .
		}

		this.context.storeReference(ref);
		return index;
	}
	execCodeReference(character) {
		if (Parse.isNot.isPeriod(character)) {
			return 0;
		}

		const Ref = this.context.getValue();

		if (typeof Ref !== 'object') {
			throw new Error('Expected to execute a code reference');
		}

		Ref.ns.eval(Ref.code);
		return 1;
	}
};

module.exports = Handler;
