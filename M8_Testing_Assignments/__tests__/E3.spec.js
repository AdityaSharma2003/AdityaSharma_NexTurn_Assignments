const {delayedGreeting} = require("../src/E3.js");

describe('Testing async funtion delayedGreeting', () => {
    it('handling delayed response from funciton', () => {
        return delayedGreeting('aditya',2000).then((data) => {
            expect(data).toBe('Hello, aditya!');
        });
    });
});