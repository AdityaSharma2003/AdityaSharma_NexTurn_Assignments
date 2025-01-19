const {capitalize, reverseString} = require("../src/E1.js");

describe('Testing the capitalize function', () => {
    it('handling the corrected string all small', () => {
        expect(capitalize("aditya")).toBe("Aditya");
    });
    it('handling the empty string', () => {
        expect(capitalize("")).toBe("");
    });
    it('handling the single character string', () => {
        expect(capitalize("a")).toBe("A");
    });
    it('handling the corrected string with first character as capital', () => {
        expect(capitalize("Aditya")).toBe("Aditya");
    });
    it('handling the incorrect string with first character as special characters and spaces', () => {
        expect(capitalize(" aditya")).toBe(" aditya");
        expect(capitalize("@34aditya")).toBe("@34aditya");
    });
});

describe('Testing the reverseString function', () => {
    it('handling the string', () => {
        expect(reverseString("aditya")).toBe("aytida");
        expect(reverseString("Aditya")).toBe("aytidA");
    });
    it('handling the empty string', () => {
        expect(reverseString("")).toBe("");
    });
    it('handling the palindromic string', () => {
        expect(reverseString("racecar")).toBe("racecar");
        expect(reverseString("abbccbba")).toBe("abbccbba");
        expect(reverseString("12321")).toBe("12321");
    });
    it('handling the string with first character as special characters and spaces', () => {
        expect(reverseString(" aditya")).toBe("aytida ");
        expect(reverseString("@34aditya")).toBe("aytida43@");
    });
});