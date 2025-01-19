const axios = require("axios");
const {fetchAndDisplayUser} = require("../src/Bonus.js"); 

jest.mock("axios"); 

describe("fetchAndDisplayUser with JSONPlaceholder mock", () => {
  let dummyElement;

  beforeEach(() => {
    let textContentValue = "";
    dummyElement = {
      get textContent() {
        return textContentValue;
      },
      set textContent(value) {
        textContentValue = value;
      },
    };
  });

  test("displays user name on successful fetch", async () => {
    axios.get.mockResolvedValueOnce({
      data: { id: 1, name: "Leanne Graham" },
    });

    const spy = jest.spyOn(dummyElement, "textContent", "set");

    await fetchAndDisplayUser(
      { getUser: (id) => axios.get(`https://jsonplaceholder.typicode.com/users/${id}`).then((res) => res.data) },
      1,
      dummyElement
    );

    expect(axios.get).toHaveBeenCalledWith("https://jsonplaceholder.typicode.com/users/1");
    expect(dummyElement.textContent).toBe("Hello, Leanne Graham");
    expect(spy).toHaveBeenCalledWith("Hello, Leanne Graham");

    spy.mockRestore();
  });

  test("displays error message when user fetch fails", async () => {
    axios.get.mockRejectedValueOnce(new Error("Network Error"));

    const spy = jest.spyOn(dummyElement, "textContent", "set");

    await fetchAndDisplayUser(
      { getUser: (id) => axios.get(`https://jsonplaceholder.typicode.com/users/${id}`).then((res) => res.data) },
      1,
      dummyElement
    );

    expect(axios.get).toHaveBeenCalledWith("https://jsonplaceholder.typicode.com/users/1");
    expect(dummyElement.textContent).toBe("Network Error");
    expect(spy).toHaveBeenCalledWith("Network Error");

    spy.mockRestore();
  });

  test("displays error message for invalid user data", async () => {
    axios.get.mockResolvedValueOnce({
      data: { id: 1 }, 
    });

    const spy = jest.spyOn(dummyElement, "textContent", "set");

    await fetchAndDisplayUser(
      { getUser: (id) => axios.get(`https://jsonplaceholder.typicode.com/users/${id}`).then((res) => res.data) },
      1,
      dummyElement
    );

    expect(axios.get).toHaveBeenCalledWith("https://jsonplaceholder.typicode.com/users/1");
    expect(dummyElement.textContent).toBe("Invalid user data");
    expect(spy).toHaveBeenCalledWith("Invalid user data");

    spy.mockRestore();
  });
});
