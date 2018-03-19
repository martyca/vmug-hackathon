var vm = (function () {
    var self = this;
    self.connection = new signalR.HubConnection('/messagehub');

    self.commands = ["__roll"];

    self.messages = ko.observableArray();
    self.isTalking = ko.observable(false);

    self.roll = ko.observable(false);

    self.messageAdded = function (element) {
        $(element).hide().slideDown();
    }

    self.messageRemoved = function (element) {
        $(element).slideUp(() => { $(element).remove() });
    }

    self.handleSpecialMoves = function (command) {
        switch (command) {
            case "__roll":
                self.roll(true);
                setTimeout(() => {
                    self.roll(false);
                }, 1000);
                break;
        }
    }

    // Constructor

    self.connection.on("send", data => {
        for (let cmd of self.commands) {
            if (data === cmd) {
                self.handleSpecialMoves(data);
                return;
            }
        }

        self.messages.unshift(data);

        self.isTalking(true);
        setTimeout(() => { self.isTalking(false); }, 750);

        if (self.messages().length > 10) {
            self.messages.pop();
        }
    });

    self.connection.start();
})();

ko.applyBindings(vm, document.getElementById("app"));