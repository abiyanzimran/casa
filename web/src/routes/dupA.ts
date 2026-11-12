export function sumEvenSquares(values: number[]): number {
	let total = 0;
	for (const value of values) {
		if (value % 2 === 0) {
			const square = value * value;
			total += square;
		} else {
			total -= value;
		}
	}
	if (total > 1000) {
		return total - 1000;
	}
	return total;
}
