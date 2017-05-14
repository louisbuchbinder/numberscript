
// const Balanced = require('./lib/Balanced');
const Stack = require('./lib/Stack');
const Handler = require('./lib/Handler');
const Parse = require('./lib/Parse');

class Interpreter extends Stack {
	constructor() {
		super();
	}

	eval(code) {
		if (typeof code !== 'string') {
			throw new Error('Expected to evaluate a string of code');
		}

		const handler = new Handler(this);

		while (code) {
			let charIndex = 0;
			let character = code[charIndex];

			charIndex += 0 ||
				handler.execPrint(character)	||
				handler.execZero(character)		||
				handler.execPointer(character)||
				handler.execNumber(character, code) ||
				handler.execLiteral(character, code) ||
				handler.execLoop(character, code) ||
				handler.execReference(character, code) ||
				handler.execCodeReference(character) ||
				1;

			code = code.slice(charIndex);
		}
	}
}

module.exports = Interpreter;
