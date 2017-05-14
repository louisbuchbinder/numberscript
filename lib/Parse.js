
class Parse {
	static isZero(character) {
		return character === '0';
	}
	static isNumber(character) {
		return Number.isInteger(Number(character));
	}
	static isNumericOperator(character) {
		return ['+', '-'].some(c => c === character);
	}
	static isPrint(character) {
		return character === 'π';
	}
	static isPointer(character) {
		return ['<', '>'].some(c => c === character);
	}
	static isRightPointer(character) {
		return character === '>';
	}
	static isLeftPointer(character) {
		return character === '<';
	}
	static isLoop(character) {
		return character === '(';
	}
	static isLiteral(character) {
		return character === '#';
	}
	static isPeriod(character) {
		return character === '.';
	}
	static isReference(character) {
		return character === '%';
	}
	static isSpecial(character) {
		return ['>', '<', '('].some(c => c === character);
	}
	static captureNumber(string) {
		let numberStr = '';
		let index = 0;
		let next;

		while (Parse.isNumber(next = string[index++])) {
			numberStr += next
		}

		return numberStr;
	}
	static captureLiteral(string) {
		let literal = '';
		let index = 0;
		let next;

		while (Parse.isNot.isSpecial(next = string[index++])) {
			literal += next;
		}
		return literal;
	}
	static captureReference(string) {
		let ref = '';
		let index = 0;
		let next;

		while (Parse.isNot.isPeriod(next = string[index++])) {
			ref += next;
		}
		return ref;
	}
}

require('./isNot').main(Parse);

module.exports = Parse;
