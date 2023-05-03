const Calculator = require ('./calculator.js');
const caclculator = new Calculator();

it('1 + 2 = 3 >', () => {
    expect(caclculator.add(1, 2)).toEqual(3);
})