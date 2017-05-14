
class isNot {
	static main(Class) {
		Class.isNot = isNot;

		Object
			.getOwnPropertyNames(Class)
			.filter(name => [
				'length',
				'name',
				'isNot',
				'prototype'
			].every(n => n !== name))
			.filter(name => typeof Class[name] === 'function')
			.forEach(name => isNot[name] = character => !Class[name](character));
	}
}

module.exports = isNot;
