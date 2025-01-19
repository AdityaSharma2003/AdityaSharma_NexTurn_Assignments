const {getElement} = require("../src/E2.js");

describe('Testing getElement function for element retrieval', () => {
    let arr;
    beforeEach(() => {
        arr = [1,2,3,4,5];
    });

    it('handling element retrieval for index within range', () => {
        expect(getElement(arr,0)).toBe(1);
        expect(getElement(arr,1)).toBe(2);
        expect(getElement(arr,2)).toBe(3);
        expect(getElement(arr,3)).toBe(4);
        expect(getElement(arr,4)).toBe(5);
    });
    it('handling element retrieval for index outside range (negative)', () => {
        expect(() => getElement(arr,-1)).toThrow("Index out of bounds");
        expect(() => getElement(arr,-2)).toThrow("Index out of bounds");
    });
    it('handling element retrieval for index outside range (positive greateer than or equal to array length)', () => {
        expect(() => getElement(arr,5)).toThrow("Index out of bounds");
        expect(() => getElement(arr,8)).toThrow("Index out of bounds");
    });
});