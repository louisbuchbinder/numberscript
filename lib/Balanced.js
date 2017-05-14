
class Balanced {
	static isBalanced(string) {
		let count = 0;
		let balanced = true;

		string.split('').forEach(c => {
			if (c === '(') {
				count ++;
			}
			if (c === ')') {
				count --;
			}
			if (count < 0) {
				balanced = false;
			}
		});

		if (balanced) {
			balanced = count === 0;
		}

		return balanced;
	}

	static findBalancedIndex(string) {
		if (!Balanced.isBalanced(string)) {
			throw new Error('Expected a string with balanced "()"');
		}
		if (string[0] !== '(') {
			throw new Error('Expected a leading (');
		}

		let count = 0;
		let index;
		let done = false;

		string.split('').forEach((c, i) => {
			if (c === '(') {
				count ++;
			}
			if (c === ')') {
				count --;
			}
			if (!done && count === 0) {
				index = i;
				done = true;
			}
		});

		return index;
	}
};

module.exports = Balanced;
