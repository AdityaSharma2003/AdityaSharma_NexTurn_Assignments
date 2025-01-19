const {toggleVisibility} = require("../src/E5.js");

describe("Testing toggleVisibility function", () => {
    let dummyElement;

    beforeEach(() => {
        let displayValue = "block";
        dummyElement = {
            style: {
                get display() {
                    return displayValue;
                },
                set display(value) {
                    displayValue = value;
                },
            },
        };
    });

    test("handling case: hiding the element when it is initially visible", () => {
        const spy = jest.spyOn(dummyElement.style, "display", "set");

        toggleVisibility(dummyElement);

        expect(dummyElement.style.display).toBe("none");
        expect(spy).toHaveBeenCalledWith("none");

        spy.mockRestore(); 
    });

    test("handling case: shows the element when it is initially hidden", () => {
        dummyElement.style.display = "none";

        const spy = jest.spyOn(dummyElement.style, "display", "set");

        toggleVisibility(dummyElement);

        expect(dummyElement.style.display).toBe("block");
        expect(spy).toHaveBeenCalledWith("block");

        spy.mockRestore(); 
    });
});
