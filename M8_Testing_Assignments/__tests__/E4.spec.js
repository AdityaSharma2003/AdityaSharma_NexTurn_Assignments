const {sendNotification} = require("../src/E4.js");

describe("Testing sendNotification function", () => {
    let dummyNotificationService;

    beforeEach(() => {
        dummyNotificationService = {
            send: jest.fn(),
        };
    });

    test("handling service when notification send is successful", () => {
        dummyNotificationService.send.mockReturnValue(true);

        const result = sendNotification(dummyNotificationService, "Hello aditya sharma!");
        expect(result).toBe("Notification Sent");
        expect(dummyNotificationService.send).toHaveBeenCalledWith("Hello aditya sharma!");
    });

    test("handling service when notification send fails", () => {
        dummyNotificationService.send.mockReturnValue(false);

        const result = sendNotification(dummyNotificationService, "Hello aditya sharma!");
        expect(result).toBe("Failed to Send");
        expect(dummyNotificationService.send).toHaveBeenCalledWith("Hello aditya sharma!");
    });
});
