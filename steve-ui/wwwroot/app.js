var vm = (function () {
    var self = this;
    self.connection = new signalR.HubConnection('/messagehub');

    self.messages = ko.observableArray();
    self.isTalking = ko.observable(false);

    // Constructor

    self.connection.on("send", data => {
        self.messages.push(data);

        self.isTalking(true);
        setTimeout(() => { self.isTalking(false); }, 750);

        if (self.messages().length > 10) {
            self.messages.shift();
        }
    });

    self.connection.start();
})();

ko.applyBindings(vm, document.getElementById("app"));